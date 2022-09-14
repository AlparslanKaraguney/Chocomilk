package validator

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func InitializeValidator() {

	Validate.RegisterValidation("is_valid_source", func(fl validator.FieldLevel) bool {
		source := fl.Field().String()

		switch source {
		case "turkish", "foreign", "turkish_foreign":
			return true
		default:
			return false
		}
	})

}
