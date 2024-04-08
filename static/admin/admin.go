package admin

import "github.com/labstack/echo/v4"

func AdminLogin(ctx echo.Context) error {
	return ctx.File("static/admin/admin_login.html")
}

func AdminPage(ctx echo.Context) error {
	return ctx.File("static/admin/admin_home.html")
}
