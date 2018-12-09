package drop

import "reflect"

type dispatchFunction func(items reflect.Value, item interface{}) interface{}

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
	output := make([]string, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(string)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrString(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make([]*string, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*string)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make([]bool, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(bool)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrBool(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make([]*bool, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*bool)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make([]int, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(int)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrInt(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make([]*int, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*int)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make([]int8, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(int8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrInt8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make([]*int8, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*int8)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make([]int16, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(int16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrInt16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make([]*int16, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*int16)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make([]int32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(int32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrInt32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make([]*int32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*int32)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make([]int64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(int64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrInt64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make([]*int64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*int64)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make([]uint, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(uint)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrUint(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make([]*uint, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*uint)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make([]uint8, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(uint8)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrUint8(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make([]*uint8, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*uint8)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make([]uint16, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(uint16)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrUint16(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make([]*uint16, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*uint16)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make([]uint32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(uint32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrUint32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make([]*uint32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*uint32)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make([]uint64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(uint64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrUint64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make([]*uint64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*uint64)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make([]float32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(float32)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrFloat32(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make([]*float32, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*float32)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make([]float64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(float64)
	for i := 0; i < len(input); i++ {
		if input[i] == item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}

func dropPtrFloat64(itemsValue reflect.Value, itemValue interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make([]*float64, itemsValue.Len())
	currentIndex := 0
	item := itemValue.(*float64)
	for i := 0; i < len(input); i++ {
		if *input[i] == *item {
			continue
		}
		output[currentIndex] = input[i]
		currentIndex++
	}
	return output[:currentIndex]
}
