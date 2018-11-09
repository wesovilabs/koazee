package stream

import (
	"reflect"
)

func emptySlice(arrayType reflect.Type) interface{} {
	return slice(arrayType, 0, 0)
}

func slice(arrayType reflect.Type, len, cap int) interface{} {
	slice := reflect.MakeSlice(reflect.SliceOf(arrayType), len, cap)
	return slice.Interface().(interface{})
}

func elementAt(s *stream, index int) interface{} {
	v := itemsToArray(s)
	return v.Index(index).Interface()
}

func elementLast(s *stream) interface{} {
	v := itemsToArray(s)
	return v.Index(v.Len() - 1).Interface()
}

func itemsToArray(s *stream) reflect.Value {
	return reflect.ValueOf(s.items)
}

func isPointer(value reflect.Value) bool {
	return value.Kind() == reflect.Ptr
}

func eqyalsValues(value1, value2 reflect.Value) bool {
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
