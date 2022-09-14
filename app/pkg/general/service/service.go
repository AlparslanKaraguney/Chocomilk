package service

import (
	"github.com/chocomilk/app/model"
	"github.com/chocomilk/app/pkg/general/serializer"
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	parameters := new(serializer.WelcomeRequest)
	err := parameters.Validate(c)
	if err != nil {
		return model.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	// var asd []string
	// temp := asd[4]
	// fmt.Println(temp)
	return c.SendString("Welcome to Chocomilk! " + parameters.Name)
}
