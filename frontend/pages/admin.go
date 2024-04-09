package pages

import "github.com/labstack/echo/v4"

func AdminLoginPage(ctx echo.Context) error {
	return ctx.File("static/admin/admin_login.html")
}

func AdminHomePage(ctx echo.Context) error {
	return ctx.File("static/admin/admin_home.html")
}

