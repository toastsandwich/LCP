package config

import (
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/db"
	"gorm.io/gorm"
)

var Echo *echo.Echo
var DB *gorm.DB

func init() {
	Echo = echo.New()
	DB = db.Init()
}
