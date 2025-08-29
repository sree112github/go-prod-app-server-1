package routes

import (
	"basics/internal/controllers"
	middleware "basics/internal/middlewere"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/signup", controllers.SignUp)
		api.POST("login", controllers.Login)

	}

	protectedRoutes := r.Group("/api")

	protectedRoutes.Use(middleware.AuthMiddleware())

	{

		protectedRoutes.GET("/profile", middleware.AuthMiddleware(), controllers.UserProfile)
		protectedRoutes.POST("/createcompany", middleware.AuthMiddleware(), controllers.CreateCompany)
		protectedRoutes.POST("/createplant", middleware.AuthMiddleware(), controllers.CreatePlant)
	}

	return r
}
