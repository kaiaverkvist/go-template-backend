package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AddMiddlewares(e *echo.Echo) {
	// Panics shouldn't kill the server.
	e.Use(middleware.Recover())

	e.Use(middleware.Secure())
}