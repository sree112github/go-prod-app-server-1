package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDeviceData(c *gin.Context) {
	var input models.DeviceDataInputModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scope not found"})
		return
	}

	scopeStr, ok := scope.(string)
	if !ok || scopeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scope String Not Avialble"})
		return
	}

	if scopeStr != "super_admin" && scopeStr != "company_admin" && scopeStr != "plant_admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only super,company,or plant admin can Add data"})
		return
	}

	response, err := services.CreateDeviceService(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)

}
