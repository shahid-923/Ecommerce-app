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
	svc *service.UserService
}

// SetupUserRoutes registers all user-related routes
func SetupUserRoutes(rh *rest.RestHandler) {
	app := rh.App

	svc := &service.UserService{
		Repo: repository.NewUserRepository(rh.DB),
		Auth: rh.Auth,
	}

	userHandler := &UserHandler{
		svc: svc,
	}

	// ================= PUBLIC ROUTES =================

	app.Post("/signup", userHandler.Signup)
	app.Post("/login", userHandler.Login)

	// ================= PROTECTED ROUTES =================

	protected := app.Group("/")
	protected.Use(rh.Auth.Authorize())

	protected.Get("/verify", userHandler.GetVerificationCode)
	protected.Post("/verify", userHandler.VerifyCode)

	protected.Post("/profile", userHandler.CreateProfile)
	protected.Get("/profile/:id", userHandler.GetProfile)
	protected.Put("/profile/:id", userHandler.UpdateProfile)

	protected.Get("/cart", userHandler.FindCart)
	protected.Post("/cart", userHandler.CreateCart)

	protected.Post("/orders", userHandler.CreateOrder)
	protected.Get("/orders", userHandler.GetOrders)
	protected.Get("/orders/:id", userHandler.GetOrderById)

	protected.Post("/seller", userHandler.BecomeSeller)
}
// ================= SIGNUP =================

func (h *UserHandler) Signup(ctx fiber.Ctx) error {
	var input dto.UserSignup

	if err := ctx.Bind().JSON(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}

	_, token, err := h.svc.Signup(input)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"token":   token,
	})
}

// ================= LOGIN =================

func (h *UserHandler) Login(ctx fiber.Ctx) error {
	var input dto.UserLogin

	if err := ctx.Bind().JSON(&input); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}

	token, err := h.svc.Login(input.Email, input.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "login successful",
		"token":   token,
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
