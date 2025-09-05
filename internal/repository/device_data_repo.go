package repository

import (
	"basics/internal/config"
	"basics/internal/models"
	"encoding/json"
	"fmt"
)

func CreateDeviceData(deviceData *models.DeviceDataInputModel) (*models.DeviceDataResponseModel, error) {

	var response models.DeviceDataResponseModel

	metricsJson, err := json.Marshal(deviceData.Metrics)

	fmt.Println("The Metrics is :", string(metricsJson))
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO device_data(device_id,metrics) VALUES ($1,$2) RETURNING data_id,device_id,metrics,created_at`

	if err := config.DB.QueryRow(query, deviceData.DeviceId, metricsJson).
		Scan(&response.DataId, &response.DeviceId, &metricsJson, &response.CreatedAt); err != nil {
		fmt.Println("Error occured in the device data creation", err)
		return nil, err
	}

	err = json.Unmarshal(metricsJson, &response.Metrics)
	if err != nil {
		return nil, err
	}

	return &response, nil

}
