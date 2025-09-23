package plantrepo

import (
	"basics/internal/config"
	plantModels "basics/internal/models/plant"
	"database/sql"
	"fmt"
)

func GetAllPlantsBasedOnScope(plant *plantModels.PlantInputModel, limit int, page int) ([]plantModels.PlantResponseModel, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	var rows *sql.Rows
	var err error

	fmt.Println("the page and limit", page, limit)

	switch plant.Scope {
	case "super_admin":
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, p.plant_id, p.plant_name, p.created_at
            FROM companies c 
            JOIN plants p ON c.company_id = p.company_id 
            ORDER BY p.created_at DESC
            LIMIT $1 OFFSET $2`, limit, offset)

	case "company_admin":
		fmt.Println("This is COmpany Admin", plant.CompanyId)
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, p.plant_id, p.plant_name, p.created_at
            FROM companies c 
            JOIN plants p ON c.company_id = p.company_id 
            WHERE c.company_id = $1 
            ORDER BY p.created_at DESC
            LIMIT $2 OFFSET $3`, plant.CompanyId, limit, offset)

	case "plant_admin":
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, p.plant_id, p.plant_name, p.created_at
            FROM companies c 
            JOIN plants p ON c.company_id = p.company_id 
            WHERE c.company_id = $1 AND p.plant_id = $2 
            ORDER BY p.created_at DESC
            LIMIT $3 OFFSET $4`, plant.CompanyId, plant.PlantId, limit, offset)

	case "machine_user":
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, p.plant_id, p.plant_name, p.created_at
            FROM companies c 
            JOIN plants p ON c.company_id = p.company_id
            JOIN machines m ON p.plant_id = m.plant_id 
            WHERE c.company_id = $1 AND p.plant_id = $2 AND m.machine_id = $3 
            ORDER BY p.created_at DESC
            LIMIT $4 OFFSET $5`, plant.CompanyId, plant.PlantId, plant.MachineId, limit, offset)

	default:
		return nil, fmt.Errorf("invalid scope: %v", plant.Scope)
	}

	if err != nil {
		return nil, fmt.Errorf("query execution error: %v", err)
	}
	defer rows.Close()

	var plants []plantModels.PlantResponseModel
	for rows.Next() {
		var p plantModels.PlantResponseModel
		if err := rows.Scan(&p.CompanyId, &p.CompanyName, &p.PlantId, &p.PlantName, &p.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		fmt.Printf("DEBUG: CompanyID=%s, Name=%s, CreatedAt=%v\n",
			p.CompanyId, p.CompanyName, p.CreatedAt)

		plants = append(plants, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return plants, nil
}

func IsPlantExists(companyId,plantId string) (bool, error) {

	fmt.Println("The Company id is :", companyId)
	fmt.Println("the Plant Id is :", plantId)

	query := `SELECT EXISTS (SELECT 1 FROM plants WHERE company_id =$1 AND plant_id=$2)`

	var exists bool

	err := config.DB.QueryRow(query, companyId, plantId).Scan(&exists)

	if err != nil {
		return false, err
	}

	fmt.Println("The Plant exits is:", exists)

	return exists, nil

}
