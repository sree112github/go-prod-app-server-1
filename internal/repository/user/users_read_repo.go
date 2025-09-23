package userrepo

import (
	"basics/internal/config"
	userModels "basics/internal/models/user"
	"database/sql"

	"fmt"
)

func Login(user *userModels.LoginInput) (*userModels.LoginResponse, error) {
	var response userModels.LoginResponse

	query := `
		SELECT 
			user_id,
			email,
			name,
			scope,
			password,
			COALESCE(company_id::text, ''),
			COALESCE(plant_id::text, ''),
			COALESCE(machine_id::text, '')
		FROM users 
		WHERE email=$1
	`

	err := config.DB.QueryRow(query, user.Email).
		Scan(&response.UserId, &response.Email, &response.UserName, &response.Scope,
			&response.Password, &response.CompanyId, &response.PlantId, &response.MachineId)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return &response, nil
}

func GetAllUserInfoBasedOnScope(input *userModels.AssignRoleInput, limit, offset int) ([]userModels.UserModel, error) {

	var rows *sql.Rows
	var err error

	switch input.Scope {
	case "super_admin":
		rows, err = config.DB.Query(`SELECT user_id,name,email,scope,company_id,plant_id,machine_id,created_at FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)

	case "company_admin":
		rows, err = config.DB.Query(
			`SELECT user_id,name,email,scope,company_id,plant_id,machine_id,created_at FROM users WHERE company_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, input.CompanyId, limit, offset)

	case "plant_admin":
		rows, err = config.DB.Query(
			`SELECT user_id,name,email,scope,company_id,plant_id,machine_id,created_at FROM users WHERE plant_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, input.PlantId, limit, offset)

	case "machine_admin":
		rows, err = config.DB.Query(
			`SELECT user_id,name,email,scope,company_id,plant_id,machine_id,created_at FROM users WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, input.UserId, limit, offset)
	default:
		return nil, fmt.Errorf("invalid Scope is Defined %v", input.Scope)

	}

	if err != nil {
		return nil, fmt.Errorf("the error Occured in the scope %v Due to %v", input.Scope, err)
	}

	defer rows.Close()

	var users []userModels.UserModel

	for rows.Next() {

		var u userModels.UserModel

		if err := rows.Scan(&u.UserId, &u.UserName, &u.Email, &u.UserScope, &u.ComapanyId, &u.PlantId, &u.MachineId, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}
