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
	Scope       string    `json:"scope,omitempty"`
	PlantId     uuid.UUID `json:"plant_id,omitempty"`
	CompanyId   uuid.UUID `json:"company_id,omitempty"`
	MachineId   uuid.UUID `json:"machine_id,omitempty"`
	MachineName string    `json:"machine_name"`
}

type MachineResponseModel struct {
	MachineId   uuid.UUID `json:"machine_id"`
	MachineName string    `json:"machine_name"`
	PlantId     uuid.UUID `json:"plant_id"`
	PlantName   string    `json:"plant_name"`
	CreatedAt   time.Time `json:"created_at"`
}
