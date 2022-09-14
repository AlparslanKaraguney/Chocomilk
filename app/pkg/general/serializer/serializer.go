package serializer

import (
	"github.com/chocomilk/app/pkg/general/validator"
	"github.com/gofiber/fiber/v2"
)

type WelcomeRequest struct {
	Name string `json:"name" validate:"required"`
}

func (s *WelcomeRequest) Validate(c *fiber.Ctx) error {

	err := c.QueryParser(s)
	if err != nil {
		return err
	}

	err = validator.Validate.Struct(s)
	if err != nil {
		return err
	}

	return nil
}
