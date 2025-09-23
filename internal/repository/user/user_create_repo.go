package userrepo

import (
	"basics/internal/config"
	userModels "basics/internal/models/user"
	"fmt"
)


func CreateUser(user *userModels.SignUpInput) (*userModels.SignUpResponse, error) {
	fmt.Println("Reached repository for Signup")

	var response userModels.SignUpResponse

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