package adapter

import (
	"fmt"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground/parsers"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground/validations"
	"github.com/go-playground/validator/v10"
)

type Tags string

const (
	TagPassword Tags = "password"
)

func Register(validate *validator.Validate) error {
	parsers.RegisterID(validate)
	err := validations.RegisterPassword(validate, string(TagPassword))
	if err != nil {
		return err
	}

	return nil
}

func TryRegister(validate any) {
	validator, ok := validate.(*validator.Validate)
	if !ok {
		panic("Invalid validator error")
	}

	err := Register(validator)
	if err != nil {
		panic(fmt.Sprintln("Validation registration error:", err.Error()))
	}
}
