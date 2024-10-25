package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/util"
	"github.com/toastsandwich/LCP/view/pages"
)

type Handler struct{}

func Home(c echo.Context) error {
	return util.Render(c, http.StatusOK, pages.Home())
}

func Login(c echo.Context) error {
	return util.Render(c, http.StatusOK, pages.Login())
}

func LoginPost(c echo.Context) error {
	email := c.FormValue("emp-id")
	password := c.FormValue("password")
	return c.JSON(http.StatusOK, map[string]string{"email": email, "password": password})
}
