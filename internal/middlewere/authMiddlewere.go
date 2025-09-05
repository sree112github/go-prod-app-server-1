package middleware

import (
	"basics/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract token (expecting "Bearer <token>")
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyJwtToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Save user_id in Gin context for use in handlers
		c.Set("email", claims["email"])
		c.Set("user_id", claims["user_id"])
		c.Set("scope", claims["scope"])
		c.Set("company_id", claims["company_id"])
		c.Set("plant_id", claims["plant_id"])
		c.Set("machine_id", claims["machine_id"])

		c.Next()
	}
}
