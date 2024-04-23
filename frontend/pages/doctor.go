package pages

import (
	"github.com/labstack/echo/v4"
)

func DoctorHomePage(ctx echo.Context) error {
	return ctx.File("static/doctor/doctor_home.html")
}

func DoctorLoginPage(ctx echo.Context) error {
	return ctx.File("static/doctor/doctor_login.html")
}

func PatientInfoPage(ctx echo.Context) error {
	return ctx.File("static/doctor/display_patient_doctor.html")
}

func RegisterDoctorPage(ctx echo.Context) error {
	return ctx.File("static/doctor/Register_doctor.html")
}
