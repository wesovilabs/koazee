package dropWhile

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, fn interface{}) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   dropWhileString,
	"*string":  dropWhilePtrString,
	"bool":     dropWhileBool,
	"*bool":    dropWhilePtrBool,
	"int":      dropWhileInt,
	"*int":     dropWhilePtrInt,
	"int8":     dropWhileInt8,
	"*int8":    dropWhilePtrInt8,
	"int16":    dropWhileInt16,
	"*int16":   dropWhilePtrInt16,
	"int32":    dropWhileInt32,
	"*int32":   dropWhilePtrInt32,
	"int64":    dropWhileInt64,
	"*int64":   dropWhilePtrInt64,
	"uint":     dropWhileUint,
	"*uint":    dropWhilePtrUint,
	"uint8":    dropWhileUint8,
	"*uint8":   dropWhilePtrUint8,
	"uint16":   dropWhileUint16,
	"*uint16":  dropWhilePtrUint16,
	"uint32":   dropWhileUint32,
	"*uint32":  dropWhilePtrUint32,
	"uint64":   dropWhileUint64,
	"*uint64":  dropWhilePtrUint64,
	"float32":  dropWhileFloat32,
	"*float32": dropWhilePtrFloat32,
	"float64":  dropWhileFloat64,
	"*float64": dropWhilePtrFloat64,
}

func dispatch(items reflect.Value, function interface{}, info *dropWhileInfo) (bool, reflect.Value) {
	input := info.fnInputType.String()
	if fnVal, ok := dispatcher[input]; ok {
		res := reflect.ValueOf(fnVal(items, function))
		return true, res
	}
	return false, reflect.ValueOf(nil)
}

func dropWhileString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make([]string, 0)
	fn := function.(func(string) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make([]*string, 0)
	fn := function.(func(*string) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make([]bool, 0)
	fn := function.(func(bool) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make([]*bool, 0)
	fn := function.(func(*bool) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make([]int, 0)
	fn := function.(func(int) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make([]*int, 0)
	fn := function.(func(*int) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make([]int8, 0)
	fn := function.(func(int8) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make([]*int8, 0)
	fn := function.(func(*int8) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make([]int16, 0)
	fn := function.(func(int16) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make([]*int16, 0)
	fn := function.(func(*int16) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make([]int32, 0)
	fn := function.(func(int32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make([]*int32, 0)
	fn := function.(func(*int32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make([]int64, 0)
	fn := function.(func(int64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make([]*int64, 0)
	fn := function.(func(*int64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make([]uint, 0)
	fn := function.(func(uint) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make([]*uint, 0)
	fn := function.(func(*uint) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make([]uint8, 0)
	fn := function.(func(uint8) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make([]*uint8, 0)
	fn := function.(func(*uint8) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make([]uint16, 0)
	fn := function.(func(uint16) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make([]*uint16, 0)
	fn := function.(func(*uint16) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make([]uint32, 0)
	fn := function.(func(uint32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make([]*uint32, 0)
	fn := function.(func(*uint32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make([]uint64, 0)
	fn := function.(func(uint64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make([]*uint64, 0)
	fn := function.(func(*uint64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make([]float32, 0)
	fn := function.(func(float32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make([]*float32, 0)
	fn := function.(func(*float32) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhileFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make([]float64, 0)
	fn := function.(func(float64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func dropWhilePtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make([]*float64, 0)
	fn := function.(func(*float64) bool)
	for i := 0; i < len(input); i++ {
		if !fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}
