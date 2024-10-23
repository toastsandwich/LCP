package util

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, status int, component templ.Component) error {
	buffer := templ.GetBuffer()
	defer templ.ReleaseBuffer(buffer)
	if err := component.Render(ctx.Request().Context(), buffer); err != nil {
		return err
	}
	return ctx.HTML(status, buffer.String())
}
