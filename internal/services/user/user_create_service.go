package userservice

import (
	userModels "basics/internal/models/user"
	userrepo "basics/internal/repository/user"
	"basics/internal/utils"
	"fmt"
)

func SignUp(user *userModels.SignUpInput) (*userModels.SignUpResponse, error) {

	hashedPassword, err := utils.GenerateHashedPassword(user.Password)
	if err != nil {
		fmt.Println("Password hashing failed:", err)
		return nil, err
	}

	fmt.Println("Password hashed successfully")

	signUpInput := userModels.SignUpInput{
		UserName: user.UserName,
		Email:    user.Email,
		Password: hashedPassword, // ✅ use the hashed password
	}

	return userrepo.CreateUser(&signUpInput)
}
