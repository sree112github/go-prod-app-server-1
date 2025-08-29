package models

import "github.com/google/uuid"

type UserModel struct {
	UserName  string    `json:"name"`
	UserScope string    `json:"scope"`
	UserId    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
