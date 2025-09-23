package devicerepo

import (
	"basics/internal/config"
	deviceModels "basics/internal/models/device"
	"fmt"
)

func CreateDevice(device *deviceModels.DeviceInputModel) (*deviceModels.DeviceResponseModel, error) {

	var deviceResponse deviceModels.DeviceResponseModel

	query := `INSERT INTO devices(machine_id,name) VALUES($1,$2) RETURNING device_id,machine_id,name,created_at`

	err := config.DB.QueryRow(query, device.MachineId, device.Name).Scan(&deviceResponse.DeviceId, &deviceResponse.MachineId, &deviceResponse.Name, &deviceResponse.CreatedAt)

	if err != nil {
		fmt.Println("error occured during device creation,", err)
		return nil, err
	}

	return &deviceResponse, nil
}
