package validators

type Employee struct {
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	EmailAddress string `validate:"required,email"`
	Phone        string `validate:"required"`
	CompanyId    int    `validate:"required"`
}

func ValidateEmployee(employee Employee) []*ErrorResponse {

	err := validate.Struct(employee)

	errors := Validate(err)

	return errors
}
