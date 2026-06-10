package config

import (
	"errors"
	"os"
	"log"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn 	  string
	AppSecret string
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
    // Load DSN from env Data Source Name is a string that tells how to connect to db
	Dsn := os.Getenv("DSN")
	if len(Dsn) == 0 {
		return AppConfig{}, errors.New("DSN not set")
	}

	appSecret := os.Getenv("APP_SECRET") 
	if len(appSecret) < 1 {
		return AppConfig{}, errors.New("Appsecret not defined")
	}
	return AppConfig{
		ServerPort: httpPort,
		Dsn:        Dsn,
		AppSecret: appSecret,
	}, nil
}