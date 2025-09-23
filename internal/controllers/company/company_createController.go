package companycontrollers

import (
	companyModels "basics/internal/models/company"
	companyservice "basics/internal/services/company"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var input companyModels.CompanyInputModel

	// Step 1: Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Step 2: Check user scope (only super_admin can create companies)
	scope, ok := c.Get("scope")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Scope not available"})
		return
	}

	scopeStr, ok := scope.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid scope type"})
		return
	}

	if scopeStr != "super_admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only super_admin can create companies"})
		return
	}

	// Step 4: Call service
	createdCompany, err := companyservice.CreateCompany(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 5: Build response
	response := companyModels.CompanyResponseModel{
		CompanyId:   createdCompany.CompanyId,
		CompanyName: createdCompany.CompanyName,
		CreatedAt:   createdCompany.CreatedAt,
	}

	// Step 6: Return response
	c.JSON(http.StatusCreated, response)

}
