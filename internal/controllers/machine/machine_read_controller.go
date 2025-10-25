package machinecontroller

import (
	machineModels "basics/internal/models/machine"
	machineservice "basics/internal/services/machine"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllMachinesBasedonScope(c *gin.Context) {

	var machine *machineModels.MachineInputModel

	scope, ok := c.Get("scope")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scope Not found"})
		return
	}
	companyId, _ := c.Get("company_id")
	plantId, _ := c.Get("plant_id")
	machineId, _ := c.Get("machine_id")

	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")

	companyIdStr := companyId.(string)
	plantIdStr := plantId.(string)
	machineIdStr := machineId.(string)

	companyUuid, _ := uuid.Parse(companyIdStr)
	plantUuid, _ := uuid.Parse(plantIdStr)
	machineUuid, _ := uuid.Parse(machineIdStr)
	scopeStr := scope.(string)

	limit, _ := strconv.Atoi(limitStr)
	page, _ := strconv.Atoi(pageStr)

	switch scopeStr {
	case "super_admin":
		machine = &machineModels.MachineInputModel{
			Scope: scopeStr,
		}
	case "company_admin":
		machine = &machineModels.MachineInputModel{
			Scope:     scopeStr,
			CompanyId: companyUuid,
		}
	case "plant_admin":
		machine = &machineModels.MachineInputModel{
			Scope:     scopeStr,
			CompanyId: companyUuid,
			PlantId:   plantUuid,
		}
	case "machine_user":
		machine = &machineModels.MachineInputModel{
			Scope:     scopeStr,
			CompanyId: companyUuid,
			PlantId:   plantUuid,
			MachineId: machineUuid,
		}

	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": scopeStr})
		return
	}

	machines, err := machineservice.GetAllMachinesBasedonScope(machine, limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": machines,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	},
	)
}
