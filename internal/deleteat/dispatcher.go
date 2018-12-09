
package deleteat

import "reflect"


type dispatchFunction func(items reflect.Value,itemsLen int,  index int) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":  deleteAtString,
	"*string":  deleteAtPtrString,
	"bool":  deleteAtBool,
	"*bool":  deleteAtPtrBool,
	"int":  deleteAtInt,
	"*int":  deleteAtPtrInt,
	"int8":  deleteAtInt8,
	"*int8":  deleteAtPtrInt8,
	"int16":  deleteAtInt16,
	"*int16":  deleteAtPtrInt16,
	"int32":  deleteAtInt32,
	"*int32":  deleteAtPtrInt32,
	"int64":  deleteAtInt64,
	"*int64":  deleteAtPtrInt64,
	"uint":  deleteAtUint,
	"*uint":  deleteAtPtrUint,
	"uint8":  deleteAtUint8,
	"*uint8":  deleteAtPtrUint8,
	"uint16":  deleteAtUint16,
	"*uint16":  deleteAtPtrUint16,
	"uint32":  deleteAtUint32,
	"*uint32":  deleteAtPtrUint32,
	"uint64":  deleteAtUint64,
	"*uint64":  deleteAtPtrUint64,
	"float32":  deleteAtFloat32,
	"*float32":  deleteAtPtrFloat32,
	"float64":  deleteAtFloat64,
	"*float64":  deleteAtPtrFloat64,
}

func dispatch(items reflect.Value, itemsType *reflect.Type,itemsLen int,  index int) (bool,interface{}) {
	input:=(*itemsType).String()
	if fnVal,ok:=dispatcher[input];ok{
		return true, fnVal(items,itemsLen,index)
    }
	return false, nil
}

func deleteAtString(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]string)
	output := make([]string, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrString(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make([]*string, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtBool(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make([]bool, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrBool(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make([]*bool, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtInt(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]int)
	output := make([]int, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrInt(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make([]*int, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtInt8(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make([]int8, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrInt8(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make([]*int8, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtInt16(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make([]int16, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrInt16(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make([]*int16, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtInt32(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make([]int32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrInt32(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make([]*int32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtInt64(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make([]int64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrInt64(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make([]*int64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtUint(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make([]uint, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrUint(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make([]*uint, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtUint8(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make([]uint8, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrUint8(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make([]*uint8, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtUint16(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make([]uint16, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrUint16(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make([]*uint16, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtUint32(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make([]uint32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrUint32(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make([]*uint32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtUint64(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make([]uint64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrUint64(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make([]*uint64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtFloat32(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make([]float32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrFloat32(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make([]*float32, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}

func deleteAtFloat64(itemsValue reflect.Value,itemsLen int, index int) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make([]float64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output 
}

func deleteAtPtrFloat64(itemsValue reflect.Value,itemsLen int,  index int) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make([]*float64, itemsLen-1)
	copy(output, input[:index])
	copy(output[index:], input[index+1:])
	return output
}
