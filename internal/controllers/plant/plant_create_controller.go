package plantcontroller

import (
	plantModels "basics/internal/models/plant"
	plantservice "basics/internal/services/plant"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePlant(c *gin.Context) {

	var input plantModels.PlantInputModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Body"})
		return
	}

	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scope is Invalid"})
		return
	}

	scopeStr, ok := scope.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid scope type"})
		return
	}

	companyId, ok := c.Get("company_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company id not found"})
		return
	}
	companyIdStr, ok := companyId.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Companyid"})
		return
	}


	companyUUID, err := uuid.Parse(companyIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Company ID format"})
		return
	}

	if scopeStr != "super_admin" && scopeStr != "company_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only super_admin or Company admin  can create Plants"})
		return
	}

	if scopeStr == "super_admin" {
		input = plantModels.PlantInputModel{
			CompanyId: input.CompanyId,
			PlantName: input.PlantName,
		}
	}

	if scopeStr == "company_admin" {
		input = plantModels.PlantInputModel{
			CompanyId: companyUUID,
			PlantName: input.PlantName,
		}
	}

	response, err := plantservice.CreatePlant(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)

}
