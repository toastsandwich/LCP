//doctor
//search and patient see patient

// lab assitant
// CRUD

// admin
// C doc and LA

package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/LCP/models"
)

func Init(echo *echo.Echo) {
	// echo.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
	echo.GET("/getAllDoctors", models.GetAllDoctors)
	echo.GET("/getDoctorByID/:id", models.GetDoctorByID)
	echo.POST("/addDoctor/", models.AddDoctor, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteDoctor/:id", models.DeleteDoctor)

	echo.GET("/getAllLabAssistants", models.GetAllLabAssistants)
	echo.GET("getLabAssistantByID/:id", models.GetLabAssistantByID)
	echo.POST("/addLabAssistant/", models.AddLabAssistant, middleware.BodyLimit("1M"))
	echo.DELETE("/deleteLabAssistant/:id", models.DeleteLabAssistant)
}
