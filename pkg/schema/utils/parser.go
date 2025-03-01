package utils

import (
	"reflect"
)

func ModifySchema[T any](data T, handle func(nodeType reflect.Value) (reflect.Value, bool)) any {
	value := reflect.ValueOf(data)
	newValue := BuildValue(value, handle)

	return newValue.Addr().Interface()
}

func ParseSchema[R any](data any, handle func(nodeType reflect.Value) (reflect.Value, bool)) R {
	value := reflect.ValueOf(data).Elem()
	filler := BuildValue(value, handle)

	typed := reflect.New(reflect.TypeOf((*R)(nil)).Elem()).Elem()
	FillValue(typed, filler)

	return typed.Interface().(R)
}
