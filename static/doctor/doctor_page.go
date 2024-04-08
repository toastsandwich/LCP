package doctor

import "github.com/labstack/echo/v4"

func HomePage(ctx echo.Context) error {
	return ctx.File("/static/doctor/doctor_home.html")
}

func LoginPage(ctx echo.Context) error {
	return ctx.File("/static/doctor/doctor_login.html")
}
