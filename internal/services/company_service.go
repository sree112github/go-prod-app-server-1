package services

import (
	"basics/internal/models"
	"basics/internal/repository"
)

func CreateCompany(company *models.CompanyModel) (*models.CompanyModel, error) {

	return repository.CreateCompany(company)
}
