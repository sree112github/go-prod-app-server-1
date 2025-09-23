package companyrepo

import (
	"basics/internal/config"
	companyModels "basics/internal/models/company"
	userModels "basics/internal/models/user"
	"database/sql"
	"fmt"
)

func GetAllCompanyBasedOnScope(company *userModels.AssignRoleInput, limit, page int) ([]companyModels.CompanyResponseModel, error) {
	var rows *sql.Rows
	var err error

	offset := (page - 1) * limit

	switch company.Scope {
	case "super_admin":
		rows, err = config.DB.Query(`
            SELECT company_id, company_name, created_at 
            FROM companies 
            ORDER BY created_at DESC 
            LIMIT $1 OFFSET $2`, limit, offset)

	case "company_admin":
		rows, err = config.DB.Query(`
            SELECT company_id, company_name, created_at 
            FROM companies 
            WHERE company_id = $1 
            ORDER BY created_at DESC 
            LIMIT $2 OFFSET $3`,
			company.CompanyId, limit, offset)

	case "plant_admin":
		if company.PlantId == nil {
			return nil, fmt.Errorf("plant_id is required for plant_admin scope")
		}
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, c.created_at 
            FROM companies c
            JOIN plants p ON c.company_id = p.company_id
            WHERE p.plant_id = $1 
            ORDER BY p.created_at DESC 
            LIMIT $2 OFFSET $3`,
			*company.PlantId, limit, offset)

	case "machine_user":
		fmt.Println("Hello")
		if company.MachineId == nil {
			return nil, fmt.Errorf("machine_id is required for machine_user scope")
		}
		rows, err = config.DB.Query(`
            SELECT c.company_id, c.company_name, c.created_at 
            FROM companies c
            JOIN plants p ON c.company_id = p.company_id
            JOIN machines m ON p.plant_id = m.plant_id
            WHERE m.machine_id = $1
            ORDER BY m.created_at DESC 
            LIMIT $2 OFFSET $3`,
			*company.MachineId, limit, offset) // ✅ dereference

	default:
		return nil, fmt.Errorf("invalid scope: %v", company.Scope)
	}

	if err != nil {
		return nil, fmt.Errorf("query execution error: %v", err)
	}
	defer rows.Close()

	var companies []companyModels.CompanyResponseModel
	for rows.Next() {
		var c companyModels.CompanyResponseModel
		if err := rows.Scan(&c.CompanyId, &c.CompanyName, &c.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}

		// ✅ Debug print
		fmt.Printf("DEBUG: CompanyID=%s, Name=%s, CreatedAt=%s\n",
			c.CompanyId, c.CompanyName, c.CreatedAt)

		companies = append(companies, c)
	}

	return companies, nil
}

func IsCompanyExist(companyId *string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM companies WHERE company_id = $1)`
	var exists bool
	err := config.DB.QueryRow(query, companyId).Scan(&exists)
	if err != nil {
		// handle error appropriately, possibly log and return false
		return false, err
	}
	fmt.Println("Company exist")
	return exists, nil
}
