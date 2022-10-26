package validators

import "mime/multipart"

type Company struct {
	Name    string               `validate:"required"`
	Address string               `validate:"required"`
	Phone   string               `validate:"required"`
	Email   string               `validate:"required,email"`
	Logo    multipart.FileHeader `form:"image" validate:"required"`
	Website string               `validate:"required,url"`
}

func ValidateCompany(company Company) []*ErrorResponse {

	err := validate.Struct(company)

	errors := Validate(err)

	return errors
}
