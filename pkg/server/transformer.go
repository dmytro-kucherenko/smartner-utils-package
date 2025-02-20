package server

import (
	"reflect"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

func TransformDataToSchema[T any](data T) any {
	dataType := reflect.TypeOf(data)
	fields := []reflect.StructField{}

	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)

		if field.Type == reflect.TypeOf((*types.ID)(nil)).Elem() {
			fields = append(fields, reflect.StructField{
				Name: field.Name,
				Type: reflect.TypeOf((*types.IDBind)(nil)).Elem(),
				Tag:  field.Tag,
			})
		} else {
			fields = append(fields, field)
		}
	}

	return reflect.New(reflect.StructOf(fields)).Interface()
}

func TransformSchemaToData[T any](schema any) T {
	schemaValue := reflect.ValueOf(schema).Elem()
	dataType := reflect.TypeOf((*T)(nil)).Elem()
	data := reflect.New(dataType).Elem()

	for i := 0; i < dataType.NumField(); i++ {
		dataField := dataType.Field(i)
		schemaField := schemaValue.Field(i)

		if dataField.Type == reflect.TypeOf((*types.ID)(nil)).Elem() {
			bind := schemaField.Interface().(types.IDBind)
			value, _ := bind.Parse()

			data.Field(i).Set(reflect.ValueOf(value))
		} else {
			data.Field(i).Set(schemaField)
		}
	}

	return data.Interface().(T)
}
