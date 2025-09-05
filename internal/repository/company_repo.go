package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"fmt"
)

func CreateCompany(company *models.CompanyInputModel) (*models.CompanyResponseModel, error) {

	var response models.CompanyResponseModel

	query := `
		INSERT INTO companies(company_name)
		VALUES($1)
		RETURNING company_id, company_name, created_at
	`

	err := config.DB.QueryRow(query, company.CompanyName).Scan(
		&response.CompanyId, &response.CompanyName, &response.CreatedAt,
	)
	if err != nil {
		fmt.Println("❌ SQL Query error during scanning:", err)
		return nil, err
	}

	return &response, nil
}
