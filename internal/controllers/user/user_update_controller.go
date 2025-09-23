package usercontroller

import (
	userModels "basics/internal/models/user"
	userservice "basics/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignUserRole(c *gin.Context) {
	var input userModels.AssignRoleInput

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
	companyId, _ := userCompanyID.(string)
	plantId, _ := userPlantID.(string)

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
		// if input.CompanyId != &companyId {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Company admin can only assign roles within their company"})
		// 	return
		// }

		input = userModels.AssignRoleInput{
			UserId:    input.UserId,
			Scope:     input.Scope,
			CompanyId: &companyId,
			PlantId:   input.PlantId,
		}

	case "plant_admin":
		if input.Scope != "machine_user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Plant admin can only assign machine_user"})
			return
		}
		// // ✅ Ownership check → can assign only within their plant
		// if input.PlantId != userPlantID {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Plant admin can only assign roles within their plant"})
		// 	return
		// }

		input = userModels.AssignRoleInput{
			UserId:    input.UserId,
			Scope:     input.Scope,
			CompanyId: &companyId,
			PlantId:   &plantId,
			MachineId: input.MachineId,
		}
	case "machine_user":
		c.JSON(http.StatusForbidden, gin.H{"error": "Machine user cannot assign roles"})
		return

	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Unknown scope"})
		return
	}

	// Step 2: Call service layer
	resp, err := userservice.AssignUserRole(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
