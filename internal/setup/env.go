package setup

import (
	"github.com/joho/godotenv"
	"github.com/kaiaverkvist/go-template-backend/internal/envs"
	"github.com/labstack/gommon/log"
	"os"
)

func LoadEnv() {
	// Attempt to load an .env
	err := godotenv.Load()
	if err != nil {
		log.Debug("No .env file has been found.")
	}

	envs.SERVER_PORT = os.Getenv("SERVER_PORT")
	envs.AUTO_TLS = os.Getenv("AUTO_TLS")
	envs.TLS_CERT_PATH = os.Getenv("TLS_CERT_PATH")
	envs.TLS_PKEY_PATH = os.Getenv("TLS_PKEY_PATH")
	envs.FRIENDLY_LOGGING = os.Getenv("FRIENDLY_LOGGING")
	envs.DOMAIN = os.Getenv("DOMAIN")
}