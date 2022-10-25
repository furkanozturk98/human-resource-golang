package messages

import "errors"

var (
	NO_RECORD_FOUND  = errors.New("Record not found")
	USER_NOT_CREATED = errors.New("User not created")
	USER_NOT_UPDATED = errors.New("User not updated")
)
