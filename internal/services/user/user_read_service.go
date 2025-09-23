package userservice

import (
	userModels "basics/internal/models/user"
	userrepo "basics/internal/repository/user"
	"basics/internal/utils"
	"fmt"
)



func Login(user *userModels.LoginInput) (*userModels.LoginResponse, error) {
	response, err := userrepo.Login(user)
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




func GetAllUserInfoBasedOnScope(input *userModels.AssignRoleInput, limit, offset int) ([]userModels.UserModel, error) {

	return userrepo.GetAllUserInfoBasedOnScope(input, limit, offset)
}
