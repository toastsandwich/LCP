package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/LCP/models"
	"github.com/toastsandwich/LCP/repository/controllers"
)

func Init(echo *echo.Echo) {
	echo.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	echo.Use(middleware.CORS())
	// routes for CRUD Doctor
	echo.GET("/getAllDoctors", controllers.GetAllDoctors)
	echo.GET("/getDoctorByID/:id", controllers.GetDoctorByID)
	echo.POST("/addDoctor/", controllers.AddDoctor, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteDoctor/:id", controllers.DeleteDoctor)
	// routes for CRUD Lab_Assistant
	echo.GET("/getAllLabAssistants", controllers.GetAllLabAssistants)
	echo.GET("/getLabAssistantByID/:id", controllers.GetLabAssistantByID)
	echo.POST("/addLabAssistant/", controllers.AddLabAssistant, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteLabAssistant/:id", controllers.DeleteLabAssistant)
	//routes for CRUD Patient
	echo.GET("/getAllPatients", controllers.GetAllPatients)
	echo.GET("/getPatientByID/:id", controllers.GetPatientByID)
	echo.POST("/createPatient/", controllers.CreatePatient)
	echo.DELETE("/deletePatient/:id", controllers.DeletePatient)

	echo.POST("/adminLogin", models.AdminLogin)
}
