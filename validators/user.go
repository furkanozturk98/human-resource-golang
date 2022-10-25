package validators

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName            string `validate:"required"`
	LastName             string `validate:"required"`
	EmailAddress         string `validate:"required,email"`
	Password             string `validate:"required"`
	PasswordConfirmation string `validate:"required"`
}

/* type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New() */

func ValidateUser(user User) []*ErrorResponse {

	var errors []*ErrorResponse
	err := validate.Struct(user)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}
