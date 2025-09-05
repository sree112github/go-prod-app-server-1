package services

import (
	"basics/internal/models"
	"basics/internal/repository"
	"fmt"
)

func CreateDevice(device *models.DeviceInputModel)(*models.DeviceResponseModel,error){

	deviceResponse,err:= repository.CreateDevice(device)

	if(err != nil){
		fmt.Println("The Error Occured in the Create Device service layer")
		return nil,err
	}

	return  deviceResponse,nil

}