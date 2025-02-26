package utils

import "reflect"

func ChangeSchema[T any](data T, callback func(nodeType reflect.Value) (reflect.Value, bool)) any {
	value := reflect.ValueOf(data)
	newValue := BuildValue(value, callback)

	return newValue.Addr().Interface()
}

func ParseSchema[R any](data any, callback func(nodeType reflect.Value) (reflect.Value, bool)) R {
	value := reflect.ValueOf(data).Elem()
	filler := BuildValue(value, callback)

	typed := reflect.New(reflect.TypeOf((*R)(nil)).Elem()).Elem()
	FillValue(typed, filler)

	return typed.Interface().(R)
}
