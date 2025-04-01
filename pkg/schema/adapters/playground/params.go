package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
)

type paramsValidator func(any) error

func (validate paramsValidator) Make(data any) error {
	return validate(data)
}

func NewParamsValidator() (server.ParamsValidator, error) {
	validate, err := New()
	if err != nil {
		return nil, err
	}

	return paramsValidator(validate.Struct), nil
}
