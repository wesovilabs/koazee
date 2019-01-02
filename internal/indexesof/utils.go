package indexesof

import "reflect"

func equalsValues(value1, value2 reflect.Value) bool {
	if value1.Kind() == reflect.Ptr && value2.Kind() == reflect.Ptr {
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
