package userModels

import "github.com/google/uuid"


type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignUpInput struct {
	UserName string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type SignUpResponse struct {
	Message  string `json:"message"`
	UserId   string `json:"user_id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Scope    string `json:"scope"`
}

type LoginResponse struct {
	UserId    uuid.UUID `json:"user_id"`
	CompanyId string    `json:"company_id"`
	PlantId   string    `json:"plant_id"`
	MachineId string    `json:"machine_id"`
	UserName  string    `json:"name"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	Scope     string    `json:"scope"`
}
