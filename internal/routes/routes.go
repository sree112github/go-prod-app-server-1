package routes

import (
	"basics/internal/controllers"
	companycontrollers "basics/internal/controllers/company"
	devicecontrollers "basics/internal/controllers/device"
	devicedatacontroller "basics/internal/controllers/device_data"
	machinecontroller "basics/internal/controllers/machine"
	plantcontroller "basics/internal/controllers/plant"
	usercontroller "basics/internal/controllers/user"
	middleware "basics/internal/middlewere"
	_ "time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// ✅ Enable CORS
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"}, // change "*" to your Flutter web URL in production
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	api := r.Group("/api")
	api.Use(middleware.RateLimitMiddleware())

	// api.Use(middleware.CORSMiddleware())

	{
		//Auth End points

		api.POST("/signup", usercontroller.SignUp)
		api.POST("/login", usercontroller.Login)

	}

	protectedRoutes := r.Group("/api")

	protectedRoutes.Use(middleware.RateLimitMiddleware(), middleware.AuthMiddleware())

	{

		//User endpoints
		protectedRoutes.GET("/profile", controllers.UserProfile)
		protectedRoutes.POST("/assignuserrole", usercontroller.AssignUserRole)
		protectedRoutes.GET("/getallusers", usercontroller.GetAllUserInfoBasedOnScope) // ?page=2&limit=10

		//Company endpoints
		protectedRoutes.POST("/createcompany", companycontrollers.CreateCompany)
		protectedRoutes.GET("/getallcompanies", companycontrollers.GetAllCompanyBasedOnScope) //// ?page=2&limit=10

		//plant endpoints
		protectedRoutes.POST("/createplant", plantcontroller.CreatePlant)
		protectedRoutes.GET("/getallplants", plantcontroller.GetAllPlantsBasedOnScope) //// ?page=2&limit=10

		//Machine Endpoints
		protectedRoutes.POST("/createmachine", machinecontroller.CreateMachine)
		protectedRoutes.GET("/getallmachines", machinecontroller.GetAllMachinesBasedonScope)

		//Device Endpoints
		protectedRoutes.POST("/createdevice", devicecontrollers.CreateDevice)
		protectedRoutes.GET("/getalldevices", devicecontrollers.GetAllDevicesBasedOnSCope) // ?page=2&limit=10

		//Device Data enpoints
		protectedRoutes.POST("/createdevicedata", devicedatacontroller.CreateDeviceData)

	}

	return r
}
