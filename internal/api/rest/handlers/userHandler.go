package handlers

import (
	"net/http"

	"ecommerce-app/internal/api/rest"
	"ecommerce-app/internal/dto"
	"ecommerce-app/internal/repository"
	"ecommerce-app/internal/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	svc service.UserService
}

// SetupUserRoutes registers all user-related routes
func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
	}
	userHandler := &UserHandler{
		svc: svc,
	}

	api := app.Group("/api")
	user := api.Group("/users")

	user.Post("/signup", userHandler.Signup)
	user.Post("/login", userHandler.Login)

	user.Get("/verify", userHandler.GetVerificationCode)
	user.Post("/verify", userHandler.VerifyCode)

	user.Post("/profile", userHandler.CreateProfile)
	user.Get("/profile/:id", userHandler.GetProfile)
	user.Put("/profile/:id", userHandler.UpdateProfile)

	user.Get("/cart", userHandler.FindCart)
	user.Post("/cart", userHandler.CreateCart)

	user.Post("/orders", userHandler.CreateOrder)
	user.Get("/orders", userHandler.GetOrders)
	user.Get("/orders/:id", userHandler.GetOrderById)

	user.Post("/seller", userHandler.BecomeSeller)
}

// ================= HANDLERS =================

func (h *UserHandler) Signup(ctx fiber.Ctx) error {
	var user dto.UserSignup

	if err := ctx.Bind().Body(&user); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":  "invalid request body",
			"detail": err.Error(),
		})
	}

	createdUser, err := h.svc.Signup(user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to sign up user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"user": fiber.Map{
			"id":        createdUser.ID,
			"email":     createdUser.Email,
			"user_type": createdUser.UserType,
		},
	})
}

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User logged in successfully",
	})
}

func (h *UserHandler) GetVerificationCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Verification code fetched",
	})
}

func (h *UserHandler) VerifyCode(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Verification code verified",
	})
}

func (h *UserHandler) CreateProfile(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Profile created",
	})
}

func (h *UserHandler) GetProfile(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":    "Profile fetched",
		"profile_id": id,
	})
}

func (h *UserHandler) UpdateProfile(ctx fiber.Ctx) error {
	id := ctx.Params("id")

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message":    "Profile updated",
		"profile_id": id,
	})
}

func (h *UserHandler) FindCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Cart fetched",
	})
}

func (h *UserHandler) CreateCart(ctx fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Item added to cart",
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

func (h *UserHandler) GetOrderById(ctx fiber.Ctx) error {
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
