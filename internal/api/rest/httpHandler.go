package rest

import (
	"ecommerce-app/internal/helper"
	"ecommerce-app/config"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type RestHandler struct {
	App *fiber.App
	DB *gorm.DB
	Auth helper.Auth
	Config config.AppConfig
}