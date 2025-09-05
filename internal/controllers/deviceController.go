package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateDevice(c *gin.Context) {
	var input models.DeviceInputModel

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request Body"})
		return
	}

	if input.MachineId == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid machine_id"})
		return
	}

	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No scope Found"})
		return
	}

	scopeStr, ok := scope.(string)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No scope string Conversion failed"})
		return
	}

	if scopeStr != "super_admin" && scopeStr != "company_admin" && scopeStr != "plant_admin" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only the super,company,plant can Create a device"})
		return
	}

	response, err := services.CreateDevice(&input)
	if err != nil {
		fmt.Println("❌ Error creating device:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)

}
