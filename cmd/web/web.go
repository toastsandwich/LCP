package web

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/cmd/handler"
)

func Start() error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", handler.Test)

	return e.Start(":15001")
}
