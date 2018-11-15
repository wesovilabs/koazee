package reflection

import (
	"reflect"
)

// IntIntToInt function that iterates over function like this func(int int)int
func IntIntToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int, int) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// StringIntToString function that iterates over function like this func(string int)string
func StringIntToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string, int) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolIntToBool function that iterates over function like this func(string int)string
func BoolIntToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool, int) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float32IntToFloat32 function that iterates over function like this func(string int)string
func Float32IntToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32, int) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float64IntToFloat64 function that iterates over function like this func(string int)string
func Float64IntToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float64, int) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// StringStringToString function that iterates over function like this func(int int)int
func IntStringToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int, string) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// StringStringToString function that iterates over function like this func(string int)string
func StringStringToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string, string) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolStringToBool function that iterates over function like this func(string int)string
func BoolStringToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool, string) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float32StringToFloat32 function that iterates over function like this func(string int)string
func Float32StringToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32, string) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float64StringToFloat64 function that iterates over function like this func(string int)string
func Float64StringToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float64, string) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float32Float32ToFloat32 function that iterates over function like this func(int int)int
func IntFloat32ToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int, float32) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float32Float32ToFloat32 function that iterates over function like this func(string int)string
func StringFloat32ToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string, float32) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolFloat32ToBool function that iterates over function like this func(string int)string
func BoolFloat32ToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool, float32) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float32Float32ToFloat32 function that iterates over function like this func(string int)string
func Float32Float32ToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32, float32) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float64Float32ToFloat64 function that iterates over function like this func(string int)string
func Float64Float32ToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float64, float32) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolBoolToBool function that iterates over function like this func(int int)int
func IntBoolToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int, bool) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolBoolToBool function that iterates over function like this func(string int)string
func StringBoolToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string, bool) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolBoolToBool function that iterates over function like this func(string int)string
func BoolBoolToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool, bool) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// BoolBoolToBool function that iterates over function like this func(string int)string
func Float32BoolToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32, bool) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}

// Float64BoolToFloat64 function that iterates over function like this func(string int)string
func Float64BoolToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float64, bool) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	return acc
}
