package models

import (
	"time"

	"github.com/google/uuid"
)

type DeviceModel struct {
	DeviceId  uuid.UUID `json:"device_id"`
	MachineId uuid.UUID `json:"machine_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type DeviceInputModel struct {
	MachineId uuid.UUID `json:"machine_id"`
	Name      string    `json:"name"`
}

type DeviceResponseModel struct {
	DeviceId  uuid.UUID `json:"device_id"`
	MachineId uuid.UUID `json:"machine_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
