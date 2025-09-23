package plantcontroller

import (
	plantModels "basics/internal/models/plant"
	plantservice "basics/internal/services/plant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllPlantsBasedOnScope(c *gin.Context) {

	var input plantModels.PlantInputModel

	scope, ok := c.Get("scope")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Scope is invalid"})
		return
	}

	scopeStr, _ := scope.(string)

	limit := c.DefaultQuery("limit", "10")
	page := c.DefaultQuery("page", "1")

	limitStr, _ := strconv.Atoi(limit)
	pageStr, _ := strconv.Atoi(page)

	companyId, ok := c.Get("company_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company id not found"})
		return
	}
	companyIdStr, _ := companyId.(string)
	companyUuid, _ := uuid.Parse(companyIdStr)

	plantId, ok := c.Get("plant_id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "plant id not found"})
		return
	}
	plantIdStr, _ := plantId.(string)

	plantUuid, _ := uuid.Parse(plantIdStr)

	machineId, ok := c.Get("machine_id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "machine id not found"})
		return
	}
	machineIdStr, _ := machineId.(string)

	machineUuid, _ := uuid.Parse(machineIdStr)

	switch scopeStr {

	case "super_admin":
		input = plantModels.PlantInputModel{
			Scope: scopeStr,
		}
	case "company_admin":

		input = plantModels.PlantInputModel{
			CompanyId: companyUuid,
			Scope:     scopeStr,
		}
	case "plant_admin":

		input = plantModels.PlantInputModel{
			CompanyId: companyUuid,
			PlantId:   plantUuid,
			Scope:     scopeStr,
		}

	case "machine_user":
		input = plantModels.PlantInputModel{
			CompanyId: companyUuid,
			PlantId:   plantUuid,
			MachineId: machineUuid,
			Scope:     scopeStr,
		}

	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "Unknown scope"})
		return
	}

	plants, err := plantservice.GetAllPlantsBasedOnScope(&input, limitStr, pageStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": plants,
		"pagination": gin.H{
			"limit":  limit,
			"offset": pageStr,
		},
	})

}
