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
