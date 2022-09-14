package config

import (
	"github.com/chocomilk/app/pkg/general/router"
	"github.com/chocomilk/app/pkg/general/validator"
	"github.com/gofiber/fiber/v2"
)

type ChocomilkGeneral struct {
	// General
}

func (s ChocomilkGeneral) InitializeRouter(app fiber.Router) {
	router.InitializeRouter(app)
}

func (s ChocomilkGeneral) InitializeValidator() {
	validator.InitializeValidator()
}

func (s ChocomilkGeneral) MakeMigrations() {

}
