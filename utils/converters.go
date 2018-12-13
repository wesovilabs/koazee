package utils

import "reflect"

// Bool converts value to bool
func Bool(value reflect.Value) bool {
	return value.Interface().(bool)
}

// String converts value to string
func String(value reflect.Value) string {
	return value.Interface().(string)
}

// Float32 converts value to float32
func Float32(value reflect.Value) float32 {
	return value.Interface().(float32)
}

// Float64 converts value to Float64
func Float64(value reflect.Value) float64 {
	return value.Interface().(float64)
}

// Int converts value to int
func Int(value reflect.Value) int {
	return value.Interface().(int)
}

// Int8 converts value to int8
func Int8(value reflect.Value) int8 {
	return value.Interface().(int8)
}

// Int16 converts value to int16
func Int16(value reflect.Value) int16 {
	return value.Interface().(int16)
}

// Int32 converts value to int32
func Int32(value reflect.Value) int32 {
	return value.Interface().(int32)
}

// Int64 converts value to int64
func Int64(value reflect.Value) int64 {
	return value.Int()
}
