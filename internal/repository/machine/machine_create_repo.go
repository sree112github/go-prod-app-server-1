package machinerepo

import (
	"basics/internal/config"
	machineModels "basics/internal/models/machine"
	"fmt"
)

func CreateMachine(machine *machineModels.MachineInputModel) (*machineModels.MachineResponseModel, error) {

	query := `INSERT INTO machines(plant_id,machine_name) VALUES($1,$2) RETURNING machine_id,plant_id,machine_name,created_at`

	var machineResponse machineModels.MachineResponseModel

	err := config.DB.QueryRow(query, machine.PlantId, machine.MachineName).
		Scan(&machineResponse.MachineId, &machineResponse.PlantId, &machineResponse.MachineName, &machineResponse.CreatedAt)

	if err != nil {
		fmt.Println("❌ SQL Query error during scanning:", err)
		return nil, err
	}

	return &machineResponse, nil
}
