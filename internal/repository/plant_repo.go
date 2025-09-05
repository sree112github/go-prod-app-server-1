package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"fmt"

	"github.com/lib/pq"
)

func CreatePlant(plant *models.PlantInputModel) (*models.PlantResponseModel, error) {

	var response models.PlantResponseModel

	tnx, err := config.DB.Begin()

	if err != nil {
		fmt.Println("Database error due to ", err)
		return nil, err
	}
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM companies WHERE company_id = $1)`

	err = tnx.QueryRow(query, plant.CompanyId).Scan(&exist)

	if err != nil {
		tnx.Rollback()
		fmt.Println("Failed to chack Company existance")
		return nil, err
	}

	query = `INSERT INTO plants(company_id,plant_name) VALUES($1,$2) RETURNING company_id,plant_id,plant_name,Created_at`

	fmt.Println("the Plant Name is ", plant.PlantName)
	err = tnx.QueryRow(query, plant.CompanyId, plant.PlantName).Scan(&response.CompanyId, &response.PlantId, &response.PlantName, &response.CreatedAt)

	if err != nil {
		// 👇 Check if it's a Postgres error
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23503" { // 23503 = foreign_key_violation
				return nil, fmt.Errorf("company with id %s does not exist", plant.CompanyId)
			}
		}
		tnx.Rollback()
		fmt.Println("Error occured During Creating Plant due to:", err)
		return nil, err
	}

	// ✅ Commit if both succeed
	if err := tnx.Commit(); err != nil {
		return nil, err
	}

	return &response, nil

}
