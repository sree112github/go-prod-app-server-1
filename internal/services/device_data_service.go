package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"fmt"
)

func CreateDeviceService(deviceData *models.DeviceDataInputModel) (*models.DeviceDataResponseModel, error) {

	response, err := repository.CreateDeviceData(deviceData)

	if err != nil {
		fmt.Println("The error Occur in service layer ", err)
		return nil, err
	}

	return response, nil
}
