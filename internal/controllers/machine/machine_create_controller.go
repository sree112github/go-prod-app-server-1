package machinecontroller

import (
	machineModels "basics/internal/models/machine"
	machineservice "basics/internal/services/machine"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMachine(c *gin.Context) {
	var input machineModels.MachineInputModel
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

	machine := machineModels.MachineInputModel{
		PlantId:     input.PlantId,
		MachineName: input.MachineName,
	}

	createMachine, err := machineservice.CreateMachine(&machine)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := machineModels.MachineResponseModel{
		MachineId:   createMachine.MachineId,
		PlantId:     createMachine.PlantId,
		MachineName: createMachine.MachineName,
		CreatedAt:   createMachine.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
