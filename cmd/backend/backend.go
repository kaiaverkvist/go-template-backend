package main

import (
	"github.com/kaiaverkvist/go-template-backend/internal/server"
	"github.com/kaiaverkvist/go-template-backend/internal/setup"
)

// Should initialize configuration and start relevant services.
func main() {
	setup.LoadEnv()

	instance := server.Server{}
	instance.Init()
}