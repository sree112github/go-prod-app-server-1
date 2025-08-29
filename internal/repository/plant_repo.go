package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"fmt"

	"github.com/google/uuid"
)

func CreatePlant(plant *models.PlantModel) (*models.PlantModel, error) {

	query := `INSERT INTO plants(company_id,plant_name) VALUES($1,$2) RETURNING company_id,plant_id,plant_name,Created_at`

	var companyId string
	var plantId string
	fmt.Println("the Plant Name is ", plant.PlantName)
	err := config.DB.QueryRow(query, plant.CompanyId, plant.PlantName).Scan(&companyId, &plantId, &plant.PlantName, &plant.CreatedAt)

	if err != nil {
		fmt.Println("Error occured During Creating Plant due to:", err)
		return nil, err
	}

	plant.CompanyId, err = uuid.Parse(companyId)
	if err != nil {
		fmt.Println("❌ UUID parsing error:", err)
		return nil, fmt.Errorf("failed to parse company ID: %w", err)
	}

	plant.PlantId, err = uuid.Parse(plantId)
	if err != nil {
		fmt.Println("❌ UUID parsing error:", err)
		return nil, fmt.Errorf("failed to parse plantID ID: %w", err)
	}

	return plant, nil

}
