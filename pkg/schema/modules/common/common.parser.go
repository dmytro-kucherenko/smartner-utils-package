package common

import (
	"reflect"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/utils"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

func ModifySchema[T any](data T) any {
	return utils.ModifySchema(data, func(nodeValue reflect.Value) (newType reflect.Value, isChanged bool) {
		if nodeValue.Type() == reflect.TypeOf((*types.ID)(nil)).Elem() {
			return reflect.New(reflect.TypeOf((types.IDBind)(""))).Elem(), true
		}

		return
	})
}

func ParseSchema[T any](schema any) T {
	return utils.ParseSchema[T](schema, func(nodeValue reflect.Value) (newType reflect.Value, isChanged bool) {
		if nodeValue.Type() == reflect.TypeOf((*types.IDBind)(nil)).Elem() {
			idBind := nodeValue.Interface().(types.IDBind)
			id, _ := idBind.Parse()

			value := reflect.New(reflect.TypeOf((types.ID)(id))).Elem()
			value.Set(reflect.ValueOf(id))

			return value, true
		}

		return
	})
}
