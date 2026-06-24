package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func ErrorMessage(ctx fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func InternalError(ctx fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": err.Error(),
	})
}

func SuccessResponse(ctx fiber.Ctx, msg string, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}