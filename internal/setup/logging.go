package setup

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Logging(e *echo.Echo, friendly bool) {
	// Whether we will use the easily readable
	if friendly {
		if l, ok := e.Logger.(*log.Logger); ok {
			l.SetHeader("⇨ ${time_rfc3339} ${level} <${short_file}:${line}")
		}

		e.HideBanner = true

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "⇨ ${method} ${uri} -> RESP ${status} (took ${latency_human}) (▼ ${bytes_in}B ▲ ${bytes_out}B)\n",
		}))
	} else {
		e.HideBanner = true

		e.Use(middleware.Logger())
	}
}