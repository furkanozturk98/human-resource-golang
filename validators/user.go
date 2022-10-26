package validators

type User struct {
	FirstName            string `validate:"required"`
	LastName             string `validate:"required"`
	Email                string `validate:"required,email"`
	Password             string `validate:"required"`
	PasswordConfirmation string `validate:"required"`
}

func ValidateUser(user User) []*ErrorResponse {

	err := validate.Struct(user)

	errors := Validate(err)

	if user.Password != user.PasswordConfirmation {
		var element ErrorResponse
		element.FailedField = "Password"
		element.Tag = "notConfirmed"
		element.Value = user.Password
		errors = append(errors, &element)
	}

	return errors
}
