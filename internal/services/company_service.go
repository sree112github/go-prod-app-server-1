package services

import (
	"basics/internal/models"
	"basics/internal/repository"
)

func CreateCompany(company *models.CompanyInputModel) (*models.CompanyResponseModel, error) {

	return repository.CreateCompany(company)
}
