package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUserInfoBasedOnScope(c *gin.Context) {

	var input models.AssignRoleInput

	// if err := c.ShouldBindJSON(&input); err != nil {

	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scope not found"})
		return
	}

	scopeStr := scope.(string)

	// Extract IDs from JWT or context
	companyID, _ := c.Get("company_id")
	plantID, _ := c.Get("plant_id")
	machineId, _ := c.Get("machine_id")
	userID, _ := c.Get("user_id")

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	var uId, cId, pId, mId *string

	if companyID != nil {
		val := companyID.(string)
		cId = &val
	}
	if plantID != nil {
		val := plantID.(string)
		pId = &val
	}
	if userID != nil {
		val := userID.(string)
		uId = &val
	}

	if machineId != nil {
		val := machineId.(string)
		mId = &val
	}

	input = models.AssignRoleInput{
		UserId:    *uId,
		Scope:     scopeStr,
		CompanyId: cId,
		PlantId:   pId,
		MachineId: mId,
	}

	users, err := services.GetAllUserInfoBasedOnScope(&input, limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
		"pagination": gin.H{
			"limit":  limit,
			"offset": offset,
		},
	})

}
