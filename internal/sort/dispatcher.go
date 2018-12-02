package sort

import (
	"reflect"
	"sort"
)

type dispatchFunction func(items reflect.Value, fn interface{}) (bool, interface{})

var dispatcher = map[string]dispatchFunction{
	"string":   sortString,
	"*string":  sortPtrString,
	"bool":     sortBool,
	"*bool":    sortPtrBool,
	"int":      sortInt,
	"*int":     sortPtrInt,
	"int8":     sortInt8,
	"*int8":    sortPtrInt8,
	"int16":    sortInt16,
	"*int16":   sortPtrInt16,
	"int32":    sortInt32,
	"*int32":   sortPtrInt32,
	"int64":    sortInt64,
	"*int64":   sortPtrInt64,
	"uint":     sortUint,
	"*uint":    sortPtrUint,
	"uint8":    sortUint8,
	"*uint8":   sortPtrUint8,
	"uint16":   sortUint16,
	"*uint16":  sortPtrUint16,
	"uint32":   sortUint32,
	"*uint32":  sortPtrUint32,
	"uint64":   sortUint64,
	"*uint64":  sortPtrUint64,
	"float32":  sortFloat32,
	"*float32": sortPtrFloat32,
	"float64":  sortFloat64,
	"*float64": sortPtrFloat64,
}

func dispatch(items reflect.Value, function interface{}, info *sortInfo) bool {
	if fnVal, ok := dispatcher[info.fnInput1Type.String()]; ok {
		fnVal(items, function)
		return true
	}
	return false
}

func sortString(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]string)
	fn := function.(func(string, string) int)
	sort.Strings(input)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrString(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*string)
	fn := function.(func(*string, *string) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortBool(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]bool)
	fn := function.(func(bool, bool) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrBool(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*bool)
	fn := function.(func(*bool, *bool) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortInt(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]int)
	fn := function.(func(int, int) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrInt(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*int)
	fn := function.(func(*int, *int) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortInt8(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]int8)
	fn := function.(func(int8, int8) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrInt8(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*int8)
	fn := function.(func(*int8, *int8) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortInt16(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]int16)
	fn := function.(func(int16, int16) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrInt16(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*int16)
	fn := function.(func(*int16, *int16) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortInt32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]int32)
	fn := function.(func(int32, int32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrInt32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*int32)
	fn := function.(func(*int32, *int32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortInt64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]int64)
	fn := function.(func(int64, int64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrInt64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*int64)
	fn := function.(func(*int64, *int64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortUint(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]uint)
	fn := function.(func(uint, uint) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrUint(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint, *uint) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortUint8(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]uint8)
	fn := function.(func(uint8, uint8) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrUint8(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint8, *uint8) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortUint16(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]uint16)
	fn := function.(func(uint16, uint16) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrUint16(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint16, *uint16) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortUint32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]uint32)
	fn := function.(func(uint32, uint32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrUint32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint32, *uint32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortUint64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]uint64)
	fn := function.(func(uint64, uint64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrUint64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint64, *uint64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortFloat32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]float32)
	fn := function.(func(float32, float32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrFloat32(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*float32)
	fn := function.(func(*float32, *float32) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}

func sortFloat64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]float64)
	fn := function.(func(float64, float64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input

}

func sortPtrFloat64(itemsValue reflect.Value, function interface{}) (bool, interface{}) {
	input := itemsValue.Interface().([]*float64)
	fn := function.(func(*float64, *float64) int)
	sort.Slice(input, func(i, j int) bool {
		left := input[i]
		right := input[j]
		return fn(left, right) <= 0
	})
	return true, input
}
