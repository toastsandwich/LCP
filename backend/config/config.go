package config

import (
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/db"
	"gorm.io/gorm"
)

var Echo *echo.Echo
var DB *gorm.DB
var SessionManager *scs.SessionManager

func init() {
	Echo = echo.New()
	DB = db.Init()
	SessionManager = scs.New()
	SessionManager.Lifetime = 24 * time.Hour
}
