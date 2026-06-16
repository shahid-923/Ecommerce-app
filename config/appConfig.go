package config

import (
	"errors"
	"os"
    "log"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
	AppSecret  string

	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPassword string
	SMTPFrom     string
}

func SetupEnvironment() (AppConfig, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env not loaded")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("HTTP_PORT not set")
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		return AppConfig{}, errors.New("DSN not set")
	}

	appSecret := os.Getenv("APP_SECRET")
	if appSecret == "" {
		return AppConfig{}, errors.New("APP_SECRET not set")
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpFrom := os.Getenv("SMTP_FROM")

	if smtpHost == "" {
		return AppConfig{}, errors.New("SMTP_HOST not set")
	}

	if smtpPort == "" {
		return AppConfig{}, errors.New("SMTP_PORT not set")
	}

	if smtpUser == "" {
		return AppConfig{}, errors.New("SMTP_USER not set")
	}

	if smtpPassword == "" {
		return AppConfig{}, errors.New("SMTP_PASSWORD not set")
	}

	if smtpFrom == "" {
		return AppConfig{}, errors.New("SMTP_FROM not set")
	}

	return AppConfig{
		ServerPort:   httpPort,
		Dsn:          dsn,
		AppSecret:    appSecret,
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUser:     smtpUser,
		SMTPPassword: smtpPassword,
		SMTPFrom:     smtpFrom,
	}, nil
}