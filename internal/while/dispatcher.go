package while

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, sfunction, dfunction interface{}) interface{}

var dispatcher = map[string]dispatchFunction{
	"string":   whileString,
	"*string":  whilePtrString,
	"bool":     whileBool,
	"*bool":    whilePtrBool,
	"int":      whileInt,
	"*int":     whilePtrInt,
	"int8":     whileInt8,
	"*int8":    whilePtrInt8,
	"int16":    whileInt16,
	"*int16":   whilePtrInt16,
	"int32":    whileInt32,
	"*int32":   whilePtrInt32,
	"int64":    whileInt64,
	"*int64":   whilePtrInt64,
	"uint":     whileUint,
	"*uint":    whilePtrUint,
	"uint8":    whileUint8,
	"*uint8":   whilePtrUint8,
	"uint16":   whileUint16,
	"*uint16":  whilePtrUint16,
	"uint32":   whileUint32,
	"*uint32":  whilePtrUint32,
	"uint64":   whileUint64,
	"*uint64":  whilePtrUint64,
	"float32":  whileFloat32,
	"*float32": whilePtrFloat32,
	"float64":  whileFloat64,
	"*float64": whilePtrFloat64,
}

func dispatch(items reflect.Value, sfunction, dfunction interface{}, info *whileInfo) (bool, interface{}) {
	input := info.fnInputType.String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, sfunction, dfunction)

	}
	return false, nil
}

func whileString(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	sfn := sfunction.(func(string) bool)
	dfn := dfunction.(func(string) string)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrString(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	sfn := sfunction.(func(*string) bool)
	dfn := dfunction.(func(*string) *string)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileBool(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	sfn := sfunction.(func(bool) bool)
	dfn := dfunction.(func(bool) bool)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrBool(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	sfn := sfunction.(func(*bool) bool)
	dfn := dfunction.(func(*bool) *bool)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileInt(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	sfn := sfunction.(func(int) bool)
	dfn := dfunction.(func(int) int)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrInt(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	sfn := sfunction.(func(*int) bool)
	dfn := dfunction.(func(*int) *int)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileInt8(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	sfn := sfunction.(func(int8) bool)
	dfn := dfunction.(func(int8) int8)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrInt8(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	sfn := sfunction.(func(*int8) bool)
	dfn := dfunction.(func(*int8) *int8)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileInt16(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	sfn := sfunction.(func(int16) bool)
	dfn := dfunction.(func(int16) int16)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrInt16(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	sfn := sfunction.(func(*int16) bool)
	dfn := dfunction.(func(*int16) *int16)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileInt32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	sfn := sfunction.(func(int32) bool)
	dfn := dfunction.(func(int32) int32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrInt32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	sfn := sfunction.(func(*int32) bool)
	dfn := dfunction.(func(*int32) *int32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileInt64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	sfn := sfunction.(func(int64) bool)
	dfn := dfunction.(func(int64) int64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrInt64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	sfn := sfunction.(func(*int64) bool)
	dfn := dfunction.(func(*int64) *int64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileUint(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	sfn := sfunction.(func(uint) bool)
	dfn := dfunction.(func(uint) uint)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrUint(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	sfn := sfunction.(func(*uint) bool)
	dfn := dfunction.(func(*uint) *uint)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileUint8(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	sfn := sfunction.(func(uint8) bool)
	dfn := dfunction.(func(uint8) uint8)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrUint8(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	sfn := sfunction.(func(*uint8) bool)
	dfn := dfunction.(func(*uint8) *uint8)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileUint16(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	sfn := sfunction.(func(uint16) bool)
	dfn := dfunction.(func(uint16) uint16)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrUint16(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	sfn := sfunction.(func(*uint16) bool)
	dfn := dfunction.(func(*uint16) *uint16)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileUint32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	sfn := sfunction.(func(uint32) bool)
	dfn := dfunction.(func(uint32) uint32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrUint32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	sfn := sfunction.(func(*uint32) bool)
	dfn := dfunction.(func(*uint32) *uint32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileUint64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	sfn := sfunction.(func(uint64) bool)
	dfn := dfunction.(func(uint64) uint64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrUint64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	sfn := sfunction.(func(*uint64) bool)
	dfn := dfunction.(func(*uint64) *uint64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileFloat32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	sfn := sfunction.(func(float32) bool)
	dfn := dfunction.(func(float32) float32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrFloat32(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	sfn := sfunction.(func(*float32) bool)
	dfn := dfunction.(func(*float32) *float32)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whileFloat64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	sfn := sfunction.(func(float64) bool)
	dfn := dfunction.(func(float64) float64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}

func whilePtrFloat64(itemsValue reflect.Value, sfunction, dfunction interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	sfn := sfunction.(func(*float64) bool)
	dfn := dfunction.(func(*float64) *float64)
	for i := 0; i < len(input); i++ {
		if sfn(input[i]) {
			input[i] = dfn(input[i])
		}
	}
	return input
}
