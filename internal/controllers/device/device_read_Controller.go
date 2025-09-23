package devicecontrollers

import (
	userModels "basics/internal/models/user"
	deviceservice "basics/internal/services/device"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDevicesBasedOnSCope(c *gin.Context) {

	var input userModels.AssignRoleInput
	

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
	pageStr := c.DefaultQuery("page", "1")
	search := c.DefaultQuery("search", "")

	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)

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

	input = userModels.AssignRoleInput{
		UserId:    *uId,
		Scope:     scopeStr,
		CompanyId: cId,
		PlantId:   pId,
		MachineId: mId,
	}

	devices, err := deviceservice.GetAllDevicesBasedOnSCope(&input, limit, page, search)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": devices,
		"pagination": gin.H{
			"limit": limit,
			"page":  page,
		},
	})
}
