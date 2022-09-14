package model

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func ErrorResponse(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(
		Response{
			Message: msg,
			Data:    []any{},
			Success: false,
		})
}

func SuccessResponse(c *fiber.Ctx, status int, msg string, data interface{}) error {
	return c.Status(status).JSON(
		Response{
			Message: msg,
			Data:    data,
			Success: true,
		})
}
