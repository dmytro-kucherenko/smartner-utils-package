package common

import (
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/utils"
	"github.com/go-playground/validator/v10"
)

func DecodeStruct[T any](data any) (result T, err error) {
	schema := ModifySchema(result)
	err = utils.DecodeStruct(data, schema)
	if err != nil {
		return
	}

	validate := validator.New()
	err = adapter.Register(validate)
	if err != nil {
		return
	}

	err = validate.Struct(schema)
	if err != nil {
		return
	}

	result = ParseSchema[T](schema)

	return
}

func EncodeStruct(data any) (result map[string]any, err error) {
	return utils.EncodeStruct(data)
}
