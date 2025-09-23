package machineModels

import (
	"time"

	"github.com/google/uuid"
)

type MachineModel struct {
	MachineId   uuid.UUID `json:"machine_id"`
	PlantId     uuid.UUID `json:"plant_id"`
	MachineName string    `json:"machine_name"`
	CreatedAt   string    `json:"created_at"`
}

type MachineInputModel struct {
	PlantId     uuid.UUID `json:"plant_id"`
	MachineName string    `json:"machine_name"`
}

type MachineResponseModel struct {
	MachineId   uuid.UUID `json:"machine_id"`
	PlantId     uuid.UUID `json:"plant_id"`
	MachineName string    `json:"machine_name"`
	CreatedAt   time.Time `json:"created_at"`
}
