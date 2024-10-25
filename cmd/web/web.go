package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/toastsandwich/LCP/cmd/handler"
)

func Start() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	e.Static("/assets", "assets")

	e.GET("/", handler.Home)
	e.GET("/login", handler.Login)

	e.POST("/login", handler.LoginPost)
	return e.Start(":15001")
}
