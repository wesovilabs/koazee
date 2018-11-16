package reflection

import "reflect"

// IntToInt function that iterates over function like this func(int int)int
func IntToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// IntToString function that iterates over function like this func(string int)string
func IntToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// IntToBool function that iterates over function like this func(string int)string
func IntToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// IntToFloat32 function that iterates over function like this func(string int)string
func IntToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// IntToFloat64 function that iterates over function like this func(string int)string
func IntToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// StringToInt function that iterates over function like this func(int int)int
func StringToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// StringToString function that iterates over function like this func(string int)string
func StringToString(input []string, fn interface{}) []string {
	c := make(map[string]string)
	fn2 := fn.(func(string) string)
	for i := 0; i < len(input); i++ {
		if val, ok := c[input[i]]; ok {
			input[1] = val
			continue
		}
		input[1] = fn2(input[i])
	}
	return input
}

// StringToBool function that iterates over function like this func(string int)string
func StringToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// StringToFloat32 function that iterates over function like this func(string int)string
func StringToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// StringToFloat64 function that iterates over function like this func(string int)string
func StringToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(string) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := String(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// Float32ToFloat32 function that iterates over function like this func(int int)int
func Float32ToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// Float32ToFloat32 function that iterates over function like this func(string int)string
func Float32ToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// Float32ToBool function that iterates over function like this func(string int)string
func Float32ToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// Float32ToFloat32 function that iterates over function like this func(string int)string
func Float32ToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// Float32ToFloat64 function that iterates over function like this func(string int)string
func Float32ToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(float32) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Float32(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// BoolToInt function that iterates over function like this func(int int)int
func BoolToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// BoolToString function that iterates over function like this func(string int)string
func BoolToString(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool) string)
	acc := ""
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// BoolToBool function that iterates over function like this func(string int)string
func BoolToBool(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool) bool)
	acc := false
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// BoolToFloat32 function that iterates over function like this func(string int)string
func BoolToFloat32(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool) float32)
	acc := float32(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}

// BoolToFloat64 function that iterates over function like this func(string int)string
func BoolToFloat64(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(bool) float64)
	acc := float64(0.0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Bool(itemsValue.Index(i))
		fn2(val)
	}
	return acc
}
