//doctor
//search and patient see patient

// lab assitant
// CRUD

// admin
// C doc and LA

package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/models"
)

func Init(echo *echo.Echo) {
	echo.GET("/getAllDoctors", models.GetAllDoctors)
	echo.GET("/getDoctorByID/:id", models.GetDoctorByID)
	echo.POST("/addDoctor/", models.AddDoctor)
	echo.DELETE("/deleteDoctor/:id", models.DeleteDoctor)

	echo.GET("/getAllLabAssistants", models.GetAllLabAssistants)
	echo.GET("getLabAssistantByID/:id", models.GetLabAssistantByID)
	echo.POST("/addLabAssistant/", models.AddLabAssistant)
	echo.DELETE("/deleteLabAssistant/:id", models.DeleteLabAssistant)
}
