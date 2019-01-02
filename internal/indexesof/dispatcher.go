package indexesof

import "reflect"

type dispatchFunction func(items reflect.Value, fn interface{}) []int

var dispatcher = map[string]dispatchFunction{
	"string":   indexesOfString,
	"*string":  indexesOfPtrString,
	"bool":     indexesOfBool,
	"*bool":    indexesOfPtrBool,
	"int":      indexesOfInt,
	"*int":     indexesOfPtrInt,
	"int8":     indexesOfInt8,
	"*int8":    indexesOfPtrInt8,
	"int16":    indexesOfInt16,
	"*int16":   indexesOfPtrInt16,
	"int32":    indexesOfInt32,
	"*int32":   indexesOfPtrInt32,
	"int64":    indexesOfInt64,
	"*int64":   indexesOfPtrInt64,
	"uint":     indexesOfUint,
	"*uint":    indexesOfPtrUint,
	"uint8":    indexesOfUint8,
	"*uint8":   indexesOfPtrUint8,
	"uint16":   indexesOfUint16,
	"*uint16":  indexesOfPtrUint16,
	"uint32":   indexesOfUint32,
	"*uint32":  indexesOfPtrUint32,
	"uint64":   indexesOfUint64,
	"*uint64":  indexesOfPtrUint64,
	"float32":  indexesOfFloat32,
	"*float32": indexesOfPtrFloat32,
	"float64":  indexesOfFloat64,
	"*float64": indexesOfPtrFloat64,
}

func dispatch(items reflect.Value, val interface{}, itemsType reflect.Type) (bool, []int) {
	input := itemsType.String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, val)
	}
	return false, []int{}
}

func indexesOfString(itemsValue reflect.Value, item interface{}) []int {
	element := item.(string)
	array := itemsValue.Interface().([]string)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrString(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*string)
	array := itemsValue.Interface().([]*string)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfBool(itemsValue reflect.Value, item interface{}) []int {
	element := item.(bool)
	array := itemsValue.Interface().([]bool)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrBool(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*bool)
	array := itemsValue.Interface().([]*bool)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfInt(itemsValue reflect.Value, item interface{}) []int {
	element := item.(int)
	array := itemsValue.Interface().([]int)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrInt(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*int)
	array := itemsValue.Interface().([]*int)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfInt8(itemsValue reflect.Value, item interface{}) []int {
	element := item.(int8)
	array := itemsValue.Interface().([]int8)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrInt8(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*int8)
	array := itemsValue.Interface().([]*int8)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfInt16(itemsValue reflect.Value, item interface{}) []int {
	element := item.(int16)
	array := itemsValue.Interface().([]int16)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrInt16(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*int16)
	array := itemsValue.Interface().([]*int16)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfInt32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(int32)
	array := itemsValue.Interface().([]int32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrInt32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*int32)
	array := itemsValue.Interface().([]*int32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfInt64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(int64)
	array := itemsValue.Interface().([]int64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrInt64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*int64)
	array := itemsValue.Interface().([]*int64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfUint(itemsValue reflect.Value, item interface{}) []int {
	element := item.(uint)
	array := itemsValue.Interface().([]uint)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrUint(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*uint)
	array := itemsValue.Interface().([]*uint)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfUint8(itemsValue reflect.Value, item interface{}) []int {
	element := item.(uint8)
	array := itemsValue.Interface().([]uint8)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrUint8(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*uint8)
	array := itemsValue.Interface().([]*uint8)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfUint16(itemsValue reflect.Value, item interface{}) []int {
	element := item.(uint16)
	array := itemsValue.Interface().([]uint16)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrUint16(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*uint16)
	array := itemsValue.Interface().([]*uint16)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfUint32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(uint32)
	array := itemsValue.Interface().([]uint32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrUint32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*uint32)
	array := itemsValue.Interface().([]*uint32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfUint64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(uint64)
	array := itemsValue.Interface().([]uint64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrUint64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*uint64)
	array := itemsValue.Interface().([]*uint64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfFloat32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(float32)
	array := itemsValue.Interface().([]float32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrFloat32(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*float32)
	array := itemsValue.Interface().([]*float32)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfFloat64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(float64)
	array := itemsValue.Interface().([]float64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if array[k] == element {
			output = append(output, k)
		}
	}
	return output
}

func indexesOfPtrFloat64(itemsValue reflect.Value, item interface{}) []int {
	element := item.(*float64)
	array := itemsValue.Interface().([]*float64)
	output := make([]int, 0)
	for k := 0; k <= len(array)-1; k++ {
		if *array[k] == *element {
			output = append(output, k)
		}
	}
	return output
}
