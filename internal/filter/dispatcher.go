package filter

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, fn interface{}) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   filterString,
	"*string":  filterPtrString,
	"bool":     filterBool,
	"*bool":    filterPtrBool,
	"int":      filterInt,
	"*int":     filterPtrInt,
	"int8":     filterInt8,
	"*int8":    filterPtrInt8,
	"int16":    filterInt16,
	"*int16":   filterPtrInt16,
	"int32":    filterInt32,
	"*int32":   filterPtrInt32,
	"int64":    filterInt64,
	"*int64":   filterPtrInt64,
	"uint":     filterUint,
	"*uint":    filterPtrUint,
	"uint8":    filterUint8,
	"*uint8":   filterPtrUint8,
	"uint16":   filterUint16,
	"*uint16":  filterPtrUint16,
	"uint32":   filterUint32,
	"*uint32":  filterPtrUint32,
	"uint64":   filterUint64,
	"*uint64":  filterPtrUint64,
	"float32":  filterFloat32,
	"*float32": filterPtrFloat32,
	"float64":  filterFloat64,
	"*float64": filterPtrFloat64,
}

func dispatch(items reflect.Value, function interface{}, info *filterInfo) (bool, reflect.Value) {
	input := info.fnIn1Type.String()
	if fnVal, ok := dispatcher[input]; ok {
		if info.fnIn2Type == nil {
			res := reflect.ValueOf(fnVal(items, function))
			return true, res
		}
	}
	return false, reflect.ValueOf(nil)
}

func filterString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make([]string, 0)
	fn := function.(func(string) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make([]*string, 0)
	fn := function.(func(*string) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make([]bool, 0)
	fn := function.(func(bool) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make([]*bool, 0)
	fn := function.(func(*bool) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make([]int, 0)
	fn := function.(func(int) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make([]*int, 0)
	fn := function.(func(*int) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make([]int8, 0)
	fn := function.(func(int8) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make([]*int8, 0)
	fn := function.(func(*int8) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make([]int16, 0)
	fn := function.(func(int16) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make([]*int16, 0)
	fn := function.(func(*int16) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make([]int32, 0)
	fn := function.(func(int32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make([]*int32, 0)
	fn := function.(func(*int32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make([]int64, 0)
	fn := function.(func(int64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make([]*int64, 0)
	fn := function.(func(*int64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make([]uint, 0)
	fn := function.(func(uint) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make([]*uint, 0)
	fn := function.(func(*uint) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make([]uint8, 0)
	fn := function.(func(uint8) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make([]*uint8, 0)
	fn := function.(func(*uint8) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make([]uint16, 0)
	fn := function.(func(uint16) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make([]*uint16, 0)
	fn := function.(func(*uint16) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make([]uint32, 0)
	fn := function.(func(uint32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make([]*uint32, 0)
	fn := function.(func(*uint32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make([]uint64, 0)
	fn := function.(func(uint64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make([]*uint64, 0)
	fn := function.(func(*uint64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make([]float32, 0)
	fn := function.(func(float32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make([]*float32, 0)
	fn := function.(func(*float32) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make([]float64, 0)
	fn := function.(func(float64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func filterPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make([]*float64, 0)
	fn := function.(func(*float64) bool)
	for i := 0; i < len(input); i++ {
		if fn(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}
