package groupby

import (
	"fmt"
	"reflect"
)

type dispatchFunction func(itemsValue reflect.Value, function interface{}) interface{}

var dispatcher = map[string]map[string]dispatchFunction{
	"string": {
		"string":  groupByStringToMapOfString,
		"bool":    groupByStringToMapOfBool,
		"int":     groupByStringToMapOfInt,
		"int8":    groupByStringToMapOfInt8,
		"int16":   groupByStringToMapOfInt16,
		"int32":   groupByStringToMapOfInt32,
		"int64":   groupByStringToMapOfInt64,
		"uint":    groupByStringToMapOfUint,
		"uint8":   groupByStringToMapOfUint8,
		"uint16":  groupByStringToMapOfUint16,
		"uint32":  groupByStringToMapOfUint32,
		"uint64":  groupByStringToMapOfUint64,
		"float32": groupByStringToMapOfFloat32,
		"float64": groupByStringToMapOfFloat64,
	},
	"*string": {
		"string":  groupByStringToMapOfPtrString,
		"bool":    groupByStringToMapOfPtrBool,
		"int":     groupByStringToMapOfPtrInt,
		"int8":    groupByStringToMapOfPtrInt8,
		"int16":   groupByStringToMapOfPtrInt16,
		"int32":   groupByStringToMapOfPtrInt32,
		"int64":   groupByStringToMapOfPtrInt64,
		"uint":    groupByStringToMapOfPtrUint,
		"uint8":   groupByStringToMapOfPtrUint8,
		"uint16":  groupByStringToMapOfPtrUint16,
		"uint32":  groupByStringToMapOfPtrUint32,
		"uint64":  groupByStringToMapOfPtrUint64,
		"float32": groupByStringToMapOfPtrFloat32,
		"float64": groupByStringToMapOfPtrFloat64,
	},
	"bool": {
		"string":  groupByBoolToMapOfString,
		"bool":    groupByBoolToMapOfBool,
		"int":     groupByBoolToMapOfInt,
		"int8":    groupByBoolToMapOfInt8,
		"int16":   groupByBoolToMapOfInt16,
		"int32":   groupByBoolToMapOfInt32,
		"int64":   groupByBoolToMapOfInt64,
		"uint":    groupByBoolToMapOfUint,
		"uint8":   groupByBoolToMapOfUint8,
		"uint16":  groupByBoolToMapOfUint16,
		"uint32":  groupByBoolToMapOfUint32,
		"uint64":  groupByBoolToMapOfUint64,
		"float32": groupByBoolToMapOfFloat32,
		"float64": groupByBoolToMapOfFloat64,
	},
	"*bool": {
		"string":  groupByBoolToMapOfPtrString,
		"bool":    groupByBoolToMapOfPtrBool,
		"int":     groupByBoolToMapOfPtrInt,
		"int8":    groupByBoolToMapOfPtrInt8,
		"int16":   groupByBoolToMapOfPtrInt16,
		"int32":   groupByBoolToMapOfPtrInt32,
		"int64":   groupByBoolToMapOfPtrInt64,
		"uint":    groupByBoolToMapOfPtrUint,
		"uint8":   groupByBoolToMapOfPtrUint8,
		"uint16":  groupByBoolToMapOfPtrUint16,
		"uint32":  groupByBoolToMapOfPtrUint32,
		"uint64":  groupByBoolToMapOfPtrUint64,
		"float32": groupByBoolToMapOfPtrFloat32,
		"float64": groupByBoolToMapOfPtrFloat64,
	},
	"int": {
		"string":  groupByIntToMapOfString,
		"bool":    groupByIntToMapOfBool,
		"int":     groupByIntToMapOfInt,
		"int8":    groupByIntToMapOfInt8,
		"int16":   groupByIntToMapOfInt16,
		"int32":   groupByIntToMapOfInt32,
		"int64":   groupByIntToMapOfInt64,
		"uint":    groupByIntToMapOfUint,
		"uint8":   groupByIntToMapOfUint8,
		"uint16":  groupByIntToMapOfUint16,
		"uint32":  groupByIntToMapOfUint32,
		"uint64":  groupByIntToMapOfUint64,
		"float32": groupByIntToMapOfFloat32,
		"float64": groupByIntToMapOfFloat64,
	},
	"*int": {
		"string":  groupByIntToMapOfPtrString,
		"bool":    groupByIntToMapOfPtrBool,
		"int":     groupByIntToMapOfPtrInt,
		"int8":    groupByIntToMapOfPtrInt8,
		"int16":   groupByIntToMapOfPtrInt16,
		"int32":   groupByIntToMapOfPtrInt32,
		"int64":   groupByIntToMapOfPtrInt64,
		"uint":    groupByIntToMapOfPtrUint,
		"uint8":   groupByIntToMapOfPtrUint8,
		"uint16":  groupByIntToMapOfPtrUint16,
		"uint32":  groupByIntToMapOfPtrUint32,
		"uint64":  groupByIntToMapOfPtrUint64,
		"float32": groupByIntToMapOfPtrFloat32,
		"float64": groupByIntToMapOfPtrFloat64,
	},
	"int8": {
		"string":  groupByInt8ToMapOfString,
		"bool":    groupByInt8ToMapOfBool,
		"int":     groupByInt8ToMapOfInt,
		"int8":    groupByInt8ToMapOfInt8,
		"int16":   groupByInt8ToMapOfInt16,
		"int32":   groupByInt8ToMapOfInt32,
		"int64":   groupByInt8ToMapOfInt64,
		"uint":    groupByInt8ToMapOfUint,
		"uint8":   groupByInt8ToMapOfUint8,
		"uint16":  groupByInt8ToMapOfUint16,
		"uint32":  groupByInt8ToMapOfUint32,
		"uint64":  groupByInt8ToMapOfUint64,
		"float32": groupByInt8ToMapOfFloat32,
		"float64": groupByInt8ToMapOfFloat64,
	},
	"*int8": {
		"string":  groupByInt8ToMapOfPtrString,
		"bool":    groupByInt8ToMapOfPtrBool,
		"int":     groupByInt8ToMapOfPtrInt,
		"int8":    groupByInt8ToMapOfPtrInt8,
		"int16":   groupByInt8ToMapOfPtrInt16,
		"int32":   groupByInt8ToMapOfPtrInt32,
		"int64":   groupByInt8ToMapOfPtrInt64,
		"uint":    groupByInt8ToMapOfPtrUint,
		"uint8":   groupByInt8ToMapOfPtrUint8,
		"uint16":  groupByInt8ToMapOfPtrUint16,
		"uint32":  groupByInt8ToMapOfPtrUint32,
		"uint64":  groupByInt8ToMapOfPtrUint64,
		"float32": groupByInt8ToMapOfPtrFloat32,
		"float64": groupByInt8ToMapOfPtrFloat64,
	},
	"int16": {
		"string":  groupByInt16ToMapOfString,
		"bool":    groupByInt16ToMapOfBool,
		"int":     groupByInt16ToMapOfInt,
		"int8":    groupByInt16ToMapOfInt8,
		"int16":   groupByInt16ToMapOfInt16,
		"int32":   groupByInt16ToMapOfInt32,
		"int64":   groupByInt16ToMapOfInt64,
		"uint":    groupByInt16ToMapOfUint,
		"uint8":   groupByInt16ToMapOfUint8,
		"uint16":  groupByInt16ToMapOfUint16,
		"uint32":  groupByInt16ToMapOfUint32,
		"uint64":  groupByInt16ToMapOfUint64,
		"float32": groupByInt16ToMapOfFloat32,
		"float64": groupByInt16ToMapOfFloat64,
	},
	"*int16": {
		"string":  groupByInt16ToMapOfPtrString,
		"bool":    groupByInt16ToMapOfPtrBool,
		"int":     groupByInt16ToMapOfPtrInt,
		"int8":    groupByInt16ToMapOfPtrInt8,
		"int16":   groupByInt16ToMapOfPtrInt16,
		"int32":   groupByInt16ToMapOfPtrInt32,
		"int64":   groupByInt16ToMapOfPtrInt64,
		"uint":    groupByInt16ToMapOfPtrUint,
		"uint8":   groupByInt16ToMapOfPtrUint8,
		"uint16":  groupByInt16ToMapOfPtrUint16,
		"uint32":  groupByInt16ToMapOfPtrUint32,
		"uint64":  groupByInt16ToMapOfPtrUint64,
		"float32": groupByInt16ToMapOfPtrFloat32,
		"float64": groupByInt16ToMapOfPtrFloat64,
	},
	"int32": {
		"string":  groupByInt32ToMapOfString,
		"bool":    groupByInt32ToMapOfBool,
		"int":     groupByInt32ToMapOfInt,
		"int8":    groupByInt32ToMapOfInt8,
		"int16":   groupByInt32ToMapOfInt16,
		"int32":   groupByInt32ToMapOfInt32,
		"int64":   groupByInt32ToMapOfInt64,
		"uint":    groupByInt32ToMapOfUint,
		"uint8":   groupByInt32ToMapOfUint8,
		"uint16":  groupByInt32ToMapOfUint16,
		"uint32":  groupByInt32ToMapOfUint32,
		"uint64":  groupByInt32ToMapOfUint64,
		"float32": groupByInt32ToMapOfFloat32,
		"float64": groupByInt32ToMapOfFloat64,
	},
	"*int32": {
		"string":  groupByInt32ToMapOfPtrString,
		"bool":    groupByInt32ToMapOfPtrBool,
		"int":     groupByInt32ToMapOfPtrInt,
		"int8":    groupByInt32ToMapOfPtrInt8,
		"int16":   groupByInt32ToMapOfPtrInt16,
		"int32":   groupByInt32ToMapOfPtrInt32,
		"int64":   groupByInt32ToMapOfPtrInt64,
		"uint":    groupByInt32ToMapOfPtrUint,
		"uint8":   groupByInt32ToMapOfPtrUint8,
		"uint16":  groupByInt32ToMapOfPtrUint16,
		"uint32":  groupByInt32ToMapOfPtrUint32,
		"uint64":  groupByInt32ToMapOfPtrUint64,
		"float32": groupByInt32ToMapOfPtrFloat32,
		"float64": groupByInt32ToMapOfPtrFloat64,
	},
	"int64": {
		"string":  groupByInt64ToMapOfString,
		"bool":    groupByInt64ToMapOfBool,
		"int":     groupByInt64ToMapOfInt,
		"int8":    groupByInt64ToMapOfInt8,
		"int16":   groupByInt64ToMapOfInt16,
		"int32":   groupByInt64ToMapOfInt32,
		"int64":   groupByInt64ToMapOfInt64,
		"uint":    groupByInt64ToMapOfUint,
		"uint8":   groupByInt64ToMapOfUint8,
		"uint16":  groupByInt64ToMapOfUint16,
		"uint32":  groupByInt64ToMapOfUint32,
		"uint64":  groupByInt64ToMapOfUint64,
		"float32": groupByInt64ToMapOfFloat32,
		"float64": groupByInt64ToMapOfFloat64,
	},
	"*int64": {
		"string":  groupByInt64ToMapOfPtrString,
		"bool":    groupByInt64ToMapOfPtrBool,
		"int":     groupByInt64ToMapOfPtrInt,
		"int8":    groupByInt64ToMapOfPtrInt8,
		"int16":   groupByInt64ToMapOfPtrInt16,
		"int32":   groupByInt64ToMapOfPtrInt32,
		"int64":   groupByInt64ToMapOfPtrInt64,
		"uint":    groupByInt64ToMapOfPtrUint,
		"uint8":   groupByInt64ToMapOfPtrUint8,
		"uint16":  groupByInt64ToMapOfPtrUint16,
		"uint32":  groupByInt64ToMapOfPtrUint32,
		"uint64":  groupByInt64ToMapOfPtrUint64,
		"float32": groupByInt64ToMapOfPtrFloat32,
		"float64": groupByInt64ToMapOfPtrFloat64,
	},
	"uint": {
		"string":  groupByUintToMapOfString,
		"bool":    groupByUintToMapOfBool,
		"int":     groupByUintToMapOfInt,
		"int8":    groupByUintToMapOfInt8,
		"int16":   groupByUintToMapOfInt16,
		"int32":   groupByUintToMapOfInt32,
		"int64":   groupByUintToMapOfInt64,
		"uint":    groupByUintToMapOfUint,
		"uint8":   groupByUintToMapOfUint8,
		"uint16":  groupByUintToMapOfUint16,
		"uint32":  groupByUintToMapOfUint32,
		"uint64":  groupByUintToMapOfUint64,
		"float32": groupByUintToMapOfFloat32,
		"float64": groupByUintToMapOfFloat64,
	},
	"*uint": {
		"string":  groupByUintToMapOfPtrString,
		"bool":    groupByUintToMapOfPtrBool,
		"int":     groupByUintToMapOfPtrInt,
		"int8":    groupByUintToMapOfPtrInt8,
		"int16":   groupByUintToMapOfPtrInt16,
		"int32":   groupByUintToMapOfPtrInt32,
		"int64":   groupByUintToMapOfPtrInt64,
		"uint":    groupByUintToMapOfPtrUint,
		"uint8":   groupByUintToMapOfPtrUint8,
		"uint16":  groupByUintToMapOfPtrUint16,
		"uint32":  groupByUintToMapOfPtrUint32,
		"uint64":  groupByUintToMapOfPtrUint64,
		"float32": groupByUintToMapOfPtrFloat32,
		"float64": groupByUintToMapOfPtrFloat64,
	},
	"uint8": {
		"string":  groupByUint8ToMapOfString,
		"bool":    groupByUint8ToMapOfBool,
		"int":     groupByUint8ToMapOfInt,
		"int8":    groupByUint8ToMapOfInt8,
		"int16":   groupByUint8ToMapOfInt16,
		"int32":   groupByUint8ToMapOfInt32,
		"int64":   groupByUint8ToMapOfInt64,
		"uint":    groupByUint8ToMapOfUint,
		"uint8":   groupByUint8ToMapOfUint8,
		"uint16":  groupByUint8ToMapOfUint16,
		"uint32":  groupByUint8ToMapOfUint32,
		"uint64":  groupByUint8ToMapOfUint64,
		"float32": groupByUint8ToMapOfFloat32,
		"float64": groupByUint8ToMapOfFloat64,
	},
	"*uint8": {
		"string":  groupByUint8ToMapOfPtrString,
		"bool":    groupByUint8ToMapOfPtrBool,
		"int":     groupByUint8ToMapOfPtrInt,
		"int8":    groupByUint8ToMapOfPtrInt8,
		"int16":   groupByUint8ToMapOfPtrInt16,
		"int32":   groupByUint8ToMapOfPtrInt32,
		"int64":   groupByUint8ToMapOfPtrInt64,
		"uint":    groupByUint8ToMapOfPtrUint,
		"uint8":   groupByUint8ToMapOfPtrUint8,
		"uint16":  groupByUint8ToMapOfPtrUint16,
		"uint32":  groupByUint8ToMapOfPtrUint32,
		"uint64":  groupByUint8ToMapOfPtrUint64,
		"float32": groupByUint8ToMapOfPtrFloat32,
		"float64": groupByUint8ToMapOfPtrFloat64,
	},
	"uint16": {
		"string":  groupByUint16ToMapOfString,
		"bool":    groupByUint16ToMapOfBool,
		"int":     groupByUint16ToMapOfInt,
		"int8":    groupByUint16ToMapOfInt8,
		"int16":   groupByUint16ToMapOfInt16,
		"int32":   groupByUint16ToMapOfInt32,
		"int64":   groupByUint16ToMapOfInt64,
		"uint":    groupByUint16ToMapOfUint,
		"uint8":   groupByUint16ToMapOfUint8,
		"uint16":  groupByUint16ToMapOfUint16,
		"uint32":  groupByUint16ToMapOfUint32,
		"uint64":  groupByUint16ToMapOfUint64,
		"float32": groupByUint16ToMapOfFloat32,
		"float64": groupByUint16ToMapOfFloat64,
	},
	"*uint16": {
		"string":  groupByUint16ToMapOfPtrString,
		"bool":    groupByUint16ToMapOfPtrBool,
		"int":     groupByUint16ToMapOfPtrInt,
		"int8":    groupByUint16ToMapOfPtrInt8,
		"int16":   groupByUint16ToMapOfPtrInt16,
		"int32":   groupByUint16ToMapOfPtrInt32,
		"int64":   groupByUint16ToMapOfPtrInt64,
		"uint":    groupByUint16ToMapOfPtrUint,
		"uint8":   groupByUint16ToMapOfPtrUint8,
		"uint16":  groupByUint16ToMapOfPtrUint16,
		"uint32":  groupByUint16ToMapOfPtrUint32,
		"uint64":  groupByUint16ToMapOfPtrUint64,
		"float32": groupByUint16ToMapOfPtrFloat32,
		"float64": groupByUint16ToMapOfPtrFloat64,
	},
	"uint32": {
		"string":  groupByUint32ToMapOfString,
		"bool":    groupByUint32ToMapOfBool,
		"int":     groupByUint32ToMapOfInt,
		"int8":    groupByUint32ToMapOfInt8,
		"int16":   groupByUint32ToMapOfInt16,
		"int32":   groupByUint32ToMapOfInt32,
		"int64":   groupByUint32ToMapOfInt64,
		"uint":    groupByUint32ToMapOfUint,
		"uint8":   groupByUint32ToMapOfUint8,
		"uint16":  groupByUint32ToMapOfUint16,
		"uint32":  groupByUint32ToMapOfUint32,
		"uint64":  groupByUint32ToMapOfUint64,
		"float32": groupByUint32ToMapOfFloat32,
		"float64": groupByUint32ToMapOfFloat64,
	},
	"*uint32": {
		"string":  groupByUint32ToMapOfPtrString,
		"bool":    groupByUint32ToMapOfPtrBool,
		"int":     groupByUint32ToMapOfPtrInt,
		"int8":    groupByUint32ToMapOfPtrInt8,
		"int16":   groupByUint32ToMapOfPtrInt16,
		"int32":   groupByUint32ToMapOfPtrInt32,
		"int64":   groupByUint32ToMapOfPtrInt64,
		"uint":    groupByUint32ToMapOfPtrUint,
		"uint8":   groupByUint32ToMapOfPtrUint8,
		"uint16":  groupByUint32ToMapOfPtrUint16,
		"uint32":  groupByUint32ToMapOfPtrUint32,
		"uint64":  groupByUint32ToMapOfPtrUint64,
		"float32": groupByUint32ToMapOfPtrFloat32,
		"float64": groupByUint32ToMapOfPtrFloat64,
	},
	"uint64": {
		"string":  groupByUint64ToMapOfString,
		"bool":    groupByUint64ToMapOfBool,
		"int":     groupByUint64ToMapOfInt,
		"int8":    groupByUint64ToMapOfInt8,
		"int16":   groupByUint64ToMapOfInt16,
		"int32":   groupByUint64ToMapOfInt32,
		"int64":   groupByUint64ToMapOfInt64,
		"uint":    groupByUint64ToMapOfUint,
		"uint8":   groupByUint64ToMapOfUint8,
		"uint16":  groupByUint64ToMapOfUint16,
		"uint32":  groupByUint64ToMapOfUint32,
		"uint64":  groupByUint64ToMapOfUint64,
		"float32": groupByUint64ToMapOfFloat32,
		"float64": groupByUint64ToMapOfFloat64,
	},
	"*uint64": {
		"string":  groupByUint64ToMapOfPtrString,
		"bool":    groupByUint64ToMapOfPtrBool,
		"int":     groupByUint64ToMapOfPtrInt,
		"int8":    groupByUint64ToMapOfPtrInt8,
		"int16":   groupByUint64ToMapOfPtrInt16,
		"int32":   groupByUint64ToMapOfPtrInt32,
		"int64":   groupByUint64ToMapOfPtrInt64,
		"uint":    groupByUint64ToMapOfPtrUint,
		"uint8":   groupByUint64ToMapOfPtrUint8,
		"uint16":  groupByUint64ToMapOfPtrUint16,
		"uint32":  groupByUint64ToMapOfPtrUint32,
		"uint64":  groupByUint64ToMapOfPtrUint64,
		"float32": groupByUint64ToMapOfPtrFloat32,
		"float64": groupByUint64ToMapOfPtrFloat64,
	},
	"float32": {
		"string":  groupByFloat32ToMapOfString,
		"bool":    groupByFloat32ToMapOfBool,
		"int":     groupByFloat32ToMapOfInt,
		"int8":    groupByFloat32ToMapOfInt8,
		"int16":   groupByFloat32ToMapOfInt16,
		"int32":   groupByFloat32ToMapOfInt32,
		"int64":   groupByFloat32ToMapOfInt64,
		"uint":    groupByFloat32ToMapOfUint,
		"uint8":   groupByFloat32ToMapOfUint8,
		"uint16":  groupByFloat32ToMapOfUint16,
		"uint32":  groupByFloat32ToMapOfUint32,
		"uint64":  groupByFloat32ToMapOfUint64,
		"float32": groupByFloat32ToMapOfFloat32,
		"float64": groupByFloat32ToMapOfFloat64,
	},
	"*float32": {
		"string":  groupByFloat32ToMapOfPtrString,
		"bool":    groupByFloat32ToMapOfPtrBool,
		"int":     groupByFloat32ToMapOfPtrInt,
		"int8":    groupByFloat32ToMapOfPtrInt8,
		"int16":   groupByFloat32ToMapOfPtrInt16,
		"int32":   groupByFloat32ToMapOfPtrInt32,
		"int64":   groupByFloat32ToMapOfPtrInt64,
		"uint":    groupByFloat32ToMapOfPtrUint,
		"uint8":   groupByFloat32ToMapOfPtrUint8,
		"uint16":  groupByFloat32ToMapOfPtrUint16,
		"uint32":  groupByFloat32ToMapOfPtrUint32,
		"uint64":  groupByFloat32ToMapOfPtrUint64,
		"float32": groupByFloat32ToMapOfPtrFloat32,
		"float64": groupByFloat32ToMapOfPtrFloat64,
	},
	"float64": {
		"string":  groupByFloat64ToMapOfString,
		"bool":    groupByFloat64ToMapOfBool,
		"int":     groupByFloat64ToMapOfInt,
		"int8":    groupByFloat64ToMapOfInt8,
		"int16":   groupByFloat64ToMapOfInt16,
		"int32":   groupByFloat64ToMapOfInt32,
		"int64":   groupByFloat64ToMapOfInt64,
		"uint":    groupByFloat64ToMapOfUint,
		"uint8":   groupByFloat64ToMapOfUint8,
		"uint16":  groupByFloat64ToMapOfUint16,
		"uint32":  groupByFloat64ToMapOfUint32,
		"uint64":  groupByFloat64ToMapOfUint64,
		"float32": groupByFloat64ToMapOfFloat32,
		"float64": groupByFloat64ToMapOfFloat64,
	},
	"*float64": {
		"string":  groupByFloat64ToMapOfPtrString,
		"bool":    groupByFloat64ToMapOfPtrBool,
		"int":     groupByFloat64ToMapOfPtrInt,
		"int8":    groupByFloat64ToMapOfPtrInt8,
		"int16":   groupByFloat64ToMapOfPtrInt16,
		"int32":   groupByFloat64ToMapOfPtrInt32,
		"int64":   groupByFloat64ToMapOfPtrInt64,
		"uint":    groupByFloat64ToMapOfPtrUint,
		"uint8":   groupByFloat64ToMapOfPtrUint8,
		"uint16":  groupByFloat64ToMapOfPtrUint16,
		"uint32":  groupByFloat64ToMapOfPtrUint32,
		"uint64":  groupByFloat64ToMapOfPtrUint64,
		"float32": groupByFloat64ToMapOfPtrFloat32,
		"float64": groupByFloat64ToMapOfPtrFloat64,
	},
}

func dispatch(items reflect.Value, function interface{}, info *groupByInfo) (bool, interface{}) {
	input := info.fnInputType.String()
	output := info.fnOutputType.String()
	if inputVal, ok := dispatcher[input]; ok {
		if outputVal, ok := inputVal[output]; ok {
			return true, outputVal(items, function)
		}
	}
	return false, nil
}

func groupByStringToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[string][]string)
	fn := function.(func(string) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[string][]*string)
	fn := function.(func(*string) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[bool][]string)
	fn := function.(func(string) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[bool][]*string)
	fn := function.(func(*string) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[int][]string)
	fn := function.(func(string) int)
	fmt.Println()
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[int][]*string)
	fn := function.(func(*string) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[int8][]string)
	fn := function.(func(string) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[int8][]*string)
	fn := function.(func(*string) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[int16][]string)
	fn := function.(func(string) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[int16][]*string)
	fn := function.(func(*string) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[int32][]string)
	fn := function.(func(string) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[int32][]*string)
	fn := function.(func(*string) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[int64][]string)
	fn := function.(func(string) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[int64][]*string)
	fn := function.(func(*string) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[uint][]string)
	fn := function.(func(string) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[uint][]*string)
	fn := function.(func(*string) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[uint8][]string)
	fn := function.(func(string) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[uint8][]*string)
	fn := function.(func(*string) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[uint16][]string)
	fn := function.(func(string) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[uint16][]*string)
	fn := function.(func(*string) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[uint32][]string)
	fn := function.(func(string) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[uint32][]*string)
	fn := function.(func(*string) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[uint64][]string)
	fn := function.(func(string) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[uint64][]*string)
	fn := function.(func(*string) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[float32][]string)
	fn := function.(func(string) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[float32][]*string)
	fn := function.(func(*string) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByStringToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]string)
	output := make(map[float64][]string)
	fn := function.(func(string) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []string{input[i]}
	}
	return output
}

func groupByStringToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*string)
	output := make(map[float64][]*string)
	fn := function.(func(*string) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*string{input[i]}
	}
	return output
}

func groupByBoolToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[string][]bool)
	fn := function.(func(bool) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[string][]*bool)
	fn := function.(func(*bool) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[bool][]bool)
	fn := function.(func(bool) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[bool][]*bool)
	fn := function.(func(*bool) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[int][]bool)
	fn := function.(func(bool) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[int][]*bool)
	fn := function.(func(*bool) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[int8][]bool)
	fn := function.(func(bool) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[int8][]*bool)
	fn := function.(func(*bool) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[int16][]bool)
	fn := function.(func(bool) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[int16][]*bool)
	fn := function.(func(*bool) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[int32][]bool)
	fn := function.(func(bool) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[int32][]*bool)
	fn := function.(func(*bool) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[int64][]bool)
	fn := function.(func(bool) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[int64][]*bool)
	fn := function.(func(*bool) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[uint][]bool)
	fn := function.(func(bool) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[uint][]*bool)
	fn := function.(func(*bool) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[uint8][]bool)
	fn := function.(func(bool) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[uint8][]*bool)
	fn := function.(func(*bool) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[uint16][]bool)
	fn := function.(func(bool) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[uint16][]*bool)
	fn := function.(func(*bool) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[uint32][]bool)
	fn := function.(func(bool) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[uint32][]*bool)
	fn := function.(func(*bool) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[uint64][]bool)
	fn := function.(func(bool) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[uint64][]*bool)
	fn := function.(func(*bool) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[float32][]bool)
	fn := function.(func(bool) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[float32][]*bool)
	fn := function.(func(*bool) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]bool)
	output := make(map[float64][]bool)
	fn := function.(func(bool) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []bool{input[i]}
	}
	return output
}

func groupByBoolToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*bool)
	output := make(map[float64][]*bool)
	fn := function.(func(*bool) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*bool{input[i]}
	}
	return output
}

func groupByIntToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[string][]int)
	fn := function.(func(int) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[string][]*int)
	fn := function.(func(*int) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[bool][]int)
	fn := function.(func(int) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[bool][]*int)
	fn := function.(func(*int) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[int][]int)
	fn := function.(func(int) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[int][]*int)
	fn := function.(func(*int) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[int8][]int)
	fn := function.(func(int) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[int8][]*int)
	fn := function.(func(*int) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[int16][]int)
	fn := function.(func(int) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[int16][]*int)
	fn := function.(func(*int) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[int32][]int)
	fn := function.(func(int) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[int32][]*int)
	fn := function.(func(*int) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[int64][]int)
	fn := function.(func(int) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[int64][]*int)
	fn := function.(func(*int) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[uint][]int)
	fn := function.(func(int) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[uint][]*int)
	fn := function.(func(*int) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[uint8][]int)
	fn := function.(func(int) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[uint8][]*int)
	fn := function.(func(*int) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[uint16][]int)
	fn := function.(func(int) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[uint16][]*int)
	fn := function.(func(*int) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[uint32][]int)
	fn := function.(func(int) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[uint32][]*int)
	fn := function.(func(*int) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[uint64][]int)
	fn := function.(func(int) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[uint64][]*int)
	fn := function.(func(*int) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[float32][]int)
	fn := function.(func(int) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[float32][]*int)
	fn := function.(func(*int) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByIntToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int)
	output := make(map[float64][]int)
	fn := function.(func(int) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int{input[i]}
	}
	return output
}

func groupByIntToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int)
	output := make(map[float64][]*int)
	fn := function.(func(*int) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int{input[i]}
	}
	return output
}

func groupByInt8ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[string][]int8)
	fn := function.(func(int8) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[string][]*int8)
	fn := function.(func(*int8) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[bool][]int8)
	fn := function.(func(int8) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[bool][]*int8)
	fn := function.(func(*int8) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[int][]int8)
	fn := function.(func(int8) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[int][]*int8)
	fn := function.(func(*int8) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[int8][]int8)
	fn := function.(func(int8) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[int8][]*int8)
	fn := function.(func(*int8) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[int16][]int8)
	fn := function.(func(int8) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[int16][]*int8)
	fn := function.(func(*int8) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[int32][]int8)
	fn := function.(func(int8) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[int32][]*int8)
	fn := function.(func(*int8) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[int64][]int8)
	fn := function.(func(int8) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[int64][]*int8)
	fn := function.(func(*int8) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[uint][]int8)
	fn := function.(func(int8) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[uint][]*int8)
	fn := function.(func(*int8) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[uint8][]int8)
	fn := function.(func(int8) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[uint8][]*int8)
	fn := function.(func(*int8) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[uint16][]int8)
	fn := function.(func(int8) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[uint16][]*int8)
	fn := function.(func(*int8) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[uint32][]int8)
	fn := function.(func(int8) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[uint32][]*int8)
	fn := function.(func(*int8) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[uint64][]int8)
	fn := function.(func(int8) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[uint64][]*int8)
	fn := function.(func(*int8) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[float32][]int8)
	fn := function.(func(int8) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[float32][]*int8)
	fn := function.(func(*int8) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int8)
	output := make(map[float64][]int8)
	fn := function.(func(int8) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int8{input[i]}
	}
	return output
}

func groupByInt8ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int8)
	output := make(map[float64][]*int8)
	fn := function.(func(*int8) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int8{input[i]}
	}
	return output
}

func groupByInt16ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[string][]int16)
	fn := function.(func(int16) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[string][]*int16)
	fn := function.(func(*int16) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[bool][]int16)
	fn := function.(func(int16) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[bool][]*int16)
	fn := function.(func(*int16) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[int][]int16)
	fn := function.(func(int16) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[int][]*int16)
	fn := function.(func(*int16) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[int8][]int16)
	fn := function.(func(int16) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[int8][]*int16)
	fn := function.(func(*int16) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[int16][]int16)
	fn := function.(func(int16) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[int16][]*int16)
	fn := function.(func(*int16) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[int32][]int16)
	fn := function.(func(int16) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[int32][]*int16)
	fn := function.(func(*int16) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[int64][]int16)
	fn := function.(func(int16) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[int64][]*int16)
	fn := function.(func(*int16) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[uint][]int16)
	fn := function.(func(int16) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[uint][]*int16)
	fn := function.(func(*int16) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[uint8][]int16)
	fn := function.(func(int16) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[uint8][]*int16)
	fn := function.(func(*int16) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[uint16][]int16)
	fn := function.(func(int16) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[uint16][]*int16)
	fn := function.(func(*int16) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[uint32][]int16)
	fn := function.(func(int16) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[uint32][]*int16)
	fn := function.(func(*int16) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[uint64][]int16)
	fn := function.(func(int16) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[uint64][]*int16)
	fn := function.(func(*int16) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[float32][]int16)
	fn := function.(func(int16) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[float32][]*int16)
	fn := function.(func(*int16) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int16)
	output := make(map[float64][]int16)
	fn := function.(func(int16) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int16{input[i]}
	}
	return output
}

func groupByInt16ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int16)
	output := make(map[float64][]*int16)
	fn := function.(func(*int16) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int16{input[i]}
	}
	return output
}

func groupByInt32ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[string][]int32)
	fn := function.(func(int32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[string][]*int32)
	fn := function.(func(*int32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[bool][]int32)
	fn := function.(func(int32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[bool][]*int32)
	fn := function.(func(*int32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[int][]int32)
	fn := function.(func(int32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[int][]*int32)
	fn := function.(func(*int32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[int8][]int32)
	fn := function.(func(int32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[int8][]*int32)
	fn := function.(func(*int32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[int16][]int32)
	fn := function.(func(int32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[int16][]*int32)
	fn := function.(func(*int32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[int32][]int32)
	fn := function.(func(int32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[int32][]*int32)
	fn := function.(func(*int32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[int64][]int32)
	fn := function.(func(int32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[int64][]*int32)
	fn := function.(func(*int32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[uint][]int32)
	fn := function.(func(int32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[uint][]*int32)
	fn := function.(func(*int32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[uint8][]int32)
	fn := function.(func(int32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[uint8][]*int32)
	fn := function.(func(*int32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[uint16][]int32)
	fn := function.(func(int32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[uint16][]*int32)
	fn := function.(func(*int32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[uint32][]int32)
	fn := function.(func(int32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[uint32][]*int32)
	fn := function.(func(*int32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[uint64][]int32)
	fn := function.(func(int32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[uint64][]*int32)
	fn := function.(func(*int32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[float32][]int32)
	fn := function.(func(int32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[float32][]*int32)
	fn := function.(func(*int32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int32)
	output := make(map[float64][]int32)
	fn := function.(func(int32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int32{input[i]}
	}
	return output
}

func groupByInt32ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int32)
	output := make(map[float64][]*int32)
	fn := function.(func(*int32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int32{input[i]}
	}
	return output
}

func groupByInt64ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[string][]int64)
	fn := function.(func(int64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[string][]*int64)
	fn := function.(func(*int64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[bool][]int64)
	fn := function.(func(int64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[bool][]*int64)
	fn := function.(func(*int64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[int][]int64)
	fn := function.(func(int64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[int][]*int64)
	fn := function.(func(*int64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[int8][]int64)
	fn := function.(func(int64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[int8][]*int64)
	fn := function.(func(*int64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[int16][]int64)
	fn := function.(func(int64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[int16][]*int64)
	fn := function.(func(*int64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[int32][]int64)
	fn := function.(func(int64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[int32][]*int64)
	fn := function.(func(*int64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[int64][]int64)
	fn := function.(func(int64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[int64][]*int64)
	fn := function.(func(*int64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[uint][]int64)
	fn := function.(func(int64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[uint][]*int64)
	fn := function.(func(*int64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[uint8][]int64)
	fn := function.(func(int64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[uint8][]*int64)
	fn := function.(func(*int64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[uint16][]int64)
	fn := function.(func(int64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[uint16][]*int64)
	fn := function.(func(*int64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[uint32][]int64)
	fn := function.(func(int64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[uint32][]*int64)
	fn := function.(func(*int64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[uint64][]int64)
	fn := function.(func(int64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[uint64][]*int64)
	fn := function.(func(*int64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[float32][]int64)
	fn := function.(func(int64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[float32][]*int64)
	fn := function.(func(*int64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]int64)
	output := make(map[float64][]int64)
	fn := function.(func(int64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []int64{input[i]}
	}
	return output
}

func groupByInt64ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*int64)
	output := make(map[float64][]*int64)
	fn := function.(func(*int64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*int64{input[i]}
	}
	return output
}

func groupByUintToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[string][]uint)
	fn := function.(func(uint) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[string][]*uint)
	fn := function.(func(*uint) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[bool][]uint)
	fn := function.(func(uint) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[bool][]*uint)
	fn := function.(func(*uint) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[int][]uint)
	fn := function.(func(uint) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[int][]*uint)
	fn := function.(func(*uint) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[int8][]uint)
	fn := function.(func(uint) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[int8][]*uint)
	fn := function.(func(*uint) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[int16][]uint)
	fn := function.(func(uint) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[int16][]*uint)
	fn := function.(func(*uint) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[int32][]uint)
	fn := function.(func(uint) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[int32][]*uint)
	fn := function.(func(*uint) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[int64][]uint)
	fn := function.(func(uint) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[int64][]*uint)
	fn := function.(func(*uint) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[uint][]uint)
	fn := function.(func(uint) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[uint][]*uint)
	fn := function.(func(*uint) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[uint8][]uint)
	fn := function.(func(uint) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[uint8][]*uint)
	fn := function.(func(*uint) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[uint16][]uint)
	fn := function.(func(uint) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[uint16][]*uint)
	fn := function.(func(*uint) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[uint32][]uint)
	fn := function.(func(uint) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[uint32][]*uint)
	fn := function.(func(*uint) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[uint64][]uint)
	fn := function.(func(uint) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[uint64][]*uint)
	fn := function.(func(*uint) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[float32][]uint)
	fn := function.(func(uint) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[float32][]*uint)
	fn := function.(func(*uint) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUintToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint)
	output := make(map[float64][]uint)
	fn := function.(func(uint) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint{input[i]}
	}
	return output
}

func groupByUintToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint)
	output := make(map[float64][]*uint)
	fn := function.(func(*uint) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint{input[i]}
	}
	return output
}

func groupByUint8ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[string][]uint8)
	fn := function.(func(uint8) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[string][]*uint8)
	fn := function.(func(*uint8) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[bool][]uint8)
	fn := function.(func(uint8) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[bool][]*uint8)
	fn := function.(func(*uint8) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[int][]uint8)
	fn := function.(func(uint8) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[int][]*uint8)
	fn := function.(func(*uint8) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[int8][]uint8)
	fn := function.(func(uint8) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[int8][]*uint8)
	fn := function.(func(*uint8) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[int16][]uint8)
	fn := function.(func(uint8) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[int16][]*uint8)
	fn := function.(func(*uint8) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[int32][]uint8)
	fn := function.(func(uint8) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[int32][]*uint8)
	fn := function.(func(*uint8) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[int64][]uint8)
	fn := function.(func(uint8) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[int64][]*uint8)
	fn := function.(func(*uint8) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[uint][]uint8)
	fn := function.(func(uint8) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[uint][]*uint8)
	fn := function.(func(*uint8) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[uint8][]uint8)
	fn := function.(func(uint8) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[uint8][]*uint8)
	fn := function.(func(*uint8) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[uint16][]uint8)
	fn := function.(func(uint8) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[uint16][]*uint8)
	fn := function.(func(*uint8) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[uint32][]uint8)
	fn := function.(func(uint8) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[uint32][]*uint8)
	fn := function.(func(*uint8) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[uint64][]uint8)
	fn := function.(func(uint8) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[uint64][]*uint8)
	fn := function.(func(*uint8) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[float32][]uint8)
	fn := function.(func(uint8) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[float32][]*uint8)
	fn := function.(func(*uint8) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint8)
	output := make(map[float64][]uint8)
	fn := function.(func(uint8) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint8{input[i]}
	}
	return output
}

func groupByUint8ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint8)
	output := make(map[float64][]*uint8)
	fn := function.(func(*uint8) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint8{input[i]}
	}
	return output
}

func groupByUint16ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[string][]uint16)
	fn := function.(func(uint16) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[string][]*uint16)
	fn := function.(func(*uint16) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[bool][]uint16)
	fn := function.(func(uint16) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[bool][]*uint16)
	fn := function.(func(*uint16) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[int][]uint16)
	fn := function.(func(uint16) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[int][]*uint16)
	fn := function.(func(*uint16) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[int8][]uint16)
	fn := function.(func(uint16) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[int8][]*uint16)
	fn := function.(func(*uint16) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[int16][]uint16)
	fn := function.(func(uint16) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[int16][]*uint16)
	fn := function.(func(*uint16) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[int32][]uint16)
	fn := function.(func(uint16) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[int32][]*uint16)
	fn := function.(func(*uint16) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[int64][]uint16)
	fn := function.(func(uint16) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[int64][]*uint16)
	fn := function.(func(*uint16) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[uint][]uint16)
	fn := function.(func(uint16) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[uint][]*uint16)
	fn := function.(func(*uint16) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[uint8][]uint16)
	fn := function.(func(uint16) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[uint8][]*uint16)
	fn := function.(func(*uint16) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[uint16][]uint16)
	fn := function.(func(uint16) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[uint16][]*uint16)
	fn := function.(func(*uint16) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[uint32][]uint16)
	fn := function.(func(uint16) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[uint32][]*uint16)
	fn := function.(func(*uint16) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[uint64][]uint16)
	fn := function.(func(uint16) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[uint64][]*uint16)
	fn := function.(func(*uint16) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[float32][]uint16)
	fn := function.(func(uint16) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[float32][]*uint16)
	fn := function.(func(*uint16) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint16)
	output := make(map[float64][]uint16)
	fn := function.(func(uint16) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint16{input[i]}
	}
	return output
}

func groupByUint16ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint16)
	output := make(map[float64][]*uint16)
	fn := function.(func(*uint16) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint16{input[i]}
	}
	return output
}

func groupByUint32ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[string][]uint32)
	fn := function.(func(uint32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[string][]*uint32)
	fn := function.(func(*uint32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[bool][]uint32)
	fn := function.(func(uint32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[bool][]*uint32)
	fn := function.(func(*uint32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[int][]uint32)
	fn := function.(func(uint32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[int][]*uint32)
	fn := function.(func(*uint32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[int8][]uint32)
	fn := function.(func(uint32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[int8][]*uint32)
	fn := function.(func(*uint32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[int16][]uint32)
	fn := function.(func(uint32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[int16][]*uint32)
	fn := function.(func(*uint32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[int32][]uint32)
	fn := function.(func(uint32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[int32][]*uint32)
	fn := function.(func(*uint32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[int64][]uint32)
	fn := function.(func(uint32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[int64][]*uint32)
	fn := function.(func(*uint32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[uint][]uint32)
	fn := function.(func(uint32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[uint][]*uint32)
	fn := function.(func(*uint32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[uint8][]uint32)
	fn := function.(func(uint32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[uint8][]*uint32)
	fn := function.(func(*uint32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[uint16][]uint32)
	fn := function.(func(uint32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[uint16][]*uint32)
	fn := function.(func(*uint32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[uint32][]uint32)
	fn := function.(func(uint32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[uint32][]*uint32)
	fn := function.(func(*uint32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[uint64][]uint32)
	fn := function.(func(uint32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[uint64][]*uint32)
	fn := function.(func(*uint32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[float32][]uint32)
	fn := function.(func(uint32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[float32][]*uint32)
	fn := function.(func(*uint32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint32)
	output := make(map[float64][]uint32)
	fn := function.(func(uint32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint32{input[i]}
	}
	return output
}

func groupByUint32ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint32)
	output := make(map[float64][]*uint32)
	fn := function.(func(*uint32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint32{input[i]}
	}
	return output
}

func groupByUint64ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[string][]uint64)
	fn := function.(func(uint64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[string][]*uint64)
	fn := function.(func(*uint64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[bool][]uint64)
	fn := function.(func(uint64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[bool][]*uint64)
	fn := function.(func(*uint64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[int][]uint64)
	fn := function.(func(uint64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[int][]*uint64)
	fn := function.(func(*uint64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[int8][]uint64)
	fn := function.(func(uint64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[int8][]*uint64)
	fn := function.(func(*uint64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[int16][]uint64)
	fn := function.(func(uint64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[int16][]*uint64)
	fn := function.(func(*uint64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[int32][]uint64)
	fn := function.(func(uint64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[int32][]*uint64)
	fn := function.(func(*uint64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[int64][]uint64)
	fn := function.(func(uint64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[int64][]*uint64)
	fn := function.(func(*uint64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[uint][]uint64)
	fn := function.(func(uint64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[uint][]*uint64)
	fn := function.(func(*uint64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[uint8][]uint64)
	fn := function.(func(uint64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[uint8][]*uint64)
	fn := function.(func(*uint64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[uint16][]uint64)
	fn := function.(func(uint64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[uint16][]*uint64)
	fn := function.(func(*uint64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[uint32][]uint64)
	fn := function.(func(uint64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[uint32][]*uint64)
	fn := function.(func(*uint64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[uint64][]uint64)
	fn := function.(func(uint64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[uint64][]*uint64)
	fn := function.(func(*uint64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[float32][]uint64)
	fn := function.(func(uint64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[float32][]*uint64)
	fn := function.(func(*uint64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]uint64)
	output := make(map[float64][]uint64)
	fn := function.(func(uint64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []uint64{input[i]}
	}
	return output
}

func groupByUint64ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*uint64)
	output := make(map[float64][]*uint64)
	fn := function.(func(*uint64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*uint64{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[string][]float32)
	fn := function.(func(float32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[string][]*float32)
	fn := function.(func(*float32) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[bool][]float32)
	fn := function.(func(float32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[bool][]*float32)
	fn := function.(func(*float32) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[int][]float32)
	fn := function.(func(float32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[int][]*float32)
	fn := function.(func(*float32) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[int8][]float32)
	fn := function.(func(float32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[int8][]*float32)
	fn := function.(func(*float32) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[int16][]float32)
	fn := function.(func(float32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[int16][]*float32)
	fn := function.(func(*float32) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[int32][]float32)
	fn := function.(func(float32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[int32][]*float32)
	fn := function.(func(*float32) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[int64][]float32)
	fn := function.(func(float32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[int64][]*float32)
	fn := function.(func(*float32) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[uint][]float32)
	fn := function.(func(float32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[uint][]*float32)
	fn := function.(func(*float32) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[uint8][]float32)
	fn := function.(func(float32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[uint8][]*float32)
	fn := function.(func(*float32) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[uint16][]float32)
	fn := function.(func(float32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[uint16][]*float32)
	fn := function.(func(*float32) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[uint32][]float32)
	fn := function.(func(float32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[uint32][]*float32)
	fn := function.(func(*float32) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[uint64][]float32)
	fn := function.(func(float32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[uint64][]*float32)
	fn := function.(func(*float32) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[float32][]float32)
	fn := function.(func(float32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[float32][]*float32)
	fn := function.(func(*float32) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float32)
	output := make(map[float64][]float32)
	fn := function.(func(float32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float32{input[i]}
	}
	return output
}

func groupByFloat32ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float32)
	output := make(map[float64][]*float32)
	fn := function.(func(*float32) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float32{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[string][]float64)
	fn := function.(func(float64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[string][]*float64)
	fn := function.(func(*float64) string)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[bool][]float64)
	fn := function.(func(float64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[bool][]*float64)
	fn := function.(func(*float64) bool)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[int][]float64)
	fn := function.(func(float64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[int][]*float64)
	fn := function.(func(*float64) int)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[int8][]float64)
	fn := function.(func(float64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[int8][]*float64)
	fn := function.(func(*float64) int8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[int16][]float64)
	fn := function.(func(float64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[int16][]*float64)
	fn := function.(func(*float64) int16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[int32][]float64)
	fn := function.(func(float64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[int32][]*float64)
	fn := function.(func(*float64) int32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[int64][]float64)
	fn := function.(func(float64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[int64][]*float64)
	fn := function.(func(*float64) int64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[uint][]float64)
	fn := function.(func(float64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[uint][]*float64)
	fn := function.(func(*float64) uint)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[uint8][]float64)
	fn := function.(func(float64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[uint8][]*float64)
	fn := function.(func(*float64) uint8)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[uint16][]float64)
	fn := function.(func(float64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[uint16][]*float64)
	fn := function.(func(*float64) uint16)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[uint32][]float64)
	fn := function.(func(float64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[uint32][]*float64)
	fn := function.(func(*float64) uint32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[uint64][]float64)
	fn := function.(func(float64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[uint64][]*float64)
	fn := function.(func(*float64) uint64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[float32][]float64)
	fn := function.(func(float64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[float32][]*float64)
	fn := function.(func(*float64) float32)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]float64)
	output := make(map[float64][]float64)
	fn := function.(func(float64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []float64{input[i]}
	}
	return output
}

func groupByFloat64ToMapOfPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	input := itemsValue.Interface().([]*float64)
	output := make(map[float64][]*float64)
	fn := function.(func(*float64) float64)
	for i := 0; i < len(input); i++ {
		res := fn(input[i])
		if val, ok := output[res]; ok {
			output[res] = append(val, input[i])
			continue
		}
		output[res] = []*float64{input[i]}
	}
	return output
}
