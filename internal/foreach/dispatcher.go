package foreach

import (
	"reflect"
)

type dispatchFunction func(items reflect.Value, fn interface{})

var dispatcher = map[string]dispatchFunction{
	"string":   forEachString,
	"*string":  forEachPtrString,
	"bool":     forEachBool,
	"*bool":    forEachPtrBool,
	"int":      forEachInt,
	"*int":     forEachPtrInt,
	"int8":     forEachInt8,
	"*int8":    forEachPtrInt8,
	"int16":    forEachInt16,
	"*int16":   forEachPtrInt16,
	"int32":    forEachInt32,
	"*int32":   forEachPtrInt32,
	"int64":    forEachInt64,
	"*int64":   forEachPtrInt64,
	"uint":     forEachUint,
	"*uint":    forEachPtrUint,
	"uint8":    forEachUint8,
	"*uint8":   forEachPtrUint8,
	"uint16":   forEachUint16,
	"*uint16":  forEachPtrUint16,
	"uint32":   forEachUint32,
	"*uint32":  forEachPtrUint32,
	"uint64":   forEachUint64,
	"*uint64":  forEachPtrUint64,
	"float32":  forEachFloat32,
	"*float32": forEachPtrFloat32,
	"float64":  forEachFloat64,
	"*float64": forEachPtrFloat64,
}

func dispatch(items reflect.Value, itemsType reflect.Type, info *forEachInfo) bool {
	input := itemsType.String()
	if fnVal, ok := dispatcher[input]; ok {
		fnVal(items, info.fnValue.Interface())
		return true
	}
	return false
}

func forEachString(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]string)
	fn := function.(func(string))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrString(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*string)
	fn := function.(func(*string))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachBool(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]bool)
	fn := function.(func(bool))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrBool(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*bool)
	fn := function.(func(*bool))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachInt(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]int)
	fn := function.(func(int))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrInt(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*int)
	fn := function.(func(*int))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachInt8(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]int8)
	fn := function.(func(int8))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrInt8(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*int8)
	fn := function.(func(*int8))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachInt16(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]int16)
	fn := function.(func(int16))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrInt16(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*int16)
	fn := function.(func(*int16))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachInt32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]int32)
	fn := function.(func(int32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrInt32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*int32)
	fn := function.(func(*int32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachInt64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]int64)
	fn := function.(func(int64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrInt64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*int64)
	fn := function.(func(*int64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachUint(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]uint)
	fn := function.(func(uint))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrUint(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachUint8(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]uint8)
	fn := function.(func(uint8))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrUint8(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint8))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachUint16(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]uint16)
	fn := function.(func(uint16))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrUint16(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint16))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachUint32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]uint32)
	fn := function.(func(uint32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrUint32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachUint64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]uint64)
	fn := function.(func(uint64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrUint64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachFloat32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]float32)
	fn := function.(func(float32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrFloat32(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*float32)
	fn := function.(func(*float32))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachFloat64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]float64)
	fn := function.(func(float64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}

func forEachPtrFloat64(itemsValue reflect.Value, function interface{}) {
	input := itemsValue.Interface().([]*float64)
	fn := function.(func(*float64))
	for i := 0; i < len(input); i++ {
		fn(input[i])
	}
}
