package deviceModels

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

// Get all device Response structure
type GetAllDeviceResponseStructure struct {
	DeviceId    uuid.UUID `json:"device_id"`
	DeviceName  string    `json:"name"`
	MachineId   uuid.UUID `json:"machine_id"`
	MachineName string    `json:"machine_name"`
	PlantId     uuid.UUID `json:"plant_id"`
	PlantName   string    `json:"plant_name"`
	CompanyId   uuid.UUID `json:"company_id"`
	CompanyName string    `json:"company_name"`
}
