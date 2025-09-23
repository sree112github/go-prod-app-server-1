package deviceDataModels
import (
	"time"

	"github.com/google/uuid"
)

type DeviceDataModel struct {
	DataId int64 `json:"data_id"`

	DeviceId uuid.UUID `json:"device_id"`

	Metrics   map[string]any `json:"metrics"`
	CreatedAt time.Time      `json:"created_at"`
}

type DeviceDataInputModel struct {
	DeviceId uuid.UUID      `json:"device_id"`
	Metrics  map[string]any `json:"metrics"`
}

type DeviceDataResponseModel struct {
	DataId    int64          `json:"data_id"`
	DeviceId  uuid.UUID      `json:"device_id"`
	Metrics   map[string]any `json:"metrics"`
	CreatedAt time.Time      `json:"created_at"`
}
