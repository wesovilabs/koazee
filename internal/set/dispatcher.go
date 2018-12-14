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
	input[index] = item
	return input
}

func setPtrString(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	item := itemValue.(*string)
	input[index] = item
	return input
}

func setBool(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	item := itemValue.(bool)
	input[index] = item
	return input
}

func setPtrBool(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	item := itemValue.(*bool)
	input[index] = item
	return input
}

func setInt(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	item := itemValue.(int)
	input[index] = item
	return input
}

func setPtrInt(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	item := itemValue.(*int)
	input[index] = item
	return input
}

func setInt8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	item := itemValue.(int8)
	input[index] = item
	return input
}

func setPtrInt8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	item := itemValue.(*int8)
	input[index] = item
	return input
}

func setInt16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	item := itemValue.(int16)
	input[index] = item
	return input
}

func setPtrInt16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	item := itemValue.(*int16)
	input[index] = item
	return input
}

func setInt32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	item := itemValue.(int32)
	input[index] = item
	return input
}

func setPtrInt32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	item := itemValue.(*int32)
	input[index] = item
	return input
}

func setInt64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	item := itemValue.(int64)
	input[index] = item
	return input
}

func setPtrInt64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	item := itemValue.(*int64)
	input[index] = item
	return input
}

func setUint(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	item := itemValue.(uint)
	input[index] = item
	return input
}

func setPtrUint(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	item := itemValue.(*uint)
	input[index] = item
	return input
}

func setUint8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	item := itemValue.(uint8)
	input[index] = item
	return input
}

func setPtrUint8(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	item := itemValue.(*uint8)
	input[index] = item
	return input
}

func setUint16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	item := itemValue.(uint16)
	input[index] = item
	return input
}

func setPtrUint16(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	item := itemValue.(*uint16)
	input[index] = item
	return input
}

func setUint32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	item := itemValue.(uint32)
	input[index] = item
	return input
}

func setPtrUint32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	item := itemValue.(*uint32)
	input[index] = item
	return input
}

func setUint64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	item := itemValue.(uint64)
	input[index] = item
	return input
}

func setPtrUint64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	item := itemValue.(*uint64)
	input[index] = item
	return input
}

func setFloat32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	item := itemValue.(float32)
	input[index] = item
	return input
}

func setPtrFloat32(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	item := itemValue.(*float32)
	input[index] = item
	return input
}

func setFloat64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	item := itemValue.(float64)
	input[index] = item
	return input
}

func setPtrFloat64(itemsValue reflect.Value, index int, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	item := itemValue.(*float64)
	input[index] = item
	return input
}
