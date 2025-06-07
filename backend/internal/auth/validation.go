package auth

import (
	"errors"
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

var (
	ErrInvalidEmail = errors.New("EMAIL HAS AN INVALID FORMAT")
	ErrWeakPassword = errors.New("PASSWORD IS TOO WEAK")
)

// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()

	if err := validate.RegisterValidation("email", validateEmail); err != nil {
		panic("failed to register email validator: " + err.Error())
	}

	if err := validate.RegisterValidation("password", validatePassword); err != nil {
		panic("failed to register password validator: " + err.Error())
	}

	return &Validation{validate}
}

func validateEmail(fl validator.FieldLevel) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

	return emailRegex.MatchString(fl.Field().String())
}

func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	if len(password) < 8 {
		return false
	}

	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasDigit   = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#$%^&*_\-]`).MatchString
	)

	return hasUpper(password) && hasLower(password) && hasDigit(password) && hasSpecial(password)
}

func (v *Validation) Validate(input interface{}) error {
	err := v.validate.Struct(input)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, fieldErr := range err.(validator.ValidationErrors) {
			switch fieldErr.Field() {
			case "Email":
				return ErrInvalidEmail
			case "Password":
				return ErrWeakPassword
			}
		}

		return errors.New("validation failed")
	}

	return nil
}
