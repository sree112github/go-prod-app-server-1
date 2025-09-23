package machinerepo

import (
	"basics/internal/config"
)

func IsMachineExists(companyId,plantId,machineId string) (bool, error) {
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
