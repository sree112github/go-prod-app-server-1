package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var input models.SignUpInput

	//Step1; bind with the json from request

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := services.SignUp(&input)
	if err != nil {
		fmt.Println("❌ Signup service error:", err) // backend log
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.SignUpResponse{
		Message:  "User Created Successfully",
		UserId:   users.UserId,   // convert UUID to string
		UserName: users.UserName, // matches UserModel field
		Email:    users.Email,
		Scope:    users.Scope,
	}

	c.JSON(http.StatusCreated, response)

}

func Login(c *gin.Context) {

	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.Login(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := models.LoginResponse{
		UserId:    user.UserId,
		UserName:  user.UserName,
		Email:     user.Email,
		Token:     user.Token,
		Scope:     user.Scope,
		CompanyId: user.CompanyId,
		PlantId:   user.PlantId,
		MachineId: user.MachineId,
	}

	c.JSON(http.StatusCreated, response)

}

func AssignUserRole(c *gin.Context) {
	var input models.AssignRoleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 🔑 Extract current user metadata from JWT/context
	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Scope not available"})
		return
	}
	scopeStr, ok := scope.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid scope type"})
		return
	}
	// useruserId, ok := c.Get("user_id")

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UserID not obtained from token"})
		return
	}

	userCompanyID, _ := c.Get("company_id")
	userPlantID, _ := c.Get("plant_id")

	// if input.UserId != useruserId {
	// 	c.JSON(http.StatusForbidden, gin.H{"error": "UserID not obtained from token not match with request userId"})
	// 	return
	// }

	// Step 1: Hierarchy-based role restrictions
	switch scopeStr {
	case "super_admin":
		if input.Scope == "super_admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Super admin cannot assign another super_admin"})
			return
		}

	case "company_admin":
		if input.Scope != "plant_admin" && input.Scope != "machine_user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Company admin can only assign plant_admin or machine_user"})
			return
		}
		// ✅ Ownership check → can assign only within their company
		if input.CompanyId != userCompanyID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Company admin can only assign roles within their company"})
			return
		}

	case "plant_admin":
		if input.Scope != "machine_user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Plant admin can only assign machine_user"})
			return
		}
		// ✅ Ownership check → can assign only within their plant
		if input.PlantId != userPlantID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Plant admin can only assign roles within their plant"})
			return
		}

	case "machine_user":
		c.JSON(http.StatusForbidden, gin.H{"error": "Machine user cannot assign roles"})
		return

	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Unknown scope"})
		return
	}

	// Step 2: Call service layer
	resp, err := services.AssignUserRole(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
