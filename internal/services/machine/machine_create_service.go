package machineservice

import (
	machineModels "basics/internal/models/machine"
	machinerepo "basics/internal/repository/machine"
	"fmt"
)

func CreateMachine(machine *machineModels.MachineInputModel) (*machineModels.MachineResponseModel, error) {

	machineResponse, err := machinerepo.CreateMachine(machine)

	if err != nil {
		fmt.Println("error occured in the Service Layer During Machine Creation")
		return nil, err
	}

	return machineResponse, err
}
