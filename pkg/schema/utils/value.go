package utils

import "reflect"

func BuildValue(nodeValue reflect.Value, handle func(nodeValue reflect.Value) (reflect.Value, bool)) reflect.Value {
	value, isChanged := handle(nodeValue)
	if isChanged {
		return value
	}

	switch nodeValue.Kind() {
	case reflect.Struct:
		fields := []reflect.StructField{}
		values := map[string]reflect.Value{}

		for i := 0; i < nodeValue.NumField(); i++ {
			field := nodeValue.Type().Field(i)
			value := BuildValue(nodeValue.Field(i), handle)

			values[field.Name] = value
			fields = append(fields, reflect.StructField{
				Name: field.Name,
				Type: value.Type(),
				Tag:  field.Tag,
			})
		}

		newValue := reflect.New(reflect.StructOf(fields)).Elem()
		for i := 0; i < newValue.NumField(); i++ {
			fieldName := newValue.Type().Field(i).Name

			if value, ok := values[fieldName]; ok {
				newValue.Field(i).Set(value)
			}
		}

		return newValue
	case reflect.Slice:
		elemValue := reflect.New(nodeValue.Type().Elem()).Elem()
		newValue := BuildValue(elemValue, handle)
		sliceType := reflect.SliceOf(newValue.Type())

		newSlice := reflect.New(sliceType).Elem()
		for i := 0; i < nodeValue.Len(); i++ {
			elemValue := nodeValue.Index(i)
			newValue := BuildValue(elemValue, handle)

			newSlice = reflect.Append(newSlice, newValue)
		}

		return newSlice
	case reflect.Array:
		elemValue := reflect.New(nodeValue.Type().Elem()).Elem()
		newValue := BuildValue(elemValue, handle)
		arrayType := reflect.ArrayOf(nodeValue.Len(), newValue.Type())
		newArray := reflect.New(arrayType).Elem()

		for i := 0; i < nodeValue.Len(); i++ {
			elemValue := nodeValue.Index(i)
			newValue := BuildValue(elemValue, handle)

			newArray.Index(i).Set(newValue)
		}

		return newArray
	case reflect.Map:
		keyValue := reflect.New(nodeValue.Type().Key()).Elem()
		elemValue := reflect.New(nodeValue.Type().Elem()).Elem()
		if nodeValue.IsNil() {
			return reflect.New(reflect.MapOf(keyValue.Type(), elemValue.Type())).Elem()
		}

		mapValue := reflect.MakeMap(reflect.MapOf(keyValue.Type(), elemValue.Type()))
		for _, key := range nodeValue.MapKeys() {
			elem := nodeValue.MapIndex(key)

			newKey := BuildValue(key, handle)
			newElem := BuildValue(elem, handle)

			mapValue.SetMapIndex(newKey, newElem)
		}

		return mapValue
	case reflect.Ptr:
		if nodeValue.IsNil() {
			value := BuildValue(reflect.New(nodeValue.Type().Elem()).Elem(), handle)

			return reflect.Zero(reflect.PointerTo(value.Type()))
		}

		value := BuildValue(nodeValue.Elem(), handle)
		pointer := reflect.New(value.Type())
		pointer.Elem().Set(value)

		return pointer
	default:
		return nodeValue
	}
}

func FillValue(nodeValue reflect.Value, fillerValue reflect.Value) reflect.Value {
	if nodeValue.Type() == fillerValue.Type() {
		nodeValue.Set(fillerValue)

		return nodeValue
	}

	switch nodeValue.Kind() {
	case reflect.Struct:
		for i := 0; i < nodeValue.NumField(); i++ {
			fieldName := nodeValue.Type().Field(i).Name
			nodeField := nodeValue.FieldByName(fieldName)
			fillerField := fillerValue.FieldByName(fieldName)

			nodeField.Set(FillValue(nodeField, fillerField))
		}

		return nodeValue
	case reflect.Array:
	case reflect.Slice:
		for i := 0; i < fillerValue.Len(); i++ {
			nodeElem := nodeValue.Index(i)
			fillerElem := fillerValue.Index(i)

			nodeValue = reflect.Append(nodeValue, FillValue(nodeElem, fillerElem))
		}

		return nodeValue
	case reflect.Map:
		for _, key := range fillerValue.MapKeys() {
			nodeElem := nodeValue.MapIndex(key)
			fillerElem := fillerValue.MapIndex(key)

			nodeValue.SetMapIndex(key, FillValue(nodeElem, fillerElem))
		}

		return nodeValue
	case reflect.Ptr:
		nodeElem := nodeValue.Elem()
		fillerElem := fillerValue.Elem()

		nodeValue.Elem().Set(FillValue(nodeElem, fillerElem))

		return nodeValue
	}

	return nodeValue
}
