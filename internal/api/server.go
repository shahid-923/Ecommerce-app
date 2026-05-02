package api

import (
	"log"

	"ecommerce-app/config"
	"ecommerce-app/internal/api/rest"
	"ecommerce-app/internal/api/rest/handlers"
	"ecommerce-app/internal/domain"

	"github.com/gofiber/fiber/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer(config config.AppConfig) {
	log.Printf("Starting server")

	app := fiber.New()
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})


	if err != nil {
		log.Fatalf("Database connection error: %v\n", err )
	}
    
	log.Println("Database connected successfully")


	//run the migration
	db.AutoMigrate(&domain.User{})

	rh := &rest.RestHandler{
		App: app,
		DB: db,
	}

	setupRoutes(rh)

	app.Listen(":" + config.ServerPort)
}


func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoutes(rh)
}