package drop

import "reflect"

type dispatchFunction func(items reflect.Value, item interface{}) interface{}

func reverseIndexes(indexes []int) []int {
	for i := len(indexes)/2 - 1; i >= 0; i-- {
		opp := len(indexes) - 1 - i
		indexes[i], indexes[opp] = indexes[opp], indexes[i]
	}
	return indexes
}

var dispatcher = map[string]dispatchFunction{
	"string":   dropString,
	"*string":  dropPtrString,
	"bool":     dropBool,
	"*bool":    dropPtrBool,
	"int":      dropInt,
	"*int":     dropPtrInt,
	"int8":     dropInt8,
	"*int8":    dropPtrInt8,
	"int16":    dropInt16,
	"*int16":   dropPtrInt16,
	"int32":    dropInt32,
	"*int32":   dropPtrInt32,
	"int64":    dropInt64,
	"*int64":   dropPtrInt64,
	"uint":     dropUint,
	"*uint":    dropPtrUint,
	"uint8":    dropUint8,
	"*uint8":   dropPtrUint8,
	"uint16":   dropUint16,
	"*uint16":  dropPtrUint16,
	"uint32":   dropUint32,
	"*uint32":  dropPtrUint32,
	"uint64":   dropUint64,
	"*uint64":  dropPtrUint64,
	"float32":  dropFloat32,
	"*float32": dropPtrFloat32,
	"float64":  dropFloat64,
	"*float64": dropPtrFloat64,
}

func dispatch(items reflect.Value, itemValue interface{}, info *dropInfo) (bool, interface{}) {
	input := (*info.itemType).String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, itemValue)
	}
	return false, nil
}

func dropString(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	indexes := make([]int, 0)
	item := itemValue.(string)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrString(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	indexes := make([]int, 0)
	item := itemValue.(*string)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	indexes := make([]int, 0)
	item := itemValue.(bool)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	indexes := make([]int, 0)
	item := itemValue.(*bool)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	indexes := make([]int, 0)
	item := itemValue.(int)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	indexes := make([]int, 0)
	item := itemValue.(*int)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	indexes := make([]int, 0)
	item := itemValue.(int8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	indexes := make([]int, 0)
	item := itemValue.(*int8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	indexes := make([]int, 0)
	item := itemValue.(int16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	indexes := make([]int, 0)
	item := itemValue.(*int16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	indexes := make([]int, 0)
	item := itemValue.(int32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	indexes := make([]int, 0)
	item := itemValue.(*int32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	indexes := make([]int, 0)
	item := itemValue.(int64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	indexes := make([]int, 0)
	item := itemValue.(*int64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	indexes := make([]int, 0)
	item := itemValue.(uint)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	indexes := make([]int, 0)
	item := itemValue.(*uint)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	indexes := make([]int, 0)
	item := itemValue.(uint8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	indexes := make([]int, 0)
	item := itemValue.(*uint8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	indexes := make([]int, 0)
	item := itemValue.(uint16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	indexes := make([]int, 0)
	item := itemValue.(*uint16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	indexes := make([]int, 0)
	item := itemValue.(uint32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	indexes := make([]int, 0)
	item := itemValue.(*uint32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	indexes := make([]int, 0)
	item := itemValue.(uint64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	indexes := make([]int, 0)
	item := itemValue.(*uint64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	indexes := make([]int, 0)
	item := itemValue.(float32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	indexes := make([]int, 0)
	item := itemValue.(*float32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}

func dropFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	indexes := make([]int, 0)
	item := itemValue.(float64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		if index > 0 {
			input = append(input[:index], input[index+1:]...)
			continue
		}
		input = input[index+1:]
	}
	return input
}

func dropPtrFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	indexes := make([]int, 0)
	item := itemValue.(*float64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			indexes = append(indexes, i)
		}
	}
	indexes = reverseIndexes(indexes)
	for _, index := range indexes {
		input = append(input[:index], input[index+1:]...)
	}
	return input
}
