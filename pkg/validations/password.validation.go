package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validation(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	matchUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	matchLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	matchDigit := regexp.MustCompile(`\d`).MatchString(password)
	matchLength := regexp.MustCompile(`^.{8,}$`).MatchString(password)

	return matchUpper && matchLower && matchDigit && matchLength
}

func RegisterPassword(validate *validator.Validate, tag Tags) error {
	return validate.RegisterValidation(string(tag), validation)
}
