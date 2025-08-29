package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePlant(c *gin.Context) {

	var input models.PlantInputModel

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
	if scopeStr != "super_admin" && scopeStr != "company_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only super_admin or Company admin  can create Plants"})
		return
	}

	plant := &models.PlantModel{
		CompanyId: input.CompanyId,
		PlantName: input.PlantName,
	}

	createPlant, err := services.CreatePlant(plant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.PlantResponseModel{
		CompanyId: createPlant.CompanyId,
		PlantId:   createPlant.PlantId,
		PlantName: createPlant.PlantName,
		CreatedAt: createPlant.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)

}
