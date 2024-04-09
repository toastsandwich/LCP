package pages

import "github.com/labstack/echo/v4"

func HomePage(ctx echo.Context) error {
	return ctx.File("static/home/index.html")
}
