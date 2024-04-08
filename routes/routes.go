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
	"github.com/toastsandwich/LCP/static/admin"
	"github.com/toastsandwich/LCP/static/doctor"
	"github.com/toastsandwich/LCP/static/home"
)

func Init(echo *echo.Echo) {
	echo.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))
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

	// routes for frontend
	echo.GET("/", home.Home)
	echo.POST("/admin-login", admin.AdminLogin)
	echo.GET("/admin-home", admin.AdminPage)

	echo.GET("/doctor-login", doctor.LoginPage)
	echo.GET("/doctor-home", doctor.HomePage)
}
