package server

import (
	"fmt"
	"github.com/kaiaverkvist/go-template-backend/internal/envs"
	"github.com/kaiaverkvist/go-template-backend/internal/setup"
	"github.com/kaiaverkvist/go-template-backend/internal/web"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func (s *Server) Init() {
	//database.InitDatabase()
	//database.AutoMigrate(&models.Model{})

	s.e = echo.New()
	setup.Logging(s.e, envs.FRIENDLY_LOGGING == "true")

	web.AddMiddlewares(s.e)
	web.AddRoutes(s.e)

	address := fmt.Sprintf(":%s", envs.SERVER_PORT)
	setup.Start(s.e, address, envs.AUTO_TLS == "true", envs.TLS_CERT_PATH, envs.TLS_PKEY_PATH)
}