package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/toastsandwich/LCP/util"
	"github.com/toastsandwich/LCP/view/layout"
)

func Test(ctx echo.Context) error {
	return util.Render(ctx, 200, layout.Base("test"))
}
