package contains

import "reflect"

type dispatchFunction func(*reflect.Value, interface{}) bool

var dispatcher = map[string]dispatchFunction{
	"string":   containsString,
	"*string":  containsPtrString,
	"bool":     containsBool,
	"*bool":    containsPtrBool,
	"int":      containsInt,
	"*int":     containsPtrInt,
	"int8":     containsInt8,
	"*int8":    containsPtrInt8,
	"int16":    containsInt16,
	"*int16":   containsPtrInt16,
	"int32":    containsInt32,
	"*int32":   containsPtrInt32,
	"int64":    containsInt64,
	"*int64":   containsPtrInt64,
	"uint":     containsUint,
	"*uint":    containsPtrUint,
	"uint8":    containsUint8,
	"*uint8":   containsPtrUint8,
	"uint16":   containsUint16,
	"*uint16":  containsPtrUint16,
	"uint32":   containsUint32,
	"*uint32":  containsPtrUint32,
	"uint64":   containsUint64,
	"*uint64":  containsPtrUint64,
	"float32":  containsFloat32,
	"*float32": containsPtrFloat32,
	"float64":  containsFloat64,
	"*float64": containsPtrFloat64,
}

func dispatch(items *reflect.Value, val interface{}, itemsType reflect.Type) (bool, bool) {
	input := itemsType.String()
	if fnVal, ok := dispatcher[input]; ok {
		return true, fnVal(items, val)
	}
	return false, false
}

func containsString(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(string)
	array := itemsValue.Interface().([]string)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrString(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*string)
	array := itemsValue.Interface().([]*string)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsBool(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(bool)
	array := itemsValue.Interface().([]bool)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrBool(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*bool)
	array := itemsValue.Interface().([]*bool)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsInt(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(int)
	array := itemsValue.Interface().([]int)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrInt(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*int)
	array := itemsValue.Interface().([]*int)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsInt8(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(int8)
	array := itemsValue.Interface().([]int8)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrInt8(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*int8)
	array := itemsValue.Interface().([]*int8)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsInt16(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(int16)
	array := itemsValue.Interface().([]int16)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrInt16(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*int16)
	array := itemsValue.Interface().([]*int16)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsInt32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(int32)
	array := itemsValue.Interface().([]int32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrInt32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*int32)
	array := itemsValue.Interface().([]*int32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsInt64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(int64)
	array := itemsValue.Interface().([]int64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrInt64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*int64)
	array := itemsValue.Interface().([]*int64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsUint(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(uint)
	array := itemsValue.Interface().([]uint)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrUint(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*uint)
	array := itemsValue.Interface().([]*uint)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsUint8(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(uint8)
	array := itemsValue.Interface().([]uint8)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrUint8(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*uint8)
	array := itemsValue.Interface().([]*uint8)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsUint16(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(uint16)
	array := itemsValue.Interface().([]uint16)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrUint16(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*uint16)
	array := itemsValue.Interface().([]*uint16)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsUint32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(uint32)
	array := itemsValue.Interface().([]uint32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrUint32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*uint32)
	array := itemsValue.Interface().([]*uint32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsUint64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(uint64)
	array := itemsValue.Interface().([]uint64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrUint64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*uint64)
	array := itemsValue.Interface().([]*uint64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsFloat32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(float32)
	array := itemsValue.Interface().([]float32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrFloat32(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*float32)
	array := itemsValue.Interface().([]*float32)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsFloat64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(float64)
	array := itemsValue.Interface().([]float64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}

func containsPtrFloat64(itemsValue *reflect.Value, val interface{}) bool {
	element := val.(*float64)
	array := itemsValue.Interface().([]*float64)
	for _, val := range array {
		if val == element {
			return true
		}
	}
	return false
}
