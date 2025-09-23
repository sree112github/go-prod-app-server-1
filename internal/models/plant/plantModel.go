package plantModels

import (
	"time"

	"github.com/google/uuid"
)

type PlantModel struct {
	PlantId   uuid.UUID `json:"plant_id"`
	CompanyId uuid.UUID `json:"company_id"`
	PlantName string    `json:"plant_name"`
	CreatedAt time.Time `json:"created_at"`
}

type PlantInputModel struct {
	Scope     string    `json:"scope,omitempty"`
	PlantId   uuid.UUID `json:"plant_id,omitempty"`
	CompanyId uuid.UUID `json:"company_id,omitempty"`
	MachineId uuid.UUID `json:"machine_id,omitempty"`
	PlantName string    `json:"plant_name,omitempty"`
}

type PlantResponseModel struct {
	PlantId     uuid.UUID `json:"plant_id"`
	CompanyId   uuid.UUID `json:"company_id"`
	CompanyName string    `json:"company_name"`
	PlantName   string    `json:"plant_name"`
	CreatedAt   time.Time `json:"created_at"`
}
