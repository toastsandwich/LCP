package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/LCP/models"
)

func Init(echo *echo.Echo) {
	echo.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	echo.Use(middleware.CORS())
	// routes for CRUD Doctor
	echo.GET("/getAllDoctors", models.GetAllDoctors)
	echo.GET("/getDoctorByID/:id", models.GetDoctorByID)
	echo.POST("/addDoctor/", models.AddDoctor, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteDoctor/:id", models.DeleteDoctor)
	// routes for CRUD Lab_Assistant
	echo.GET("/getAllLabAssistants", models.GetAllLabAssistants)
	echo.GET("/getLabAssistantByID/:id", models.GetLabAssistantByID)
	echo.POST("/addLabAssistant/", models.AddLabAssistant, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteLabAssistant/:id", models.DeleteLabAssistant)
	//routes for CRUD Patient
	echo.GET("/getAllPatients", models.GetAllPatients)
	echo.GET("/getPatientByID/:id", models.GetPatientByID)
	echo.POST("/createPatient/", models.CreatePatient)
	echo.DELETE("/deletePatient/:id", models.DeletePatient)

	echo.POST("/adminLogin", models.AdminLogin)
}
