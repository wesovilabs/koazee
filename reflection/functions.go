package reflection

import (
	"fmt"
	"reflect"
)

type argArgOutSignature func(*reflect.Value, interface{}) interface{}

var functionsArgArgOuts = make(map[string]argArgOutSignature)

func init() {
	functionsArgArgOuts["func(int, int) int"] = IntIntToInt
}

// GetFunctionArgArgOut returns the functions to be called when fn has arg,arg and out
func GetFunctionArgArgOut(fn interface{}) func(*reflect.Value, interface{}) interface{} {
	return functionsArgArgOuts[reflect.TypeOf(fn).String()]
}

// IntIntToInt function that iterates over function like this func(int int)int
func IntIntToInt(itemsValue *reflect.Value, fn interface{}) interface{} {
	fn2 := fn.(func(int, int) int)
	acc := int(0)
	for i := 0; i < itemsValue.Len(); i++ {
		val := Int(itemsValue.Index(i))
		acc = fn2(acc, val)
	}
	fmt.Println(reflect.TypeOf(acc).Kind())
	return acc
}
