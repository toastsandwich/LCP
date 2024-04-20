package app

import (
	"github.com/toastsandwich/LCP/config"
	"github.com/toastsandwich/LCP/handlers"
)

func Start(port string) {
	echo := config.Echo
	handlers.Init(echo)
	echo.Logger.Fatal(echo.Start(port))
}
