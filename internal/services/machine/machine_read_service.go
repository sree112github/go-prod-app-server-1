package machineservice

import (
	machineModels "basics/internal/models/machine"
	machinerepo "basics/internal/repository/machine"
	"fmt"
)

func GetAllMachinesBasedonScope(machine *machineModels.MachineInputModel, limit, page int) ([]machineModels.MachineResponseModel, error) {

	machines, err := machinerepo.GetAllMachinesBasedonScope(machine, limit, page)

	if err != nil {
		fmt.Println("something went wrong")
		return nil, fmt.Errorf("something wrong in service : %w", err)
	}
	return machines, nil
}
