package models

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
	CompanyId uuid.UUID `json:"company_id"`
	PlantName string    `json:"plant_name"`
}

type PlantResponseModel struct {
	PlantId   uuid.UUID `json:"plant_id"`
	CompanyId uuid.UUID `json:"company_id"`
	PlantName string    `json:"plant_name"`
	CreatedAt time.Time `json:"created_at"`
}
