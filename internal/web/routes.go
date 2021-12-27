package web

import (
	"github.com/kaiaverkvist/go-template-backend/internal/view"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		v := view.New(c, "index")
		err := v.Render()
		if err != nil {
			return c.String(200, "Unable to render view")
		}
		return nil
	})
}