package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func CreateCompany(company *models.CompanyModel) (*models.CompanyModel, error) {
	query := `
		INSERT INTO companies(company_name)
		VALUES($1)
		RETURNING company_id, company_name, created_at
	`

	var companyId string
	err := config.DB.QueryRow(query, company.CompanyName).Scan(
		&companyId, &company.CompanyName, &company.CreatedAt,
	)
	if err != nil {
		fmt.Println("❌ SQL Query error during scanning:", err)
		return nil, err
	}

	company.CompanyId, err = uuid.Parse(companyId)
	if err != nil {
		fmt.Println("❌ UUID parsing error:", err)
		return nil, fmt.Errorf("failed to parse company ID: %w", err)
	}

	return company, nil
}
