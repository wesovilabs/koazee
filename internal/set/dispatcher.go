package set

import "reflect"

type dispatchFunction func(items reflect.Value, index int, item interface{}) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   setString,
	"*string":  setPtrString,
	"bool":     setBool,
	"*bool":    setPtrBool,
	"int":      setInt,
	"*int":     setPtrInt,
	"int8":     setInt8,
	"*int8":    setPtrInt8,
	"int16":    setInt16,
	"*int16":   setPtrInt16,
	"int32":    setInt32,
	"*int32":   setPtrInt32,
	"int64":    setInt64,
	"*int64":   setPtrInt64,
	"uint":     setUint,
	"*uint":    setPtrUint,
	"uint8":    setUint8,
	"*uint8":   setPtrUint8,
	"uint16":   setUint16,
	"*uint16":  setPtrUint16,
	"uint32":   setUint32,
	"*uint32":  setPtrUint32,
	"uint64":   setUint64,
	"*uint64":  setPtrUint64,
	"float32":  setFloat32,
	"*float32": setPtrFloat32,
	"float64":  setFloat64,
	"*float64": setPtrFloat64,
}

func dispatch(items reflect.Value, itemValue interface{}, info *setInfo) (bool, interface{}) {
	input := (*info.itemType).String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, info.index, itemValue)
	}
	return false, nil
}

func setString(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	item := itemValue.(string)
	n := len(input)
	output := make([]string, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrString(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	item := itemValue.(*string)
	n := len(input)
	output := make([]*string, n)
	copy(output, input)
	output[index] = item
	return output
}

func setBool(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	item := itemValue.(bool)
	n := len(input)
	output := make([]bool, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrBool(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	item := itemValue.(*bool)
	n := len(input)
	output := make([]*bool, n)
	copy(output, input)
	output[index] = item
	return output
}

func setInt(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	item := itemValue.(int)
	n := len(input)
	output := make([]int, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrInt(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	item := itemValue.(*int)
	n := len(input)
	output := make([]*int, n)
	copy(output, input)
	output[index] = item
	return output
}

func setInt8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	item := itemValue.(int8)
	n := len(input)
	output := make([]int8, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrInt8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	item := itemValue.(*int8)
	n := len(input)
	output := make([]*int8, n)
	copy(output, input)
	output[index] = item
	return output
}

func setInt16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	item := itemValue.(int16)
	n := len(input)
	output := make([]int16, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrInt16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	item := itemValue.(*int16)
	n := len(input)
	output := make([]*int16, n)
	copy(output, input)
	output[index] = item
	return output
}

func setInt32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	item := itemValue.(int32)
	n := len(input)
	output := make([]int32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrInt32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	item := itemValue.(*int32)
	n := len(input)
	output := make([]*int32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setInt64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	item := itemValue.(int64)
	n := len(input)
	output := make([]int64, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrInt64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	item := itemValue.(*int64)
	n := len(input)
	output := make([]*int64, n)
	copy(output, input)
	output[index] = item
	return output
}

func setUint(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	item := itemValue.(uint)
	n := len(input)
	output := make([]uint, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrUint(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	item := itemValue.(*uint)
	n := len(input)
	output := make([]*uint, n)
	copy(output, input)
	output[index] = item
	return output
}

func setUint8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	item := itemValue.(uint8)
	n := len(input)
	output := make([]uint8, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrUint8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	item := itemValue.(*uint8)
	n := len(input)
	output := make([]*uint8, n)
	copy(output, input)
	output[index] = item
	return output
}

func setUint16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	item := itemValue.(uint16)
	n := len(input)
	output := make([]uint16, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrUint16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	item := itemValue.(*uint16)
	n := len(input)
	output := make([]*uint16, n)
	copy(output, input)
	output[index] = item
	return output
}

func setUint32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	item := itemValue.(uint32)
	n := len(input)
	output := make([]uint32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrUint32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	item := itemValue.(*uint32)
	n := len(input)
	output := make([]*uint32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setUint64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	item := itemValue.(uint64)
	n := len(input)
	output := make([]uint64, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrUint64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	item := itemValue.(*uint64)
	n := len(input)
	output := make([]*uint64, n)
	copy(output, input)
	output[index] = item
	return output
}

func setFloat32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	item := itemValue.(float32)
	n := len(input)
	output := make([]float32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrFloat32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	item := itemValue.(*float32)
	n := len(input)
	output := make([]*float32, n)
	copy(output, input)
	output[index] = item
	return output
}

func setFloat64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	item := itemValue.(float64)
	n := len(input)
	output := make([]float64, n)
	copy(output, input)
	output[index] = item
	return output
}

func setPtrFloat64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	item := itemValue.(*float64)
	n := len(input)
	output := make([]*float64, n)
	copy(output, input)
	output[index] = item
	return output
}
