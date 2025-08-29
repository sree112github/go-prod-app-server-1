package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"fmt"
)

func CreatePlant(plant *models.PlantModel) (*models.PlantModel, error) {

	plant, err := repository.CreatePlant(plant)
	if err != nil {
		fmt.Println("error occured in the Service Layer During Plant Creation")
		return nil, err
	}

	return plant, nil
}
