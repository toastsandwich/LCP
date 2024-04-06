package app

import (
	"github.com/toastsandwich/LCP/config"
	"github.com/toastsandwich/LCP/routes"
)

func Start() {
	echo := config.Echo
	routes.Init(echo)
	echo.Logger.Fatal(echo.Start(":8080"))
}
