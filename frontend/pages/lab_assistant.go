package pages

import "github.com/labstack/echo/v4"

func LabAssisantLogin(ctx echo.Context) error {
	return ctx.File("static/labassistant/Labassistant_login.html")
}

func LabAssistantHome(ctx echo.Context) error {
	return ctx.File("static/labassistant/lab_assistant_home.html")
}

func LabAssistantAddPatient(ctx echo.Context) error {
	return ctx.File("static/labassisant/create_patient.html")
}

func RegisterLabAssistant(ctx echo.Context) error {
	return ctx.File("static/labassistant/Register_lab_assitant.html")
}

func CreatePatient(ctx echo.Context) error {
	return ctx.File("static/labassistant/create_patient.html")
}

func LabAssistantSeePatient(ctx echo.Context) error {
	return ctx.File("static/labassistant/lab_ass_view_pateint.html")
}
