package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"fmt"
)

func CreateMachine(machine *models.MachineInputModel) (*models.MachineResponseModel, error) {

	machineResponse, err := repository.CreateMachine(machine)

	if err != nil {
		fmt.Println("error occured in the Service Layer During Machine Creation")
		return nil, err
	}

	return machineResponse, err
}
