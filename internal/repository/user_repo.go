package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"database/sql"

	"fmt"
)

func CreateUser(user *models.SignUpInput) (*models.SignUpResponse, error) {
	fmt.Println("Reached repository for Signup")

	var response models.SignUpResponse

	query := `
        INSERT INTO users(name, email, password)
        VALUES($1, $2, $3)
        RETURNING user_id, name, email, scope
    `

	err := config.DB.QueryRow(query, user.UserName, user.Email, user.Password).
		Scan(&response.UserId, &response.UserName, &response.Email, &response.Scope)
	if err != nil {
		fmt.Println("❌ DB Insert Error:", err)
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	fmt.Println("✅ User inserted successfully into DB")
	return &response, nil
}

func GetUserByEmail(user *models.LoginInput) (*models.LoginResponse, error) {
	var response models.LoginResponse

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

func AssignUserRole(input *models.AssignRoleInput) (*models.AssignRoleResponse, error) {

	query := `
	UPDATE users
	SET scope = $2,
	company_id =$3,
	plant_id =$4,
	machine_id =$5
	WHERE user_id = $1
	RETURNING user_id,name,email,scope,company_id,plant_id,machine_id,created_at
	`
	var response models.AssignRoleResponse
	err := config.DB.QueryRow(
		query, input.UserId, input.Scope,
		input.CompanyId, input.PlantId,
		input.MachineId).Scan(
		&response.UserId, &response.UserName, &response.Email,
		&response.Scope, &response.CompanyId, &response.PlantId,
		&response.MachineId, &response.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to assign user role: %w", err)
	}

	return &response, nil

}

func GetAllUserInfoBasedOnScope(input *models.AssignRoleInput, limit, offset int) ([]models.UserModel, error) {

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

	var users []models.UserModel

	for rows.Next() {

		var u models.UserModel

		if err := rows.Scan(&u.UserId, &u.UserName, &u.Email, &u.UserScope, &u.ComapanyId, &u.PlantId, &u.MachineId, &u.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil

}
