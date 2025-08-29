package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignUpInput struct {
	UserName string `json:"name" binding:"required"`
	Scope    string `json:"scope" binding:"required,oneof=super_admin company_admin plant_admin machine_admin"`
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
	UserId   string `json:"user_id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Scope    string `json:"scope"`
}

func SignUp(c *gin.Context) {

	var input SignUpInput

	//Step1; bind with the json from request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := services.SignUp(input.UserName, input.Scope, input.Email, input.Password)
	if err != nil {
		fmt.Println("❌ Signup service error:", err) // backend log
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := SignUpResponse{
		Message:  "User Created Successfully",
		UserId:   users.UserId.String(), // convert UUID to string
		UserName: users.UserName,        // matches UserModel field
		Email:    users.Email,
		Scope:    users.UserScope,
	}

	c.JSON(http.StatusCreated, response)

}

func Login(c *gin.Context) {

	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, token, err := services.Login(input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := LoginResponse{
		UserId:   user.UserId.String(),
		UserName: user.UserName,
		Email:    user.Email,
		Token:    token,
		Scope:    user.UserScope,
	}

	c.JSON(http.StatusCreated, response)

}
