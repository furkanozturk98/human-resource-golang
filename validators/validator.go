package validators

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func Validate(err error) []*ErrorResponse {

	var errors []*ErrorResponse

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
