
package reverse

import (
	"reflect"
)


type dispatchFunction func(items reflect.Value)interface{} 

var dispatcher = map[string]dispatchFunction{
	"string":  reverseString,
	"*string":  reverseStringPtr,
	"bool":  reverseBool,
	"*bool":  reverseBoolPtr,
	"int":  reverseInt,
	"*int":  reverseIntPtr,
	"int8":  reverseInt8,
	"*int8":  reverseInt8Ptr,
	"int16":  reverseInt16,
	"*int16":  reverseInt16Ptr,
	"int32":  reverseInt32,
	"*int32":  reverseInt32Ptr,
	"int64":  reverseInt64,
	"*int64":  reverseInt64Ptr,
	"uint":  reverseUint,
	"*uint":  reverseUintPtr,
	"uint8":  reverseUint8,
	"*uint8":  reverseUint8Ptr,
	"uint16":  reverseUint16,
	"*uint16":  reverseUint16Ptr,
	"uint32":  reverseUint32,
	"*uint32":  reverseUint32Ptr,
	"uint64":  reverseUint64,
	"*uint64":  reverseUint64Ptr,
	"float32":  reverseFloat32,
	"*float32":  reverseFloat32Ptr,
	"float64":  reverseFloat64,
	"*float64":  reverseFloat64Ptr,
}

func dispatch(items reflect.Value, itemsType reflect.Type) (bool,interface{}) {
	if fnVal, ok := dispatcher[itemsType.String()]; ok {
		return true,fnVal(items)
	}
	return false,nil
}
func reverseString(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]string)
	len := len(input)
	output := make([]string, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseStringPtr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*string)
	len := len(input)
	output := make([]*string, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseBool(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]bool)
	len := len(input)
	output := make([]bool, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseBoolPtr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*bool)
	len := len(input)
	output := make([]*bool, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseInt(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]int)
	len := len(input)
	output := make([]int, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseIntPtr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*int)
	len := len(input)
	output := make([]*int, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseInt8(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]int8)
	len := len(input)
	output := make([]int8, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseInt8Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*int8)
	len := len(input)
	output := make([]*int8, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseInt16(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]int16)
	len := len(input)
	output := make([]int16, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseInt16Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*int16)
	len := len(input)
	output := make([]*int16, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseInt32(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]int32)
	len := len(input)
	output := make([]int32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseInt32Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*int32)
	len := len(input)
	output := make([]*int32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseInt64(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]int64)
	len := len(input)
	output := make([]int64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseInt64Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*int64)
	len := len(input)
	output := make([]*int64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseUint(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]uint)
	len := len(input)
	output := make([]uint, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseUintPtr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*uint)
	len := len(input)
	output := make([]*uint, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseUint8(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]uint8)
	len := len(input)
	output := make([]uint8, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseUint8Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*uint8)
	len := len(input)
	output := make([]*uint8, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseUint16(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]uint16)
	len := len(input)
	output := make([]uint16, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseUint16Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*uint16)
	len := len(input)
	output := make([]*uint16, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseUint32(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]uint32)
	len := len(input)
	output := make([]uint32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseUint32Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*uint32)
	len := len(input)
	output := make([]*uint32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseUint64(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]uint64)
	len := len(input)
	output := make([]uint64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseUint64Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*uint64)
	len := len(input)
	output := make([]*uint64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseFloat32(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]float32)
	len := len(input)
	output := make([]float32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseFloat32Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*float32)
	len := len(input)
	output := make([]*float32, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
func reverseFloat64(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]float64)
	len := len(input)
	output := make([]float64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}

func reverseFloat64Ptr(itemsValue reflect.Value) interface{}{
	input := itemsValue.Interface().([]*float64)
	len := len(input)
	output := make([]*float64, len)
	for index := 0; index < (len/2)+1; index++ {
		output[index], output[len-1-index] = input[len-1-index], input[index]
	}
	return output
}
