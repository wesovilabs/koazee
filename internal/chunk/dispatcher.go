package chunk

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, size int) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   chunkString,
	"*string":  chunkStringPtr,
	"bool":     chunkBool,
	"*bool":    chunkBoolPtr,
	"int":      chunkInt,
	"*int":     chunkIntPtr,
	"int8":     chunkInt8,
	"*int8":    chunkInt8Ptr,
	"int16":    chunkInt16,
	"*int16":   chunkInt16Ptr,
	"int32":    chunkInt32,
	"*int32":   chunkInt32Ptr,
	"int64":    chunkInt64,
	"*int64":   chunkInt64Ptr,
	"uint":     chunkUint,
	"*uint":    chunkUintPtr,
	"uint8":    chunkUint8,
	"*uint8":   chunkUint8Ptr,
	"uint16":   chunkUint16,
	"*uint16":  chunkUint16Ptr,
	"uint32":   chunkUint32,
	"*uint32":  chunkUint32Ptr,
	"uint64":   chunkUint64,
	"*uint64":  chunkUint64Ptr,
	"float32":  chunkFloat32,
	"*float32": chunkFloat32Ptr,
	"float64":  chunkFloat64,
	"*float64": chunkFloat64Ptr,
}

func dispatch(items reflect.Value, itemsType reflect.Type, size uint) (bool, reflect.Value) {
	if fnVal, ok := dispatcher[itemsType.String()]; ok {
		return true, reflect.ValueOf(fnVal(items, int(size)))
	}
	return false, reflect.ValueOf(nil)
}
func chunkString(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]string)
	len := len(input)
	output := make([][]string, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkStringPtr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*string)
	len := len(input)
	output := make([][]*string, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkBool(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]bool)
	len := len(input)
	output := make([][]bool, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkBoolPtr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*bool)
	len := len(input)
	output := make([][]*bool, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkInt(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]int)
	len := len(input)
	output := make([][]int, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkIntPtr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*int)
	len := len(input)
	output := make([][]*int, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkInt8(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]int8)
	len := len(input)
	output := make([][]int8, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkInt8Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*int8)
	len := len(input)
	output := make([][]*int8, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkInt16(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]int16)
	len := len(input)
	output := make([][]int16, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkInt16Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*int16)
	len := len(input)
	output := make([][]*int16, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkInt32(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]int32)
	len := len(input)
	output := make([][]int32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkInt32Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*int32)
	len := len(input)
	output := make([][]*int32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkInt64(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]int64)
	len := len(input)
	output := make([][]int64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkInt64Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*int64)
	len := len(input)
	output := make([][]*int64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkUint(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]uint)
	len := len(input)
	output := make([][]uint, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkUintPtr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*uint)
	len := len(input)
	output := make([][]*uint, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkUint8(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]uint8)
	len := len(input)
	output := make([][]uint8, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkUint8Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*uint8)
	len := len(input)
	output := make([][]*uint8, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkUint16(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]uint16)
	len := len(input)
	output := make([][]uint16, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkUint16Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*uint16)
	len := len(input)
	output := make([][]*uint16, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkUint32(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]uint32)
	len := len(input)
	output := make([][]uint32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkUint32Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*uint32)
	len := len(input)
	output := make([][]*uint32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkUint64(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]uint64)
	len := len(input)
	output := make([][]uint64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkUint64Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*uint64)
	len := len(input)
	output := make([][]*uint64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkFloat32(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]float32)
	len := len(input)
	output := make([][]float32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkFloat32Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*float32)
	len := len(input)
	output := make([][]*float32, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
func chunkFloat64(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]float64)
	len := len(input)
	output := make([][]float64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}

func chunkFloat64Ptr(itemsValue reflect.Value, size int) interface{} {
	input := itemsValue.Interface().([]*float64)
	len := len(input)
	output := make([][]*float64, 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = append(output, input[index:rangeLast:rangeLast])
	}
	return output
}
