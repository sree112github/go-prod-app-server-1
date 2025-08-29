package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"basics/internal/utils"
	"fmt"
)

func SignUp(user_name string, scope string, email string, password string) (*models.UserModel, error) {
	hashedPassword, err := utils.GenerateHashedPassword(password)
	if err != nil {
		fmt.Println("Password hashing failed:", err)
		return nil, err
	}

	fmt.Println("Password hashed successfully")

	user := models.UserModel{
		Email:     email,
		Password:  string(hashedPassword),
		UserScope: scope,
		UserName:  user_name,
	}

	return repository.CreateUser(&user)
}

func Login(email string, password string) (*models.UserModel, string, error) {

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		fmt.Println("!")
		return nil, "", fmt.Errorf("Invalid Email or password")
	}

	if !utils.CheckPasswordIsCorrect(password, user.Password) {
		return nil, "", fmt.Errorf("Invalid Email or Password")
	}

	//Generate JWT

	token, err := utils.GenerateJWT(user)

	if err != nil {
		return nil, "", fmt.Errorf("Could Not generate Token")
	}

	return user, token, nil

}
