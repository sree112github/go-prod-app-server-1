package devicedataservice

import (
	deviceDataModels "basics/internal/models/device_data"
	devicedatarepo "basics/internal/repository/device_data"
	"fmt"
)

func CreateDeviceData(deviceData *deviceDataModels.DeviceDataInputModel) (*deviceDataModels.DeviceDataResponseModel, error) {

	response, err := devicedatarepo.CreateDeviceData(deviceData)

	if err != nil {
		fmt.Println("The error Occur in service layer ", err)
		return nil, err
	}

	return response, nil
}
