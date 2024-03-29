package http

// import (
// 	"strings"
// 	"unicode"

// 	"github.com/go-playground/validator/v10"

// 	"myapp/pkg/constants"
// 	ce "myapp/pkg/errors"
// 	"myapp/pkg/helpers"
// )

// type UserValidator struct {
// 	validate *validator.Validate
// }

// func NewUserValidator() *UserValidator {
// 	uv := &UserValidator{
// 		validate: validator.New(),
// 	}
// 	uv.RegisterCustomValidation()

// 	return uv
// }

// func (uv *UserValidator) RegisterCustomValidation() *UserValidator {
// 	_ = uv.validate.RegisterValidation("password", uv.ValidatePassword)
// 	_ = uv.validate.RegisterValidation("phone_number", uv.ValidatePhoneNumber)
// 	_ = uv.validate.RegisterValidation("phone_code", uv.ValidatePhoneCode)

// 	return uv
// }

// func (uv *UserValidator) ValidatePassword(fl validator.FieldLevel) bool {
// 	var (
// 		letter, special, number, upper, exclude bool
// 	)
// 	p := fl.Field().String()

// 	if len(p) < 8 {
// 		return false
// 	}
// 	specialChar := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
// 	for _, c := range p {
// 		switch {
// 		case unicode.IsNumber(c):
// 			number = true
// 		case unicode.IsUpper(c):
// 			upper = true
// 		case unicode.IsSpace(c) || helpers.CheckStringInSlice(strings.Split(specialChar, ""), string(c)):
// 			special = true
// 		case unicode.IsLetter(c):
// 			letter = true
// 		case (unicode.IsPunct(c) || unicode.IsSymbol(c)) && !helpers.CheckStringInSlice(strings.Split(specialChar, ""), string(c)):
// 			exclude = true
// 		}
// 	}

// 	return !exclude && letter && number && upper && special
// }

// func (uv *UserValidator) ValidatePhoneNumber(fl validator.FieldLevel) bool {
// 	p := fl.Field().String()
// 	if len(p) < 8 || len(p) > 10 {
// 		return false
// 	}

// 	return true
// }

// func (uv *UserValidator) ValidatePhoneCode(fl validator.FieldLevel) bool {
// 	p := fl.Field().String()
// 	_, ok := constants.CountryCodeMapping[p]
// 	return ok
// }

// func (uv *UserValidator) GetError(fe validator.FieldError) *ce.Error {
// 	switch fe.Tag() {
// 	case "required":
// 		f := strings.ToLower(fe.Field())
// 		switch f {
// 		case "phonenumber":
// 			return ce.New(ce.ErrorPhoneNumberRequired, "", nil)
// 		case "phonecode":
// 			return ce.New(ce.ErrorPhoneCodeRequired, "", nil)
// 		case "email":
// 			return ce.New(ce.ErrorEmailRequired, "", nil)
// 		case "password":
// 			return ce.New(ce.ErrorPasswordRequired, "", nil)
// 		default:
// 			return ce.New(ce.ErrorMissingRequiredField, "", nil)
// 		}
// 	case "email":
// 		return ce.New(ce.ErrorInvalidEmail, "invalid email", nil)
// 	case "password":
// 		return ce.New(ce.ErrorInvalidPassword, "invalid password", nil)
// 	case "phone_number":
// 		return ce.New(ce.ErrorInvalidPhoneNumber, "invalid phone_number", nil)
// 	case "phone_code":
// 		return ce.New(ce.ErrorInvalidPhoneCode, "invalid phone_code", nil)
// 	}

// 	return nil
// }
