package validators

type Company struct {
	Name    string `validate:"required"`
	Address string `validate:"required"`
	Phone   string `validate:"required"`
	Email   string `validate:"required,email"`
	Logo    string
	Website string `validate:"required,url"`
}

/* type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New() */

func ValidateCompany(company Company) []*ErrorResponse {

	err := validate.Struct(company)

	errors := Validate(err)

	return errors
}
