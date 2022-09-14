package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/chocomilk/app/pkg/general/service"
)

func InitializeRouter(app fiber.Router) {
	// api := app.Group("/")
	app.Get("/", service.Welcome)
}
