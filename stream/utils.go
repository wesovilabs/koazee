package stream

import (
	"reflect"
)

func isPointer(value reflect.Value) bool {
	return value.Kind() == reflect.Ptr
}

func equalsValues(value1, value2 reflect.Value) bool {
	if isPointer(value1) {
		return equalsPtr(value1, value2)
	}
	return equalsValue(value1, value2)
}

func equalsPtr(value1, value2 reflect.Value) bool {
	return equalsValue(value1.Elem(), value2.Elem())
}

func equalsValue(value1, value2 reflect.Value) bool {
	return value1.Interface() == value2.Interface()
}
