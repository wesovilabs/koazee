
package add

import "reflect"


type dispatchFunction func(items reflect.Value, item interface{}) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":  addString,
	"*string":  addPtrString,
	"bool":  addBool,
	"*bool":  addPtrBool,
	"int":  addInt,
	"*int":  addPtrInt,
	"int8":  addInt8,
	"*int8":  addPtrInt8,
	"int16":  addInt16,
	"*int16":  addPtrInt16,
	"int32":  addInt32,
	"*int32":  addPtrInt32,
	"int64":  addInt64,
	"*int64":  addPtrInt64,
	"uint":  addUint,
	"*uint":  addPtrUint,
	"uint8":  addUint8,
	"*uint8":  addPtrUint8,
	"uint16":  addUint16,
	"*uint16":  addPtrUint16,
	"uint32":  addUint32,
	"*uint32":  addPtrUint32,
	"uint64":  addUint64,
	"*uint64":  addPtrUint64,
	"float32":  addFloat32,
	"*float32":  addPtrFloat32,
	"float64":  addFloat64,
	"*float64":  addPtrFloat64,
}

func dispatch(items reflect.Value, itemValue interface{}, info *addInfo) (bool,interface{}) {
	input:=(*info.itemType).String()
	if fnVal,ok:=dispatcher[input];ok{
		return true, fnVal(items,itemValue)
    }
	return false, nil
}

func addString(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	item := itemValue.(string)
	n := len(input)
	output := make([]string, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrString(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	item := itemValue.(*string)
	n := len(input)
	output := make([]*string, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	item := itemValue.(bool)
	n := len(input)
	output := make([]bool, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	item := itemValue.(*bool)
	n := len(input)
	output := make([]*bool, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	item := itemValue.(int)
	n := len(input)
	output := make([]int, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	item := itemValue.(*int)
	n := len(input)
	output := make([]*int, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	item := itemValue.(int8)
	n := len(input)
	output := make([]int8, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	item := itemValue.(*int8)
	n := len(input)
	output := make([]*int8, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	item := itemValue.(int16)
	n := len(input)
	output := make([]int16, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	item := itemValue.(*int16)
	n := len(input)
	output := make([]*int16, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	item := itemValue.(int32)
	n := len(input)
	output := make([]int32, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	item := itemValue.(*int32)
	n := len(input)
	output := make([]*int32, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	item := itemValue.(int64)
	n := len(input)
	output := make([]int64, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	item := itemValue.(*int64)
	n := len(input)
	output := make([]*int64, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	item := itemValue.(uint)
	n := len(input)
	output := make([]uint, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	item := itemValue.(*uint)
	n := len(input)
	output := make([]*uint, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	item := itemValue.(uint8)
	n := len(input)
	output := make([]uint8, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	item := itemValue.(*uint8)
	n := len(input)
	output := make([]*uint8, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	item := itemValue.(uint16)
	n := len(input)
	output := make([]uint16, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	item := itemValue.(*uint16)
	n := len(input)
	output := make([]*uint16, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	item := itemValue.(uint32)
	n := len(input)
	output := make([]uint32, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	item := itemValue.(*uint32)
	n := len(input)
	output := make([]*uint32, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	item := itemValue.(uint64)
	n := len(input)
	output := make([]uint64, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	item := itemValue.(*uint64)
	n := len(input)
	output := make([]*uint64, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	item := itemValue.(float32)
	n := len(input)
	output := make([]float32, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	item := itemValue.(*float32)
	n := len(input)
	output := make([]*float32, n+1)
	copy(output, input)
	output[n] = item
	return output
}

func addFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	item := itemValue.(float64)
	n := len(input)
	output := make([]float64, n+1)
	copy(output,input)
	output[n] = item
	return output
}

func addPtrFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	item := itemValue.(*float64)
	n := len(input)
	output := make([]*float64, n+1)
	copy(output, input)
	output[n] = item
	return output
}
