package adapter

import "github.com/go-playground/validator/v10"

var Validator *validator.Validate

func Init() error {
	validate, err := New()
	if err != nil {
		return err
	}

	Validator = validate

	return nil
}

func ValidateStruct(data any) error {
	return Validator.Struct(data)
}
