package companycontrollers

import (
	userModels "basics/internal/models/user"
	companyservice "basics/internal/services/company"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCompanyBasedOnScope(c *gin.Context) {

	var input userModels.AssignRoleInput

	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "scope not found"})
		return
	}

	scopeStr := scope.(string)
	userID, _ := c.Get("user_id")
	companyID, _ := c.Get("company_id")
	plantID, _ := c.Get("plant_id")
	machineID, _ := c.Get("machine_id")

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something wrong in limit"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something wrong in page"})
		return
	}

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

	if machineID != nil {
		val := machineID.(string)
		mId = &val
	}

	input = userModels.AssignRoleInput{
		UserId:    *uId,
		Scope:     scopeStr,
		CompanyId: cId,
		PlantId:   pId,
		MachineId: mId,
	}

	companies, err := companyservice.GetAllCompanyBasedOnScope(&input, limit, page)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": companies,
		"pagination": gin.H{
			"limit": limit,
			"page":  page,
		},
	})

}
