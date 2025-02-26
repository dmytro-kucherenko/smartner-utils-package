package parsers

import (
	"reflect"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
	"github.com/go-playground/validator/v10"
)

func parseID(field reflect.Value) any {
	value, ok := field.Interface().(types.IDBind)
	if !ok {
		return nil
	}

	_, err := value.Parse()
	if err != nil {
		return nil
	}

	return string(value)
}

func RegisterID(validate *validator.Validate) {
	var idBind types.IDBind

	validate.RegisterCustomTypeFunc(parseID, idBind)
}
