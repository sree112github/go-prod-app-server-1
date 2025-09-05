package controllers

import (
	"basics/internal/models"
	"basics/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMachine(c *gin.Context) {
	var input models.MachineInputModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
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
	if scopeStr != "super_admin" && scopeStr != "company_admin" && scopeStr != "plant_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only super_admin or Company admin  can create Plants"})
		return
	}

	machine := models.MachineInputModel{
		PlantId:     input.PlantId,
		MachineName: input.MachineName,
	}

	createMachine, err := services.CreateMachine(&machine)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.MachineResponseModel{
		MachineId:   createMachine.MachineId,
		PlantId:     createMachine.PlantId,
		MachineName: createMachine.MachineName,
		CreatedAt:   createMachine.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
