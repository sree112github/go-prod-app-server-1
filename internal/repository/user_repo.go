package repository

import (
	"basics/internal/config"
	"basics/internal/models"

	"fmt"

	"github.com/google/uuid"
)

func CreateUser(user *models.UserModel) (*models.UserModel, error) {
	fmt.Println("Reached repository for Signup")

	// SQL query
	query := `
		INSERT INTO users(name, email, password, scope)
		VALUES($1, $2, $3, $4)
		RETURNING user_id, name, email, scope
	`

	// Scan returned values
	var userID string
	err := config.DB.QueryRow(query, user.UserName, user.Email, user.Password, user.UserScope).
		Scan(&userID, &user.UserName, &user.Email, &user.UserScope)
	if err != nil {
		fmt.Println("❌ DB Insert Error:", err)
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	// Convert string to uuid.UUID
	user.UserId, err = uuid.Parse(userID)
	if err != nil {
		fmt.Println("❌ UUID parsing error:", err)
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	fmt.Println("✅ User inserted successfully into DB")
	return user, nil
}

func GetUserByEmail(email string) (*models.UserModel, error) {

	var user models.UserModel

	query := `SELECT user_id, email, name, scope, password FROM users where email=$1`

	err := config.DB.QueryRow(query, email).Scan(&user.UserId, &user.Email, &user.UserName, &user.UserScope, &user.Password)

	if err != nil {
		return nil, fmt.Errorf("user Not found %s", err)
	}

	return &user, nil
}
