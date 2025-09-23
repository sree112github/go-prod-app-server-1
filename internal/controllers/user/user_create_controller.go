package usercontroller

import (
	userModels "basics/internal/models/user"
	userservice "basics/internal/services/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var input userModels.SignUpInput

	//Step1; bind with the json from request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := userservice.SignUp(&input)
	if err != nil {
		fmt.Println("❌ Signup service error:", err) // backend log
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := userModels.SignUpResponse{
		Message:  "User Created Successfully",
		UserId:   users.UserId,   // convert UUID to string
		UserName: users.UserName, // matches UserModel field
		Email:    users.Email,
		Scope:    users.Scope,
	}

	c.JSON(http.StatusCreated, response)

}
