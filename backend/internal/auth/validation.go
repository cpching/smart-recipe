package auth

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password too weak")
)

// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("email", validateEmail)
	validate.RegisterValidation("password", validatePassword)
	return &Validation{validate}
}

func validateEmail(fl validator.FieldLevel) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	// emailRegex := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

	if emailRegex.MatchString(fl.Field().String()) {
		return true
	}

	return false
}

func validatePassword(fl validator.FieldLevel) bool {

	return false
}

func (v *Validation) Validate(input interface{}) error {
	err := v.validate.Struct(input)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			if fieldErr.Tag() == "email" {
				return ErrInvalidEmail
			}
			if fieldErr.Tag() == "password" {
				return ErrWeakPassword
			}
		}

		return errors.New("validation failed")
	}

	return nil
}
