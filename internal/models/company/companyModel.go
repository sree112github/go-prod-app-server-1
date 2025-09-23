package companyModels

import (
	"time"

	"github.com/google/uuid"
)

type CompanyModel struct {
	CompanyId   uuid.UUID `json:"company_id"`
	CompanyName string    `json:"company_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type CompanyInputModel struct {
	CompanyName string `json:"company_name"`
}

type CompanyResponseModel struct {
	CompanyId   uuid.UUID `json:"company_id"`
	CompanyName string    `json:"company_name"`
	CreatedAt   time.Time `json:"created_at"`
}
