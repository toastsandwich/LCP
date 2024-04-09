package server

import (
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/frontend/pages"
)

func Start(addr string) {
	app := echo.New()
	app.Static("/static", "static")
	app.Static("/assets", "assets")
	routes(app)
	app.Logger.Fatal(app.Start(addr))
}

func routes(app *echo.Echo) {
	app.GET("/", pages.HomePage)
	app.GET("/home", pages.HomePage)
	app.GET("/admin/login", pages.AdminLoginPage)
	app.GET("/doctor/login", pages.DoctorLoginPage)
	app.GET("/lab-assistant/login", pages.LabAssisantLogin)

	app.GET("/admin/home", pages.AdminHomePage)
	app.GET("/doctor/home", pages.DoctorHomePage)
	app.GET("/lab-assistant/home", pages.LabAssistantHome)

	app.GET("/doctor/patient/:id", pages.PatientInfoPage)
	app.GET("/lab-assistant/patient/:id", pages.CreatePatient)
	app.GET("/lab-assistant/patient/:id", pages.LabAssistantSeePatient)

	app.GET("/admin/register/doctor/", pages.RegisterDoctorPage)
	app.GET("/admin/register/labassistant/", pages.RegisterLabAssistant)
}
