
package pop

import (
	"reflect"
)


type dispatchFunction func(items reflect.Value) (interface{}, interface{}) 

var dispatcher = map[string]dispatchFunction{
	"string":  popString,
	"*string":  popStringPtr,
	"bool":  popBool,
	"*bool":  popBoolPtr,
	"int":  popInt,
	"*int":  popIntPtr,
	"int8":  popInt8,
	"*int8":  popInt8Ptr,
	"int16":  popInt16,
	"*int16":  popInt16Ptr,
	"int32":  popInt32,
	"*int32":  popInt32Ptr,
	"int64":  popInt64,
	"*int64":  popInt64Ptr,
	"uint":  popUint,
	"*uint":  popUintPtr,
	"uint8":  popUint8,
	"*uint8":  popUint8Ptr,
	"uint16":  popUint16,
	"*uint16":  popUint16Ptr,
	"uint32":  popUint32,
	"*uint32":  popUint32Ptr,
	"uint64":  popUint64,
	"*uint64":  popUint64Ptr,
	"float32":  popFloat32,
	"*float32":  popFloat32Ptr,
	"float64":  popFloat64,
	"*float64":  popFloat64Ptr,
}

func dispatch(items reflect.Value, itemsType reflect.Type) (bool,interface{},interface{}) {
	if fnVal, ok := dispatcher[itemsType.String()]; ok {
		first, tail := fnVal(items)
		return true, first, tail
	}
	return false,nil,nil
}
func popString(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]string)
	return input[0],input[1:]
}

func popStringPtr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*string)
	return input[0],input[1:]
}
func popBool(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]bool)
	return input[0],input[1:]
}

func popBoolPtr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*bool)
	return input[0],input[1:]
}
func popInt(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]int)
	return input[0],input[1:]
}

func popIntPtr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*int)
	return input[0],input[1:]
}
func popInt8(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]int8)
	return input[0],input[1:]
}

func popInt8Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*int8)
	return input[0],input[1:]
}
func popInt16(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]int16)
	return input[0],input[1:]
}

func popInt16Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*int16)
	return input[0],input[1:]
}
func popInt32(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]int32)
	return input[0],input[1:]
}

func popInt32Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*int32)
	return input[0],input[1:]
}
func popInt64(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]int64)
	return input[0],input[1:]
}

func popInt64Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*int64)
	return input[0],input[1:]
}
func popUint(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]uint)
	return input[0],input[1:]
}

func popUintPtr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*uint)
	return input[0],input[1:]
}
func popUint8(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]uint8)
	return input[0],input[1:]
}

func popUint8Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*uint8)
	return input[0],input[1:]
}
func popUint16(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]uint16)
	return input[0],input[1:]
}

func popUint16Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*uint16)
	return input[0],input[1:]
}
func popUint32(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]uint32)
	return input[0],input[1:]
}

func popUint32Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*uint32)
	return input[0],input[1:]
}
func popUint64(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]uint64)
	return input[0],input[1:]
}

func popUint64Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*uint64)
	return input[0],input[1:]
}
func popFloat32(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]float32)
	return input[0],input[1:]
}

func popFloat32Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*float32)
	return input[0],input[1:]
}
func popFloat64(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]float64)
	return input[0],input[1:]
}

func popFloat64Ptr(itemsValue reflect.Value) (interface{},interface{}){
	input := itemsValue.Interface().([]*float64)
	return input[0],input[1:]
}
