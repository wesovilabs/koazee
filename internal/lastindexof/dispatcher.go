package lastindexof

import "reflect"

type dispatchFunction func(reflect.Value, interface{}) int

var dispatcher = map[string]dispatchFunction{
	"string":   indexOfString,
	"*string":  indexOfPtrString,
	"bool":     indexOfBool,
	"*bool":    indexOfPtrBool,
	"int":      indexOfInt,
	"*int":     indexOfPtrInt,
	"int8":     indexOfInt8,
	"*int8":    indexOfPtrInt8,
	"int16":    indexOfInt16,
	"*int16":   indexOfPtrInt16,
	"int32":    indexOfInt32,
	"*int32":   indexOfPtrInt32,
	"int64":    indexOfInt64,
	"*int64":   indexOfPtrInt64,
	"uint":     indexOfUint,
	"*uint":    indexOfPtrUint,
	"uint8":    indexOfUint8,
	"*uint8":   indexOfPtrUint8,
	"uint16":   indexOfUint16,
	"*uint16":  indexOfPtrUint16,
	"uint32":   indexOfUint32,
	"*uint32":  indexOfPtrUint32,
	"uint64":   indexOfUint64,
	"*uint64":  indexOfPtrUint64,
	"float32":  indexOfFloat32,
	"*float32": indexOfPtrFloat32,
	"float64":  indexOfFloat64,
	"*float64": indexOfPtrFloat64,
}

func dispatch(items reflect.Value, val interface{}, itemsType reflect.Type) (bool, int) {
	input := itemsType.String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, val)
	}
	return false, InvalidIndex
}

func indexOfString(itemsValue reflect.Value, item interface{}) int {
	element := item.(string)
	array := itemsValue.Interface().([]string)

	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrString(itemsValue reflect.Value, item interface{}) int {
	element := item.(*string)
	array := itemsValue.Interface().([]*string)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfBool(itemsValue reflect.Value, item interface{}) int {
	element := item.(bool)
	array := itemsValue.Interface().([]bool)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrBool(itemsValue reflect.Value, item interface{}) int {
	element := item.(*bool)
	array := itemsValue.Interface().([]*bool)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfInt(itemsValue reflect.Value, item interface{}) int {
	element := item.(int)
	array := itemsValue.Interface().([]int)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrInt(itemsValue reflect.Value, item interface{}) int {
	element := item.(*int)
	array := itemsValue.Interface().([]*int)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfInt8(itemsValue reflect.Value, item interface{}) int {
	element := item.(int8)
	array := itemsValue.Interface().([]int8)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrInt8(itemsValue reflect.Value, item interface{}) int {
	element := item.(*int8)
	array := itemsValue.Interface().([]*int8)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfInt16(itemsValue reflect.Value, item interface{}) int {
	element := item.(int16)
	array := itemsValue.Interface().([]int16)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrInt16(itemsValue reflect.Value, item interface{}) int {
	element := item.(*int16)
	array := itemsValue.Interface().([]*int16)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfInt32(itemsValue reflect.Value, item interface{}) int {
	element := item.(int32)
	array := itemsValue.Interface().([]int32)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrInt32(itemsValue reflect.Value, item interface{}) int {
	element := item.(*int32)
	array := itemsValue.Interface().([]*int32)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfInt64(itemsValue reflect.Value, item interface{}) int {
	element := item.(int64)
	array := itemsValue.Interface().([]int64)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrInt64(itemsValue reflect.Value, item interface{}) int {
	element := item.(*int64)
	array := itemsValue.Interface().([]*int64)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfUint(itemsValue reflect.Value, item interface{}) int {
	element := item.(uint)
	array := itemsValue.Interface().([]uint)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrUint(itemsValue reflect.Value, item interface{}) int {
	element := item.(*uint)
	array := itemsValue.Interface().([]*uint)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfUint8(itemsValue reflect.Value, item interface{}) int {
	element := item.(uint8)
	array := itemsValue.Interface().([]uint8)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrUint8(itemsValue reflect.Value, item interface{}) int {
	element := item.(*uint8)
	array := itemsValue.Interface().([]*uint8)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfUint16(itemsValue reflect.Value, item interface{}) int {
	element := item.(uint16)
	array := itemsValue.Interface().([]uint16)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrUint16(itemsValue reflect.Value, item interface{}) int {
	element := item.(*uint16)
	array := itemsValue.Interface().([]*uint16)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfUint32(itemsValue reflect.Value, item interface{}) int {
	element := item.(uint32)
	array := itemsValue.Interface().([]uint32)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrUint32(itemsValue reflect.Value, item interface{}) int {
	element := item.(*uint32)
	array := itemsValue.Interface().([]*uint32)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfUint64(itemsValue reflect.Value, item interface{}) int {
	element := item.(uint64)
	array := itemsValue.Interface().([]uint64)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrUint64(itemsValue reflect.Value, item interface{}) int {
	element := item.(*uint64)
	array := itemsValue.Interface().([]*uint64)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfFloat32(itemsValue reflect.Value, item interface{}) int {
	element := item.(float32)
	array := itemsValue.Interface().([]float32)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrFloat32(itemsValue reflect.Value, item interface{}) int {
	element := item.(*float32)
	array := itemsValue.Interface().([]*float32)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfFloat64(itemsValue reflect.Value, item interface{}) int {
	element := item.(float64)
	array := itemsValue.Interface().([]float64)
	for k := len(array) - 1; k >= 0; k-- {
		if array[k] == element {
			return k
		}
	}
	return InvalidIndex
}

func indexOfPtrFloat64(itemsValue reflect.Value, item interface{}) int {
	element := item.(*float64)
	array := itemsValue.Interface().([]*float64)
	for k := len(array) - 1; k >= 0; k-- {
		if *array[k] == *element {
			return k
		}
	}
	return InvalidIndex
}
