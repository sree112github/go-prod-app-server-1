package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Here the credential extract from the midddle were
func UserProfile(c *gin.Context) {

	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	userScope, ok := c.Get("scope")

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Scope not found in token"})
		return
	}

	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email not found in token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"email": email, "user_id": userId, "user_scope": userScope})
}
