package machinerepo

import (
	"basics/internal/config"
	machineModels "basics/internal/models/machine"
	"database/sql"
	"fmt"
)

func GetAllMachinesBasedonScope(machine *machineModels.MachineInputModel, limit, page int) ([]machineModels.MachineResponseModel, error) {

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	var rows *sql.Rows
	var err error

	switch machine.Scope {
	case "super_admin":
		rows, err = config.DB.Query(`SELECT m.machine_id,m.machine_name,p.plant_id,p.plant_name,m.created_at 
		FROM machines m JOIN plants p ON 
		m.plant_id = p.plant_id ORDER by m.created_at LIMIT $1 OFFSET $2`, limit, offset)

	case "company_admin":
		rows, err = config.DB.Query(`SELECT m.machine_id,m.machine_name,p.plant_id,p.plant_name,m.created_at 
		FROM machines m JOIN plants p ON 
		m.plant_id = p.plant_id WHERE p.company_id= $1 ORDER by m.created_at LIMIT $2 OFFSET $3`, machine.CompanyId, limit, offset)

	case "plant_admin":
		rows, err = config.DB.Query(`SELECT m.machine_id, m.machine_name, p.plant_id, p.plant_name, m.created_at 
    FROM machines m 
    JOIN plants p ON m.plant_id = p.plant_id 
    WHERE p.company_id = $1 AND p.plant_id = $2
    ORDER BY m.created_at LIMIT $3 OFFSET $4`,
			machine.CompanyId, machine.PlantId, limit, offset)

	case "machine_user":
		rows, err = config.DB.Query(`SELECT m.machine_id,m.machine_name,p.plant_id,p.plant_name,m.created_at 
		FROM machines m JOIN plants p ON 
		m.plant_id = p.plant_id WHERE p.company_id= $1 AND p.plant_id= $2 AND m.machine_id = $3
		ORDER by m.created_at LIMIT $4 OFFSET $5`, machine.CompanyId, machine.PlantId, machine.MachineId, limit, offset)

	default:
		return nil, fmt.Errorf("invalid scope: %v", machine.Scope)
	}

	if err != nil {
		fmt.Println("Query execution error due to:", err)
		return nil, fmt.Errorf("query execution error")
	}

	defer rows.Close()

	var machines []machineModels.MachineResponseModel

	for rows.Next() {
		var m machineModels.MachineResponseModel
		if err := rows.Scan(&m.MachineId, &m.MachineName, &m.PlantId, &m.PlantName, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		machines = append(machines, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return machines, nil

}

func IsMachineExists(companyId, plantId, machineId string) (bool, error) {
	query := `SELECT EXISTS (
        SELECT 1
        FROM plants p
        JOIN machines m ON p.plant_id = m.plant_id
        WHERE p.company_id = $1 AND p.plant_id = $2 AND m.machine_id = $3
    )`
	var exists bool
	err := config.DB.QueryRow(query, companyId, plantId, machineId).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
