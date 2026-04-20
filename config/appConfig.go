package config

import (
	"errors"
	"os"
	"log"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
}

func SetupEnvironment() (AppConfig, error) {
	// Load .env first
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env not loaded")
	}

	// Now env variables are available
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("HTTP_PORT not set")
	}

	return AppConfig{
		ServerPort: httpPort,
	}, nil
}