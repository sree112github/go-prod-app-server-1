package services

import (
	"basics/internal/models"
	"basics/internal/repository"
)

func GetAllUserInfoBasedOnScope(input *models.AssignRoleInput, limit, offset int) ([]models.UserModel, error) {

	return repository.GetAllUserInfoBasedOnScope(input, limit, offset)
}
