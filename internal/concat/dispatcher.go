package concat

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, item reflect.Value) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   concatString,
	"*string":  concatPtrString,
	"bool":     concatBool,
	"*bool":    concatPtrBool,
	"int":      concatInt,
	"*int":     concatPtrInt,
	"int8":     concatInt8,
	"*int8":    concatPtrInt8,
	"int16":    concatInt16,
	"*int16":   concatPtrInt16,
	"int32":    concatInt32,
	"*int32":   concatPtrInt32,
	"int64":    concatInt64,
	"*int64":   concatPtrInt64,
	"uint":     concatUint,
	"*uint":    concatPtrUint,
	"uint8":    concatUint8,
	"*uint8":   concatPtrUint8,
	"uint16":   concatUint16,
	"*uint16":  concatPtrUint16,
	"uint32":   concatUint32,
	"*uint32":  concatPtrUint32,
	"uint64":   concatUint64,
	"*uint64":  concatPtrUint64,
	"float32":  concatFloat32,
	"*float32": concatPtrFloat32,
	"float64":  concatFloat64,
	"*float64": concatPtrFloat64,
}

func dispatch(items reflect.Value, item reflect.Value, itemsType reflect.Type) (bool, reflect.Value) {
	input := itemsType.String()
	if fnVal, ok := dispatcher[input]; ok {
		res := reflect.ValueOf(fnVal(items, item))
		return true, res
	}
	return false, reflect.ValueOf(nil)
}

func concatString(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]string)
	output := itemsValue.Interface().([]string)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrString(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*string)
	output := itemsValue.Interface().([]*string)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatBool(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]bool)
	output := itemsValue.Interface().([]bool)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrBool(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*bool)
	output := itemsValue.Interface().([]*bool)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatInt(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]int)
	output := itemsValue.Interface().([]int)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrInt(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*int)
	output := itemsValue.Interface().([]*int)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatInt8(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]int8)
	output := itemsValue.Interface().([]int8)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrInt8(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*int8)
	output := itemsValue.Interface().([]*int8)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatInt16(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]int16)
	output := itemsValue.Interface().([]int16)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrInt16(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*int16)
	output := itemsValue.Interface().([]*int16)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatInt32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]int32)
	output := itemsValue.Interface().([]int32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrInt32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*int32)
	output := itemsValue.Interface().([]*int32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatInt64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]int64)
	output := itemsValue.Interface().([]int64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrInt64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*int64)
	output := itemsValue.Interface().([]*int64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatUint(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]uint)
	output := itemsValue.Interface().([]uint)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrUint(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*uint)
	output := itemsValue.Interface().([]*uint)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatUint8(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]uint8)
	output := itemsValue.Interface().([]uint8)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrUint8(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*uint8)
	output := itemsValue.Interface().([]*uint8)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatUint16(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]uint16)
	output := itemsValue.Interface().([]uint16)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrUint16(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*uint16)
	output := itemsValue.Interface().([]*uint16)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatUint32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]uint32)
	output := itemsValue.Interface().([]uint32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrUint32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*uint32)
	output := itemsValue.Interface().([]*uint32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatUint64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]uint64)
	output := itemsValue.Interface().([]uint64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrUint64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*uint64)
	output := itemsValue.Interface().([]*uint64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatFloat32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]float32)
	output := itemsValue.Interface().([]float32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrFloat32(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*float32)
	output := itemsValue.Interface().([]*float32)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatFloat64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]float64)
	output := itemsValue.Interface().([]float64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

func concatPtrFloat64(itemsValue reflect.Value, item reflect.Value) interface{} {
	input := item.Interface().([]*float64)
	output := itemsValue.Interface().([]*float64)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}
