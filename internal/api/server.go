package api

import (
	"log"

	"ecommerce-app/config"
	"ecommerce-app/internal/api/rest"            
	"ecommerce-app/internal/api/rest/handlers"

	"github.com/gofiber/fiber/v3"
)
func StartServer(config config.AppConfig) {
	log.Printf("Starting server")

	app := fiber.New()

	// 👇 PASTE HERE
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Server is working")
	})

	rh := &rest.RestHandler{App: app}

	setupRoutes(rh)

	app.Listen(":" + config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}