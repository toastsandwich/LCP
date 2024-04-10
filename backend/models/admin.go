package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Admin struct {
	Username string
	Password string
}

func init() {
	db.AutoMigrate(&Admin{})
}

func AdminLogin(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	fmt.Println(username, ":", password)
	if username == "admin" && password == "admin" {
		return ctx.Redirect(http.StatusFound, "http://localhost:8080/admin/home")
	}
	return ctx.Redirect(http.StatusFound, "http://localhost:8080/admin/login")
}