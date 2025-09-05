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

		protectedRoutes.GET("/profile", controllers.UserProfile)
		protectedRoutes.POST("/createcompany", controllers.CreateCompany)
		protectedRoutes.POST("/createplant", controllers.CreatePlant)
		protectedRoutes.POST("/createmachine", controllers.CreateMachine)
		protectedRoutes.POST("/createdevice", controllers.CreateDevice)
		protectedRoutes.POST("/createdevicedata", controllers.CreateDeviceData)
		protectedRoutes.POST("/assignuserrole", controllers.AssignUserRole)
		protectedRoutes.GET("/getallusers", controllers.GetAllUserInfoBasedOnScope)

	}

	return r
}
