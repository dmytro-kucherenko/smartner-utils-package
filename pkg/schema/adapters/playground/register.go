package adapter

import (
	"errors"
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

func TryRegister(validate any) error {
	validator, ok := validate.(*validator.Validate)
	if !ok {
		return errors.New("invalid validator error")
	}

	err := Register(validator)
	if err != nil {
		return fmt.Errorf("validation registration error: %v", err.Error())
	}

	return nil
}
