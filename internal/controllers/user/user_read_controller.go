package usercontroller

import (
	userModels "basics/internal/models/user"
	userservice "basics/internal/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var input userModels.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userservice.Login(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := userModels.LoginResponse{
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


func GetAllUserInfoBasedOnScope(c *gin.Context) {

	var input userModels.AssignRoleInput

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

	input = userModels.AssignRoleInput{
		UserId:    *uId,
		Scope:     scopeStr,
		CompanyId: cId,
		PlantId:   pId,
		MachineId: mId,
	}

	users, err := userservice.GetAllUserInfoBasedOnScope(&input, limit, offset)

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
