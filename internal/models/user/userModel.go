package userModels

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	UserId     uuid.UUID  `json:"user_id"`
	UserName   string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"-"`
	UserScope  string     `json:"scope"`
	ComapanyId *uuid.UUID `json:"company_id,omitempty"`
	PlantId    *uuid.UUID `json:"plant_id,omitempty"`
	MachineId  *uuid.UUID `json:"machine_id,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}

type AssignRoleInput struct {
	UserId    string  `json:"user_id" binding:"required,uuid"`
	Scope     string  `json:"scope" binding:"required,oneof= super_admin company_admin plant_admin machine_user"`
	CompanyId *string `json:"company_id,omitempty"`
	PlantId   *string `json:"plant_id,omitempty"`
	MachineId *string `json:"machine_id,omitempty"`
}

// Response
type AssignRoleResponse struct {
	UserId    string    `json:"user_id"`
	UserName  string    `json:"name"`
	Email     string    `json:"email"`
	Scope     string    `json:"scope"`
	CompanyId *string   `json:"company_id,omitempty"`
	PlantId   *string   `json:"plant_id,omitempty"`
	MachineId *string   `json:"machine_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
