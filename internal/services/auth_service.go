package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"basics/internal/utils"
	"fmt"
)

func SignUp(user *models.SignUpInput) (*models.SignUpResponse, error) {

	hashedPassword, err := utils.GenerateHashedPassword(user.Password)
	if err != nil {
		fmt.Println("Password hashing failed:", err)
		return nil, err
	}

	fmt.Println("Password hashed successfully")

	signUpInput := models.SignUpInput{
		UserName: user.UserName,
		Email:    user.Email,
		Password: hashedPassword, // ✅ use the hashed password
	}

	return repository.CreateUser(&signUpInput)
}

func Login(user *models.LoginInput) (*models.LoginResponse, error) {
	response, err := repository.GetUserByEmail(user)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if !utils.CheckPasswordIsCorrect(user.Password, response.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	// Generate JWT
	response, err = utils.GenerateJWT(response)
	if err != nil {
		return nil, fmt.Errorf("could not generate token: %w", err)
	}

	return response, nil
}

func AssignUserRole(input *models.AssignRoleInput) (*models.AssignRoleResponse, error) {

	user, err := repository.AssignUserRole(input)
	if err != nil {
		return nil, err
	}

	return user, nil

}
