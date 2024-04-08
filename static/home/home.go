package home

import "github.com/labstack/echo/v4"

func Home(ctx echo.Context) error {
	return ctx.File("static/home/index.html")
}
