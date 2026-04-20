package handlers

import (
	"net/http"

	"ecommerce-app/internal/api/rest"
	"github.com/gofiber/fiber/v3"
)

type UserHandler struct{}

// SetupUserRoutes registers all user-related routes
func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	userHandler := &UserHandler{}

	api := app.Group("/api")
	user := api.Group("/users")

	// Public
	user.Post("/register", userHandler.Register)
	user.Post("/login", userHandler.Login)

	// Auth
	user.Get("/verify", userHandler.Verify)

	// Profile
	user.Post("/profile", userHandler.CreateProfile)
	user.Get("/profile/:id", userHandler.GetProfile)

	// Cart
	user.Post("/cart", userHandler.AddToCart)
	user.Get("/cart", userHandler.GetCart)

	// Orders
	user.Post("/orders", userHandler.CreateOrder)
	user.Get("/orders", userHandler.GetOrders)
	user.Get("/orders/:id", userHandler.GetOrder)

	// Seller
	user.Post("/seller", userHandler.BecomeSeller)
}

// ================= HANDLERS =================

func (h *UserHandler) Register(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
	})
}

func (h *UserHandler) Verify(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User verified",
	})
}

func (h *UserHandler) CreateProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User profile created",
	})
}

func (h *UserHandler) GetProfile(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":    "Profile fetched",
		"profile_id": id,
	})
}

func (h *UserHandler) AddToCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item added to cart",
	})
}

func (h *UserHandler) GetCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Cart fetched",
	})
}

func (h *UserHandler) CreateOrder(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Order created",
	})
}

func (h *UserHandler) GetOrders(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Orders fetched",
	})
}

func (h *UserHandler) GetOrder(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Order fetched",
		"order_id": id,
	})
}

func (h *UserHandler) BecomeSeller(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Seller application submitted",
	})
}