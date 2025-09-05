package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"fmt"
)

func CreatePlant(plant *models.PlantInputModel) (*models.PlantResponseModel, error) {

	response, err := repository.CreatePlant(plant)
	if err != nil {
		fmt.Println("error occured in the Service Layer During Plant Creation")
		return nil, err
	}

	return response, nil
}
