package deviceservice

import (
	deviceModels "basics/internal/models/device"
	devicerepo "basics/internal/repository/device"

	"fmt"
)

func CreateDevice(device *deviceModels.DeviceInputModel) (*deviceModels.DeviceResponseModel, error) {

	deviceResponse, err := devicerepo.CreateDevice(device)

	if err != nil {
		fmt.Println("The Error Occured in the Create Device service layer")
		return nil, err
	}

	return deviceResponse, nil

}