package middleware

import "github.com/go-playground/validator/v10"

// CustomValidator is a wrapper around validator
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validates the struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
