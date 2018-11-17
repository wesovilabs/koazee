
package _map

import "reflect"


type dispatchFunction func(items *reflect.Value, fn interface{}, info *mapInfo) interface{}

var dispatcher = map[string]map[string]dispatchFunction{
	"string": {
		"string": mapStringToString,
		"bool": mapStringToBool,
		"int": mapStringToInt,
		"int8": mapStringToInt8,
		"int16": mapStringToInt16,
		"int32": mapStringToInt32,
		"int64": mapStringToInt64,
		"uint": mapStringToUint,
		"uint8": mapStringToUint8,
		"uint16": mapStringToUint16,
		"uint32": mapStringToUint32,
		"uint64": mapStringToUint64,
		"float32": mapStringToFloat32,
		"float64": mapStringToFloat64,
		"*string": mapStringToPtrString,
		"*bool": mapStringToPtrBool,
		"*int": mapStringToPtrInt,
		"*int8": mapStringToPtrInt8,
		"*int16": mapStringToPtrInt16,
		"*int32": mapStringToPtrInt32,
		"*int64": mapStringToPtrInt64,
		"*uint": mapStringToPtrUint,
		"*uint8": mapStringToPtrUint8,
		"*uint16": mapStringToPtrUint16,
		"*uint32": mapStringToPtrUint32,
		"*uint64": mapStringToPtrUint64,
		"*float32": mapStringToPtrFloat32,
		"*float64": mapStringToPtrFloat64,
	},
	"*string": {
		"string": mapPtrStringToString,
		"bool": mapPtrStringToBool,
		"int": mapPtrStringToInt,
		"int8": mapPtrStringToInt8,
		"int16": mapPtrStringToInt16,
		"int32": mapPtrStringToInt32,
		"int64": mapPtrStringToInt64,
		"uint": mapPtrStringToUint,
		"uint8": mapPtrStringToUint8,
		"uint16": mapPtrStringToUint16,
		"uint32": mapPtrStringToUint32,
		"uint64": mapPtrStringToUint64,
		"float32": mapPtrStringToFloat32,
		"float64": mapPtrStringToFloat64,
		"*string": mapPtrStringToPtrString,
		"*bool": mapPtrStringToPtrBool,
		"*int": mapPtrStringToPtrInt,
		"*int8": mapPtrStringToPtrInt8,
		"*int16": mapPtrStringToPtrInt16,
		"*int32": mapPtrStringToPtrInt32,
		"*int64": mapPtrStringToPtrInt64,
		"*uint": mapPtrStringToPtrUint,
		"*uint8": mapPtrStringToPtrUint8,
		"*uint16": mapPtrStringToPtrUint16,
		"*uint32": mapPtrStringToPtrUint32,
		"*uint64": mapPtrStringToPtrUint64,
		"*float32": mapPtrStringToPtrFloat32,
		"*float64": mapPtrStringToPtrFloat64,
	},
	"bool": {
		"string": mapBoolToString,
		"bool": mapBoolToBool,
		"int": mapBoolToInt,
		"int8": mapBoolToInt8,
		"int16": mapBoolToInt16,
		"int32": mapBoolToInt32,
		"int64": mapBoolToInt64,
		"uint": mapBoolToUint,
		"uint8": mapBoolToUint8,
		"uint16": mapBoolToUint16,
		"uint32": mapBoolToUint32,
		"uint64": mapBoolToUint64,
		"float32": mapBoolToFloat32,
		"float64": mapBoolToFloat64,
		"*string": mapBoolToPtrString,
		"*bool": mapBoolToPtrBool,
		"*int": mapBoolToPtrInt,
		"*int8": mapBoolToPtrInt8,
		"*int16": mapBoolToPtrInt16,
		"*int32": mapBoolToPtrInt32,
		"*int64": mapBoolToPtrInt64,
		"*uint": mapBoolToPtrUint,
		"*uint8": mapBoolToPtrUint8,
		"*uint16": mapBoolToPtrUint16,
		"*uint32": mapBoolToPtrUint32,
		"*uint64": mapBoolToPtrUint64,
		"*float32": mapBoolToPtrFloat32,
		"*float64": mapBoolToPtrFloat64,
	},
	"*bool": {
		"string": mapPtrBoolToString,
		"bool": mapPtrBoolToBool,
		"int": mapPtrBoolToInt,
		"int8": mapPtrBoolToInt8,
		"int16": mapPtrBoolToInt16,
		"int32": mapPtrBoolToInt32,
		"int64": mapPtrBoolToInt64,
		"uint": mapPtrBoolToUint,
		"uint8": mapPtrBoolToUint8,
		"uint16": mapPtrBoolToUint16,
		"uint32": mapPtrBoolToUint32,
		"uint64": mapPtrBoolToUint64,
		"float32": mapPtrBoolToFloat32,
		"float64": mapPtrBoolToFloat64,
		"*string": mapPtrBoolToPtrString,
		"*bool": mapPtrBoolToPtrBool,
		"*int": mapPtrBoolToPtrInt,
		"*int8": mapPtrBoolToPtrInt8,
		"*int16": mapPtrBoolToPtrInt16,
		"*int32": mapPtrBoolToPtrInt32,
		"*int64": mapPtrBoolToPtrInt64,
		"*uint": mapPtrBoolToPtrUint,
		"*uint8": mapPtrBoolToPtrUint8,
		"*uint16": mapPtrBoolToPtrUint16,
		"*uint32": mapPtrBoolToPtrUint32,
		"*uint64": mapPtrBoolToPtrUint64,
		"*float32": mapPtrBoolToPtrFloat32,
		"*float64": mapPtrBoolToPtrFloat64,
	},
	"int": {
		"string": mapIntToString,
		"bool": mapIntToBool,
		"int": mapIntToInt,
		"int8": mapIntToInt8,
		"int16": mapIntToInt16,
		"int32": mapIntToInt32,
		"int64": mapIntToInt64,
		"uint": mapIntToUint,
		"uint8": mapIntToUint8,
		"uint16": mapIntToUint16,
		"uint32": mapIntToUint32,
		"uint64": mapIntToUint64,
		"float32": mapIntToFloat32,
		"float64": mapIntToFloat64,
		"*string": mapIntToPtrString,
		"*bool": mapIntToPtrBool,
		"*int": mapIntToPtrInt,
		"*int8": mapIntToPtrInt8,
		"*int16": mapIntToPtrInt16,
		"*int32": mapIntToPtrInt32,
		"*int64": mapIntToPtrInt64,
		"*uint": mapIntToPtrUint,
		"*uint8": mapIntToPtrUint8,
		"*uint16": mapIntToPtrUint16,
		"*uint32": mapIntToPtrUint32,
		"*uint64": mapIntToPtrUint64,
		"*float32": mapIntToPtrFloat32,
		"*float64": mapIntToPtrFloat64,
	},
	"*int": {
		"string": mapPtrIntToString,
		"bool": mapPtrIntToBool,
		"int": mapPtrIntToInt,
		"int8": mapPtrIntToInt8,
		"int16": mapPtrIntToInt16,
		"int32": mapPtrIntToInt32,
		"int64": mapPtrIntToInt64,
		"uint": mapPtrIntToUint,
		"uint8": mapPtrIntToUint8,
		"uint16": mapPtrIntToUint16,
		"uint32": mapPtrIntToUint32,
		"uint64": mapPtrIntToUint64,
		"float32": mapPtrIntToFloat32,
		"float64": mapPtrIntToFloat64,
		"*string": mapPtrIntToPtrString,
		"*bool": mapPtrIntToPtrBool,
		"*int": mapPtrIntToPtrInt,
		"*int8": mapPtrIntToPtrInt8,
		"*int16": mapPtrIntToPtrInt16,
		"*int32": mapPtrIntToPtrInt32,
		"*int64": mapPtrIntToPtrInt64,
		"*uint": mapPtrIntToPtrUint,
		"*uint8": mapPtrIntToPtrUint8,
		"*uint16": mapPtrIntToPtrUint16,
		"*uint32": mapPtrIntToPtrUint32,
		"*uint64": mapPtrIntToPtrUint64,
		"*float32": mapPtrIntToPtrFloat32,
		"*float64": mapPtrIntToPtrFloat64,
	},
	"int8": {
		"string": mapInt8ToString,
		"bool": mapInt8ToBool,
		"int": mapInt8ToInt,
		"int8": mapInt8ToInt8,
		"int16": mapInt8ToInt16,
		"int32": mapInt8ToInt32,
		"int64": mapInt8ToInt64,
		"uint": mapInt8ToUint,
		"uint8": mapInt8ToUint8,
		"uint16": mapInt8ToUint16,
		"uint32": mapInt8ToUint32,
		"uint64": mapInt8ToUint64,
		"float32": mapInt8ToFloat32,
		"float64": mapInt8ToFloat64,
		"*string": mapInt8ToPtrString,
		"*bool": mapInt8ToPtrBool,
		"*int": mapInt8ToPtrInt,
		"*int8": mapInt8ToPtrInt8,
		"*int16": mapInt8ToPtrInt16,
		"*int32": mapInt8ToPtrInt32,
		"*int64": mapInt8ToPtrInt64,
		"*uint": mapInt8ToPtrUint,
		"*uint8": mapInt8ToPtrUint8,
		"*uint16": mapInt8ToPtrUint16,
		"*uint32": mapInt8ToPtrUint32,
		"*uint64": mapInt8ToPtrUint64,
		"*float32": mapInt8ToPtrFloat32,
		"*float64": mapInt8ToPtrFloat64,
	},
	"*int8": {
		"string": mapPtrInt8ToString,
		"bool": mapPtrInt8ToBool,
		"int": mapPtrInt8ToInt,
		"int8": mapPtrInt8ToInt8,
		"int16": mapPtrInt8ToInt16,
		"int32": mapPtrInt8ToInt32,
		"int64": mapPtrInt8ToInt64,
		"uint": mapPtrInt8ToUint,
		"uint8": mapPtrInt8ToUint8,
		"uint16": mapPtrInt8ToUint16,
		"uint32": mapPtrInt8ToUint32,
		"uint64": mapPtrInt8ToUint64,
		"float32": mapPtrInt8ToFloat32,
		"float64": mapPtrInt8ToFloat64,
		"*string": mapPtrInt8ToPtrString,
		"*bool": mapPtrInt8ToPtrBool,
		"*int": mapPtrInt8ToPtrInt,
		"*int8": mapPtrInt8ToPtrInt8,
		"*int16": mapPtrInt8ToPtrInt16,
		"*int32": mapPtrInt8ToPtrInt32,
		"*int64": mapPtrInt8ToPtrInt64,
		"*uint": mapPtrInt8ToPtrUint,
		"*uint8": mapPtrInt8ToPtrUint8,
		"*uint16": mapPtrInt8ToPtrUint16,
		"*uint32": mapPtrInt8ToPtrUint32,
		"*uint64": mapPtrInt8ToPtrUint64,
		"*float32": mapPtrInt8ToPtrFloat32,
		"*float64": mapPtrInt8ToPtrFloat64,
	},
	"int16": {
		"string": mapInt16ToString,
		"bool": mapInt16ToBool,
		"int": mapInt16ToInt,
		"int8": mapInt16ToInt8,
		"int16": mapInt16ToInt16,
		"int32": mapInt16ToInt32,
		"int64": mapInt16ToInt64,
		"uint": mapInt16ToUint,
		"uint8": mapInt16ToUint8,
		"uint16": mapInt16ToUint16,
		"uint32": mapInt16ToUint32,
		"uint64": mapInt16ToUint64,
		"float32": mapInt16ToFloat32,
		"float64": mapInt16ToFloat64,
		"*string": mapInt16ToPtrString,
		"*bool": mapInt16ToPtrBool,
		"*int": mapInt16ToPtrInt,
		"*int8": mapInt16ToPtrInt8,
		"*int16": mapInt16ToPtrInt16,
		"*int32": mapInt16ToPtrInt32,
		"*int64": mapInt16ToPtrInt64,
		"*uint": mapInt16ToPtrUint,
		"*uint8": mapInt16ToPtrUint8,
		"*uint16": mapInt16ToPtrUint16,
		"*uint32": mapInt16ToPtrUint32,
		"*uint64": mapInt16ToPtrUint64,
		"*float32": mapInt16ToPtrFloat32,
		"*float64": mapInt16ToPtrFloat64,
	},
	"*int16": {
		"string": mapPtrInt16ToString,
		"bool": mapPtrInt16ToBool,
		"int": mapPtrInt16ToInt,
		"int8": mapPtrInt16ToInt8,
		"int16": mapPtrInt16ToInt16,
		"int32": mapPtrInt16ToInt32,
		"int64": mapPtrInt16ToInt64,
		"uint": mapPtrInt16ToUint,
		"uint8": mapPtrInt16ToUint8,
		"uint16": mapPtrInt16ToUint16,
		"uint32": mapPtrInt16ToUint32,
		"uint64": mapPtrInt16ToUint64,
		"float32": mapPtrInt16ToFloat32,
		"float64": mapPtrInt16ToFloat64,
		"*string": mapPtrInt16ToPtrString,
		"*bool": mapPtrInt16ToPtrBool,
		"*int": mapPtrInt16ToPtrInt,
		"*int8": mapPtrInt16ToPtrInt8,
		"*int16": mapPtrInt16ToPtrInt16,
		"*int32": mapPtrInt16ToPtrInt32,
		"*int64": mapPtrInt16ToPtrInt64,
		"*uint": mapPtrInt16ToPtrUint,
		"*uint8": mapPtrInt16ToPtrUint8,
		"*uint16": mapPtrInt16ToPtrUint16,
		"*uint32": mapPtrInt16ToPtrUint32,
		"*uint64": mapPtrInt16ToPtrUint64,
		"*float32": mapPtrInt16ToPtrFloat32,
		"*float64": mapPtrInt16ToPtrFloat64,
	},
	"int32": {
		"string": mapInt32ToString,
		"bool": mapInt32ToBool,
		"int": mapInt32ToInt,
		"int8": mapInt32ToInt8,
		"int16": mapInt32ToInt16,
		"int32": mapInt32ToInt32,
		"int64": mapInt32ToInt64,
		"uint": mapInt32ToUint,
		"uint8": mapInt32ToUint8,
		"uint16": mapInt32ToUint16,
		"uint32": mapInt32ToUint32,
		"uint64": mapInt32ToUint64,
		"float32": mapInt32ToFloat32,
		"float64": mapInt32ToFloat64,
		"*string": mapInt32ToPtrString,
		"*bool": mapInt32ToPtrBool,
		"*int": mapInt32ToPtrInt,
		"*int8": mapInt32ToPtrInt8,
		"*int16": mapInt32ToPtrInt16,
		"*int32": mapInt32ToPtrInt32,
		"*int64": mapInt32ToPtrInt64,
		"*uint": mapInt32ToPtrUint,
		"*uint8": mapInt32ToPtrUint8,
		"*uint16": mapInt32ToPtrUint16,
		"*uint32": mapInt32ToPtrUint32,
		"*uint64": mapInt32ToPtrUint64,
		"*float32": mapInt32ToPtrFloat32,
		"*float64": mapInt32ToPtrFloat64,
	},
	"*int32": {
		"string": mapPtrInt32ToString,
		"bool": mapPtrInt32ToBool,
		"int": mapPtrInt32ToInt,
		"int8": mapPtrInt32ToInt8,
		"int16": mapPtrInt32ToInt16,
		"int32": mapPtrInt32ToInt32,
		"int64": mapPtrInt32ToInt64,
		"uint": mapPtrInt32ToUint,
		"uint8": mapPtrInt32ToUint8,
		"uint16": mapPtrInt32ToUint16,
		"uint32": mapPtrInt32ToUint32,
		"uint64": mapPtrInt32ToUint64,
		"float32": mapPtrInt32ToFloat32,
		"float64": mapPtrInt32ToFloat64,
		"*string": mapPtrInt32ToPtrString,
		"*bool": mapPtrInt32ToPtrBool,
		"*int": mapPtrInt32ToPtrInt,
		"*int8": mapPtrInt32ToPtrInt8,
		"*int16": mapPtrInt32ToPtrInt16,
		"*int32": mapPtrInt32ToPtrInt32,
		"*int64": mapPtrInt32ToPtrInt64,
		"*uint": mapPtrInt32ToPtrUint,
		"*uint8": mapPtrInt32ToPtrUint8,
		"*uint16": mapPtrInt32ToPtrUint16,
		"*uint32": mapPtrInt32ToPtrUint32,
		"*uint64": mapPtrInt32ToPtrUint64,
		"*float32": mapPtrInt32ToPtrFloat32,
		"*float64": mapPtrInt32ToPtrFloat64,
	},
	"int64": {
		"string": mapInt64ToString,
		"bool": mapInt64ToBool,
		"int": mapInt64ToInt,
		"int8": mapInt64ToInt8,
		"int16": mapInt64ToInt16,
		"int32": mapInt64ToInt32,
		"int64": mapInt64ToInt64,
		"uint": mapInt64ToUint,
		"uint8": mapInt64ToUint8,
		"uint16": mapInt64ToUint16,
		"uint32": mapInt64ToUint32,
		"uint64": mapInt64ToUint64,
		"float32": mapInt64ToFloat32,
		"float64": mapInt64ToFloat64,
		"*string": mapInt64ToPtrString,
		"*bool": mapInt64ToPtrBool,
		"*int": mapInt64ToPtrInt,
		"*int8": mapInt64ToPtrInt8,
		"*int16": mapInt64ToPtrInt16,
		"*int32": mapInt64ToPtrInt32,
		"*int64": mapInt64ToPtrInt64,
		"*uint": mapInt64ToPtrUint,
		"*uint8": mapInt64ToPtrUint8,
		"*uint16": mapInt64ToPtrUint16,
		"*uint32": mapInt64ToPtrUint32,
		"*uint64": mapInt64ToPtrUint64,
		"*float32": mapInt64ToPtrFloat32,
		"*float64": mapInt64ToPtrFloat64,
	},
	"*int64": {
		"string": mapPtrInt64ToString,
		"bool": mapPtrInt64ToBool,
		"int": mapPtrInt64ToInt,
		"int8": mapPtrInt64ToInt8,
		"int16": mapPtrInt64ToInt16,
		"int32": mapPtrInt64ToInt32,
		"int64": mapPtrInt64ToInt64,
		"uint": mapPtrInt64ToUint,
		"uint8": mapPtrInt64ToUint8,
		"uint16": mapPtrInt64ToUint16,
		"uint32": mapPtrInt64ToUint32,
		"uint64": mapPtrInt64ToUint64,
		"float32": mapPtrInt64ToFloat32,
		"float64": mapPtrInt64ToFloat64,
		"*string": mapPtrInt64ToPtrString,
		"*bool": mapPtrInt64ToPtrBool,
		"*int": mapPtrInt64ToPtrInt,
		"*int8": mapPtrInt64ToPtrInt8,
		"*int16": mapPtrInt64ToPtrInt16,
		"*int32": mapPtrInt64ToPtrInt32,
		"*int64": mapPtrInt64ToPtrInt64,
		"*uint": mapPtrInt64ToPtrUint,
		"*uint8": mapPtrInt64ToPtrUint8,
		"*uint16": mapPtrInt64ToPtrUint16,
		"*uint32": mapPtrInt64ToPtrUint32,
		"*uint64": mapPtrInt64ToPtrUint64,
		"*float32": mapPtrInt64ToPtrFloat32,
		"*float64": mapPtrInt64ToPtrFloat64,
	},
	"uint": {
		"string": mapUintToString,
		"bool": mapUintToBool,
		"int": mapUintToInt,
		"int8": mapUintToInt8,
		"int16": mapUintToInt16,
		"int32": mapUintToInt32,
		"int64": mapUintToInt64,
		"uint": mapUintToUint,
		"uint8": mapUintToUint8,
		"uint16": mapUintToUint16,
		"uint32": mapUintToUint32,
		"uint64": mapUintToUint64,
		"float32": mapUintToFloat32,
		"float64": mapUintToFloat64,
		"*string": mapUintToPtrString,
		"*bool": mapUintToPtrBool,
		"*int": mapUintToPtrInt,
		"*int8": mapUintToPtrInt8,
		"*int16": mapUintToPtrInt16,
		"*int32": mapUintToPtrInt32,
		"*int64": mapUintToPtrInt64,
		"*uint": mapUintToPtrUint,
		"*uint8": mapUintToPtrUint8,
		"*uint16": mapUintToPtrUint16,
		"*uint32": mapUintToPtrUint32,
		"*uint64": mapUintToPtrUint64,
		"*float32": mapUintToPtrFloat32,
		"*float64": mapUintToPtrFloat64,
	},
	"*uint": {
		"string": mapPtrUintToString,
		"bool": mapPtrUintToBool,
		"int": mapPtrUintToInt,
		"int8": mapPtrUintToInt8,
		"int16": mapPtrUintToInt16,
		"int32": mapPtrUintToInt32,
		"int64": mapPtrUintToInt64,
		"uint": mapPtrUintToUint,
		"uint8": mapPtrUintToUint8,
		"uint16": mapPtrUintToUint16,
		"uint32": mapPtrUintToUint32,
		"uint64": mapPtrUintToUint64,
		"float32": mapPtrUintToFloat32,
		"float64": mapPtrUintToFloat64,
		"*string": mapPtrUintToPtrString,
		"*bool": mapPtrUintToPtrBool,
		"*int": mapPtrUintToPtrInt,
		"*int8": mapPtrUintToPtrInt8,
		"*int16": mapPtrUintToPtrInt16,
		"*int32": mapPtrUintToPtrInt32,
		"*int64": mapPtrUintToPtrInt64,
		"*uint": mapPtrUintToPtrUint,
		"*uint8": mapPtrUintToPtrUint8,
		"*uint16": mapPtrUintToPtrUint16,
		"*uint32": mapPtrUintToPtrUint32,
		"*uint64": mapPtrUintToPtrUint64,
		"*float32": mapPtrUintToPtrFloat32,
		"*float64": mapPtrUintToPtrFloat64,
	},
	"uint8": {
		"string": mapUint8ToString,
		"bool": mapUint8ToBool,
		"int": mapUint8ToInt,
		"int8": mapUint8ToInt8,
		"int16": mapUint8ToInt16,
		"int32": mapUint8ToInt32,
		"int64": mapUint8ToInt64,
		"uint": mapUint8ToUint,
		"uint8": mapUint8ToUint8,
		"uint16": mapUint8ToUint16,
		"uint32": mapUint8ToUint32,
		"uint64": mapUint8ToUint64,
		"float32": mapUint8ToFloat32,
		"float64": mapUint8ToFloat64,
		"*string": mapUint8ToPtrString,
		"*bool": mapUint8ToPtrBool,
		"*int": mapUint8ToPtrInt,
		"*int8": mapUint8ToPtrInt8,
		"*int16": mapUint8ToPtrInt16,
		"*int32": mapUint8ToPtrInt32,
		"*int64": mapUint8ToPtrInt64,
		"*uint": mapUint8ToPtrUint,
		"*uint8": mapUint8ToPtrUint8,
		"*uint16": mapUint8ToPtrUint16,
		"*uint32": mapUint8ToPtrUint32,
		"*uint64": mapUint8ToPtrUint64,
		"*float32": mapUint8ToPtrFloat32,
		"*float64": mapUint8ToPtrFloat64,
	},
	"*uint8": {
		"string": mapPtrUint8ToString,
		"bool": mapPtrUint8ToBool,
		"int": mapPtrUint8ToInt,
		"int8": mapPtrUint8ToInt8,
		"int16": mapPtrUint8ToInt16,
		"int32": mapPtrUint8ToInt32,
		"int64": mapPtrUint8ToInt64,
		"uint": mapPtrUint8ToUint,
		"uint8": mapPtrUint8ToUint8,
		"uint16": mapPtrUint8ToUint16,
		"uint32": mapPtrUint8ToUint32,
		"uint64": mapPtrUint8ToUint64,
		"float32": mapPtrUint8ToFloat32,
		"float64": mapPtrUint8ToFloat64,
		"*string": mapPtrUint8ToPtrString,
		"*bool": mapPtrUint8ToPtrBool,
		"*int": mapPtrUint8ToPtrInt,
		"*int8": mapPtrUint8ToPtrInt8,
		"*int16": mapPtrUint8ToPtrInt16,
		"*int32": mapPtrUint8ToPtrInt32,
		"*int64": mapPtrUint8ToPtrInt64,
		"*uint": mapPtrUint8ToPtrUint,
		"*uint8": mapPtrUint8ToPtrUint8,
		"*uint16": mapPtrUint8ToPtrUint16,
		"*uint32": mapPtrUint8ToPtrUint32,
		"*uint64": mapPtrUint8ToPtrUint64,
		"*float32": mapPtrUint8ToPtrFloat32,
		"*float64": mapPtrUint8ToPtrFloat64,
	},
	"uint16": {
		"string": mapUint16ToString,
		"bool": mapUint16ToBool,
		"int": mapUint16ToInt,
		"int8": mapUint16ToInt8,
		"int16": mapUint16ToInt16,
		"int32": mapUint16ToInt32,
		"int64": mapUint16ToInt64,
		"uint": mapUint16ToUint,
		"uint8": mapUint16ToUint8,
		"uint16": mapUint16ToUint16,
		"uint32": mapUint16ToUint32,
		"uint64": mapUint16ToUint64,
		"float32": mapUint16ToFloat32,
		"float64": mapUint16ToFloat64,
		"*string": mapUint16ToPtrString,
		"*bool": mapUint16ToPtrBool,
		"*int": mapUint16ToPtrInt,
		"*int8": mapUint16ToPtrInt8,
		"*int16": mapUint16ToPtrInt16,
		"*int32": mapUint16ToPtrInt32,
		"*int64": mapUint16ToPtrInt64,
		"*uint": mapUint16ToPtrUint,
		"*uint8": mapUint16ToPtrUint8,
		"*uint16": mapUint16ToPtrUint16,
		"*uint32": mapUint16ToPtrUint32,
		"*uint64": mapUint16ToPtrUint64,
		"*float32": mapUint16ToPtrFloat32,
		"*float64": mapUint16ToPtrFloat64,
	},
	"*uint16": {
		"string": mapPtrUint16ToString,
		"bool": mapPtrUint16ToBool,
		"int": mapPtrUint16ToInt,
		"int8": mapPtrUint16ToInt8,
		"int16": mapPtrUint16ToInt16,
		"int32": mapPtrUint16ToInt32,
		"int64": mapPtrUint16ToInt64,
		"uint": mapPtrUint16ToUint,
		"uint8": mapPtrUint16ToUint8,
		"uint16": mapPtrUint16ToUint16,
		"uint32": mapPtrUint16ToUint32,
		"uint64": mapPtrUint16ToUint64,
		"float32": mapPtrUint16ToFloat32,
		"float64": mapPtrUint16ToFloat64,
		"*string": mapPtrUint16ToPtrString,
		"*bool": mapPtrUint16ToPtrBool,
		"*int": mapPtrUint16ToPtrInt,
		"*int8": mapPtrUint16ToPtrInt8,
		"*int16": mapPtrUint16ToPtrInt16,
		"*int32": mapPtrUint16ToPtrInt32,
		"*int64": mapPtrUint16ToPtrInt64,
		"*uint": mapPtrUint16ToPtrUint,
		"*uint8": mapPtrUint16ToPtrUint8,
		"*uint16": mapPtrUint16ToPtrUint16,
		"*uint32": mapPtrUint16ToPtrUint32,
		"*uint64": mapPtrUint16ToPtrUint64,
		"*float32": mapPtrUint16ToPtrFloat32,
		"*float64": mapPtrUint16ToPtrFloat64,
	},
	"uint32": {
		"string": mapUint32ToString,
		"bool": mapUint32ToBool,
		"int": mapUint32ToInt,
		"int8": mapUint32ToInt8,
		"int16": mapUint32ToInt16,
		"int32": mapUint32ToInt32,
		"int64": mapUint32ToInt64,
		"uint": mapUint32ToUint,
		"uint8": mapUint32ToUint8,
		"uint16": mapUint32ToUint16,
		"uint32": mapUint32ToUint32,
		"uint64": mapUint32ToUint64,
		"float32": mapUint32ToFloat32,
		"float64": mapUint32ToFloat64,
		"*string": mapUint32ToPtrString,
		"*bool": mapUint32ToPtrBool,
		"*int": mapUint32ToPtrInt,
		"*int8": mapUint32ToPtrInt8,
		"*int16": mapUint32ToPtrInt16,
		"*int32": mapUint32ToPtrInt32,
		"*int64": mapUint32ToPtrInt64,
		"*uint": mapUint32ToPtrUint,
		"*uint8": mapUint32ToPtrUint8,
		"*uint16": mapUint32ToPtrUint16,
		"*uint32": mapUint32ToPtrUint32,
		"*uint64": mapUint32ToPtrUint64,
		"*float32": mapUint32ToPtrFloat32,
		"*float64": mapUint32ToPtrFloat64,
	},
	"*uint32": {
		"string": mapPtrUint32ToString,
		"bool": mapPtrUint32ToBool,
		"int": mapPtrUint32ToInt,
		"int8": mapPtrUint32ToInt8,
		"int16": mapPtrUint32ToInt16,
		"int32": mapPtrUint32ToInt32,
		"int64": mapPtrUint32ToInt64,
		"uint": mapPtrUint32ToUint,
		"uint8": mapPtrUint32ToUint8,
		"uint16": mapPtrUint32ToUint16,
		"uint32": mapPtrUint32ToUint32,
		"uint64": mapPtrUint32ToUint64,
		"float32": mapPtrUint32ToFloat32,
		"float64": mapPtrUint32ToFloat64,
		"*string": mapPtrUint32ToPtrString,
		"*bool": mapPtrUint32ToPtrBool,
		"*int": mapPtrUint32ToPtrInt,
		"*int8": mapPtrUint32ToPtrInt8,
		"*int16": mapPtrUint32ToPtrInt16,
		"*int32": mapPtrUint32ToPtrInt32,
		"*int64": mapPtrUint32ToPtrInt64,
		"*uint": mapPtrUint32ToPtrUint,
		"*uint8": mapPtrUint32ToPtrUint8,
		"*uint16": mapPtrUint32ToPtrUint16,
		"*uint32": mapPtrUint32ToPtrUint32,
		"*uint64": mapPtrUint32ToPtrUint64,
		"*float32": mapPtrUint32ToPtrFloat32,
		"*float64": mapPtrUint32ToPtrFloat64,
	},
	"uint64": {
		"string": mapUint64ToString,
		"bool": mapUint64ToBool,
		"int": mapUint64ToInt,
		"int8": mapUint64ToInt8,
		"int16": mapUint64ToInt16,
		"int32": mapUint64ToInt32,
		"int64": mapUint64ToInt64,
		"uint": mapUint64ToUint,
		"uint8": mapUint64ToUint8,
		"uint16": mapUint64ToUint16,
		"uint32": mapUint64ToUint32,
		"uint64": mapUint64ToUint64,
		"float32": mapUint64ToFloat32,
		"float64": mapUint64ToFloat64,
		"*string": mapUint64ToPtrString,
		"*bool": mapUint64ToPtrBool,
		"*int": mapUint64ToPtrInt,
		"*int8": mapUint64ToPtrInt8,
		"*int16": mapUint64ToPtrInt16,
		"*int32": mapUint64ToPtrInt32,
		"*int64": mapUint64ToPtrInt64,
		"*uint": mapUint64ToPtrUint,
		"*uint8": mapUint64ToPtrUint8,
		"*uint16": mapUint64ToPtrUint16,
		"*uint32": mapUint64ToPtrUint32,
		"*uint64": mapUint64ToPtrUint64,
		"*float32": mapUint64ToPtrFloat32,
		"*float64": mapUint64ToPtrFloat64,
	},
	"*uint64": {
		"string": mapPtrUint64ToString,
		"bool": mapPtrUint64ToBool,
		"int": mapPtrUint64ToInt,
		"int8": mapPtrUint64ToInt8,
		"int16": mapPtrUint64ToInt16,
		"int32": mapPtrUint64ToInt32,
		"int64": mapPtrUint64ToInt64,
		"uint": mapPtrUint64ToUint,
		"uint8": mapPtrUint64ToUint8,
		"uint16": mapPtrUint64ToUint16,
		"uint32": mapPtrUint64ToUint32,
		"uint64": mapPtrUint64ToUint64,
		"float32": mapPtrUint64ToFloat32,
		"float64": mapPtrUint64ToFloat64,
		"*string": mapPtrUint64ToPtrString,
		"*bool": mapPtrUint64ToPtrBool,
		"*int": mapPtrUint64ToPtrInt,
		"*int8": mapPtrUint64ToPtrInt8,
		"*int16": mapPtrUint64ToPtrInt16,
		"*int32": mapPtrUint64ToPtrInt32,
		"*int64": mapPtrUint64ToPtrInt64,
		"*uint": mapPtrUint64ToPtrUint,
		"*uint8": mapPtrUint64ToPtrUint8,
		"*uint16": mapPtrUint64ToPtrUint16,
		"*uint32": mapPtrUint64ToPtrUint32,
		"*uint64": mapPtrUint64ToPtrUint64,
		"*float32": mapPtrUint64ToPtrFloat32,
		"*float64": mapPtrUint64ToPtrFloat64,
	},
	"float32": {
		"string": mapFloat32ToString,
		"bool": mapFloat32ToBool,
		"int": mapFloat32ToInt,
		"int8": mapFloat32ToInt8,
		"int16": mapFloat32ToInt16,
		"int32": mapFloat32ToInt32,
		"int64": mapFloat32ToInt64,
		"uint": mapFloat32ToUint,
		"uint8": mapFloat32ToUint8,
		"uint16": mapFloat32ToUint16,
		"uint32": mapFloat32ToUint32,
		"uint64": mapFloat32ToUint64,
		"float32": mapFloat32ToFloat32,
		"float64": mapFloat32ToFloat64,
		"*string": mapFloat32ToPtrString,
		"*bool": mapFloat32ToPtrBool,
		"*int": mapFloat32ToPtrInt,
		"*int8": mapFloat32ToPtrInt8,
		"*int16": mapFloat32ToPtrInt16,
		"*int32": mapFloat32ToPtrInt32,
		"*int64": mapFloat32ToPtrInt64,
		"*uint": mapFloat32ToPtrUint,
		"*uint8": mapFloat32ToPtrUint8,
		"*uint16": mapFloat32ToPtrUint16,
		"*uint32": mapFloat32ToPtrUint32,
		"*uint64": mapFloat32ToPtrUint64,
		"*float32": mapFloat32ToPtrFloat32,
		"*float64": mapFloat32ToPtrFloat64,
	},
	"*float32": {
		"string": mapPtrFloat32ToString,
		"bool": mapPtrFloat32ToBool,
		"int": mapPtrFloat32ToInt,
		"int8": mapPtrFloat32ToInt8,
		"int16": mapPtrFloat32ToInt16,
		"int32": mapPtrFloat32ToInt32,
		"int64": mapPtrFloat32ToInt64,
		"uint": mapPtrFloat32ToUint,
		"uint8": mapPtrFloat32ToUint8,
		"uint16": mapPtrFloat32ToUint16,
		"uint32": mapPtrFloat32ToUint32,
		"uint64": mapPtrFloat32ToUint64,
		"float32": mapPtrFloat32ToFloat32,
		"float64": mapPtrFloat32ToFloat64,
		"*string": mapPtrFloat32ToPtrString,
		"*bool": mapPtrFloat32ToPtrBool,
		"*int": mapPtrFloat32ToPtrInt,
		"*int8": mapPtrFloat32ToPtrInt8,
		"*int16": mapPtrFloat32ToPtrInt16,
		"*int32": mapPtrFloat32ToPtrInt32,
		"*int64": mapPtrFloat32ToPtrInt64,
		"*uint": mapPtrFloat32ToPtrUint,
		"*uint8": mapPtrFloat32ToPtrUint8,
		"*uint16": mapPtrFloat32ToPtrUint16,
		"*uint32": mapPtrFloat32ToPtrUint32,
		"*uint64": mapPtrFloat32ToPtrUint64,
		"*float32": mapPtrFloat32ToPtrFloat32,
		"*float64": mapPtrFloat32ToPtrFloat64,
	},
	"float64": {
		"string": mapFloat64ToString,
		"bool": mapFloat64ToBool,
		"int": mapFloat64ToInt,
		"int8": mapFloat64ToInt8,
		"int16": mapFloat64ToInt16,
		"int32": mapFloat64ToInt32,
		"int64": mapFloat64ToInt64,
		"uint": mapFloat64ToUint,
		"uint8": mapFloat64ToUint8,
		"uint16": mapFloat64ToUint16,
		"uint32": mapFloat64ToUint32,
		"uint64": mapFloat64ToUint64,
		"float32": mapFloat64ToFloat32,
		"float64": mapFloat64ToFloat64,
		"*string": mapFloat64ToPtrString,
		"*bool": mapFloat64ToPtrBool,
		"*int": mapFloat64ToPtrInt,
		"*int8": mapFloat64ToPtrInt8,
		"*int16": mapFloat64ToPtrInt16,
		"*int32": mapFloat64ToPtrInt32,
		"*int64": mapFloat64ToPtrInt64,
		"*uint": mapFloat64ToPtrUint,
		"*uint8": mapFloat64ToPtrUint8,
		"*uint16": mapFloat64ToPtrUint16,
		"*uint32": mapFloat64ToPtrUint32,
		"*uint64": mapFloat64ToPtrUint64,
		"*float32": mapFloat64ToPtrFloat32,
		"*float64": mapFloat64ToPtrFloat64,
	},
	"*float64": {
		"string": mapPtrFloat64ToString,
		"bool": mapPtrFloat64ToBool,
		"int": mapPtrFloat64ToInt,
		"int8": mapPtrFloat64ToInt8,
		"int16": mapPtrFloat64ToInt16,
		"int32": mapPtrFloat64ToInt32,
		"int64": mapPtrFloat64ToInt64,
		"uint": mapPtrFloat64ToUint,
		"uint8": mapPtrFloat64ToUint8,
		"uint16": mapPtrFloat64ToUint16,
		"uint32": mapPtrFloat64ToUint32,
		"uint64": mapPtrFloat64ToUint64,
		"float32": mapPtrFloat64ToFloat32,
		"float64": mapPtrFloat64ToFloat64,
		"*string": mapPtrFloat64ToPtrString,
		"*bool": mapPtrFloat64ToPtrBool,
		"*int": mapPtrFloat64ToPtrInt,
		"*int8": mapPtrFloat64ToPtrInt8,
		"*int16": mapPtrFloat64ToPtrInt16,
		"*int32": mapPtrFloat64ToPtrInt32,
		"*int64": mapPtrFloat64ToPtrInt64,
		"*uint": mapPtrFloat64ToPtrUint,
		"*uint8": mapPtrFloat64ToPtrUint8,
		"*uint16": mapPtrFloat64ToPtrUint16,
		"*uint32": mapPtrFloat64ToPtrUint32,
		"*uint64": mapPtrFloat64ToPtrUint64,
		"*float32": mapPtrFloat64ToPtrFloat32,
		"*float64": mapPtrFloat64ToPtrFloat64,
	},
}

func dispatch(items *reflect.Value, function interface{}, info *mapInfo) (bool,interface{}) {
	input:=info.fnInputType.String()	
	output:=info.fnOutputType.String()
	if inputVal,ok:=dispatcher[input];ok{
		if outputVal,ok:=inputVal[output];ok{
			return true, outputVal(items,function,info)
		}
    }
	return false, nil
}

func mapStringToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]string)
	fn := function.(func(string) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*string)
	fn := function.(func(string) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]string)
	fn := function.(func(*string) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*string)
	fn := function.(func(*string) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]bool)
	fn := function.(func(string) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*bool)
	fn := function.(func(string) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]bool)
	fn := function.(func(*string) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*string) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]int)
	fn := function.(func(string) int)
	for i := 0; i < len(input); i++ {
		go func(index int){output = append(output,fn(input[index]))}(i)
	}
	return output
}

func mapStringToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*int)
	fn := function.(func(string) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]int)
	fn := function.(func(*string) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*int)
	fn := function.(func(*string) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]int8)
	fn := function.(func(string) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*int8)
	fn := function.(func(string) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]int8)
	fn := function.(func(*string) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*string) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]int16)
	fn := function.(func(string) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*int16)
	fn := function.(func(string) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]int16)
	fn := function.(func(*string) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*string) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]int32)
	fn := function.(func(string) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*int32)
	fn := function.(func(string) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]int32)
	fn := function.(func(*string) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*string) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]int64)
	fn := function.(func(string) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*int64)
	fn := function.(func(string) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]int64)
	fn := function.(func(*string) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*string) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]uint)
	fn := function.(func(string) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*uint)
	fn := function.(func(string) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]uint)
	fn := function.(func(*string) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*string) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]uint8)
	fn := function.(func(string) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(string) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*string) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*string) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]uint16)
	fn := function.(func(string) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(string) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*string) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*string) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]uint32)
	fn := function.(func(string) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(string) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*string) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*string) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]uint64)
	fn := function.(func(string) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(string) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*string) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*string) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]float32)
	fn := function.(func(string) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*float32)
	fn := function.(func(string) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]float32)
	fn := function.(func(*string) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*string) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]float64)
	fn := function.(func(string) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapStringToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]string)
	output := info.output.Interface().([]*float64)
	fn := function.(func(string) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]float64)
	fn := function.(func(*string) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrStringToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*string)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*string) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]string)
	fn := function.(func(bool) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*string)
	fn := function.(func(bool) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]string)
	fn := function.(func(*bool) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*string)
	fn := function.(func(*bool) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]bool)
	fn := function.(func(bool) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*bool)
	fn := function.(func(bool) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]bool)
	fn := function.(func(*bool) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*bool) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]int)
	fn := function.(func(bool) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*int)
	fn := function.(func(bool) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]int)
	fn := function.(func(*bool) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*int)
	fn := function.(func(*bool) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]int8)
	fn := function.(func(bool) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*int8)
	fn := function.(func(bool) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]int8)
	fn := function.(func(*bool) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*bool) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]int16)
	fn := function.(func(bool) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*int16)
	fn := function.(func(bool) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]int16)
	fn := function.(func(*bool) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*bool) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]int32)
	fn := function.(func(bool) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*int32)
	fn := function.(func(bool) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]int32)
	fn := function.(func(*bool) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*bool) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]int64)
	fn := function.(func(bool) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*int64)
	fn := function.(func(bool) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]int64)
	fn := function.(func(*bool) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*bool) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]uint)
	fn := function.(func(bool) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*uint)
	fn := function.(func(bool) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]uint)
	fn := function.(func(*bool) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*bool) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]uint8)
	fn := function.(func(bool) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(bool) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*bool) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*bool) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]uint16)
	fn := function.(func(bool) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(bool) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*bool) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*bool) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]uint32)
	fn := function.(func(bool) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(bool) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*bool) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*bool) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]uint64)
	fn := function.(func(bool) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(bool) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*bool) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*bool) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]float32)
	fn := function.(func(bool) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*float32)
	fn := function.(func(bool) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]float32)
	fn := function.(func(*bool) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*bool) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]float64)
	fn := function.(func(bool) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapBoolToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]bool)
	output := info.output.Interface().([]*float64)
	fn := function.(func(bool) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]float64)
	fn := function.(func(*bool) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrBoolToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*bool)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*bool) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]string)
	fn := function.(func(int) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*string)
	fn := function.(func(int) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]string)
	fn := function.(func(*int) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*string)
	fn := function.(func(*int) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]bool)
	fn := function.(func(int) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*bool)
	fn := function.(func(int) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]bool)
	fn := function.(func(*int) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*int) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]int)
	fn := function.(func(int) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*int)
	fn := function.(func(int) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]int)
	fn := function.(func(*int) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*int)
	fn := function.(func(*int) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]int8)
	fn := function.(func(int) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*int8)
	fn := function.(func(int) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]int8)
	fn := function.(func(*int) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*int) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]int16)
	fn := function.(func(int) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*int16)
	fn := function.(func(int) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]int16)
	fn := function.(func(*int) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*int) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]int32)
	fn := function.(func(int) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*int32)
	fn := function.(func(int) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]int32)
	fn := function.(func(*int) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*int) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]int64)
	fn := function.(func(int) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*int64)
	fn := function.(func(int) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]int64)
	fn := function.(func(*int) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*int) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]uint)
	fn := function.(func(int) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*uint)
	fn := function.(func(int) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]uint)
	fn := function.(func(*int) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*int) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]uint8)
	fn := function.(func(int) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(int) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*int) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*int) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]uint16)
	fn := function.(func(int) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(int) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*int) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*int) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]uint32)
	fn := function.(func(int) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(int) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*int) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*int) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]uint64)
	fn := function.(func(int) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(int) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*int) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*int) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]float32)
	fn := function.(func(int) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*float32)
	fn := function.(func(int) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]float32)
	fn := function.(func(*int) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*int) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]float64)
	fn := function.(func(int) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapIntToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int)
	output := info.output.Interface().([]*float64)
	fn := function.(func(int) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]float64)
	fn := function.(func(*int) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrIntToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*int) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]string)
	fn := function.(func(int8) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*string)
	fn := function.(func(int8) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]string)
	fn := function.(func(*int8) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*string)
	fn := function.(func(*int8) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]bool)
	fn := function.(func(int8) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*bool)
	fn := function.(func(int8) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]bool)
	fn := function.(func(*int8) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*int8) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]int)
	fn := function.(func(int8) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*int)
	fn := function.(func(int8) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]int)
	fn := function.(func(*int8) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*int)
	fn := function.(func(*int8) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]int8)
	fn := function.(func(int8) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*int8)
	fn := function.(func(int8) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]int8)
	fn := function.(func(*int8) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*int8) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]int16)
	fn := function.(func(int8) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*int16)
	fn := function.(func(int8) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]int16)
	fn := function.(func(*int8) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*int8) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]int32)
	fn := function.(func(int8) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*int32)
	fn := function.(func(int8) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]int32)
	fn := function.(func(*int8) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*int8) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]int64)
	fn := function.(func(int8) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*int64)
	fn := function.(func(int8) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]int64)
	fn := function.(func(*int8) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*int8) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]uint)
	fn := function.(func(int8) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*uint)
	fn := function.(func(int8) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]uint)
	fn := function.(func(*int8) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*int8) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]uint8)
	fn := function.(func(int8) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(int8) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*int8) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*int8) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]uint16)
	fn := function.(func(int8) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(int8) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*int8) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*int8) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]uint32)
	fn := function.(func(int8) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(int8) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*int8) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*int8) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]uint64)
	fn := function.(func(int8) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(int8) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*int8) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*int8) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]float32)
	fn := function.(func(int8) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*float32)
	fn := function.(func(int8) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]float32)
	fn := function.(func(*int8) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*int8) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]float64)
	fn := function.(func(int8) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt8ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int8)
	output := info.output.Interface().([]*float64)
	fn := function.(func(int8) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]float64)
	fn := function.(func(*int8) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt8ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int8)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*int8) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]string)
	fn := function.(func(int16) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*string)
	fn := function.(func(int16) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]string)
	fn := function.(func(*int16) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*string)
	fn := function.(func(*int16) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]bool)
	fn := function.(func(int16) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*bool)
	fn := function.(func(int16) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]bool)
	fn := function.(func(*int16) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*int16) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]int)
	fn := function.(func(int16) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*int)
	fn := function.(func(int16) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]int)
	fn := function.(func(*int16) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*int)
	fn := function.(func(*int16) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]int8)
	fn := function.(func(int16) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*int8)
	fn := function.(func(int16) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]int8)
	fn := function.(func(*int16) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*int16) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]int16)
	fn := function.(func(int16) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*int16)
	fn := function.(func(int16) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]int16)
	fn := function.(func(*int16) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*int16) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]int32)
	fn := function.(func(int16) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*int32)
	fn := function.(func(int16) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]int32)
	fn := function.(func(*int16) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*int16) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]int64)
	fn := function.(func(int16) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*int64)
	fn := function.(func(int16) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]int64)
	fn := function.(func(*int16) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*int16) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]uint)
	fn := function.(func(int16) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*uint)
	fn := function.(func(int16) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]uint)
	fn := function.(func(*int16) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*int16) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]uint8)
	fn := function.(func(int16) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(int16) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*int16) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*int16) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]uint16)
	fn := function.(func(int16) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(int16) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*int16) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*int16) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]uint32)
	fn := function.(func(int16) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(int16) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*int16) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*int16) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]uint64)
	fn := function.(func(int16) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(int16) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*int16) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*int16) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]float32)
	fn := function.(func(int16) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*float32)
	fn := function.(func(int16) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]float32)
	fn := function.(func(*int16) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*int16) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]float64)
	fn := function.(func(int16) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt16ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int16)
	output := info.output.Interface().([]*float64)
	fn := function.(func(int16) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]float64)
	fn := function.(func(*int16) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt16ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int16)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*int16) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]string)
	fn := function.(func(int32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*string)
	fn := function.(func(int32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]string)
	fn := function.(func(*int32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*string)
	fn := function.(func(*int32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]bool)
	fn := function.(func(int32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(int32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]bool)
	fn := function.(func(*int32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*int32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]int)
	fn := function.(func(int32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*int)
	fn := function.(func(int32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]int)
	fn := function.(func(*int32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*int)
	fn := function.(func(*int32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]int8)
	fn := function.(func(int32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(int32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]int8)
	fn := function.(func(*int32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*int32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]int16)
	fn := function.(func(int32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(int32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]int16)
	fn := function.(func(*int32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*int32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]int32)
	fn := function.(func(int32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(int32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]int32)
	fn := function.(func(*int32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*int32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]int64)
	fn := function.(func(int32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(int32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]int64)
	fn := function.(func(*int32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*int32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]uint)
	fn := function.(func(int32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(int32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]uint)
	fn := function.(func(*int32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*int32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(int32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(int32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*int32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*int32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(int32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(int32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*int32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*int32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(int32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(int32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*int32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*int32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(int32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(int32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*int32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*int32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]float32)
	fn := function.(func(int32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(int32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]float32)
	fn := function.(func(*int32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*int32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]float64)
	fn := function.(func(int32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(int32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]float64)
	fn := function.(func(*int32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*int32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]string)
	fn := function.(func(int64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*string)
	fn := function.(func(int64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]string)
	fn := function.(func(*int64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*string)
	fn := function.(func(*int64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]bool)
	fn := function.(func(int64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(int64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]bool)
	fn := function.(func(*int64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*int64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]int)
	fn := function.(func(int64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*int)
	fn := function.(func(int64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]int)
	fn := function.(func(*int64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*int)
	fn := function.(func(*int64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]int8)
	fn := function.(func(int64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(int64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]int8)
	fn := function.(func(*int64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*int64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]int16)
	fn := function.(func(int64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(int64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]int16)
	fn := function.(func(*int64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*int64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]int32)
	fn := function.(func(int64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(int64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]int32)
	fn := function.(func(*int64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*int64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]int64)
	fn := function.(func(int64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(int64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]int64)
	fn := function.(func(*int64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*int64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]uint)
	fn := function.(func(int64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(int64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]uint)
	fn := function.(func(*int64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*int64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(int64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(int64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*int64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*int64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(int64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(int64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*int64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*int64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(int64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(int64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*int64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*int64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(int64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(int64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*int64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*int64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]float32)
	fn := function.(func(int64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(int64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]float32)
	fn := function.(func(*int64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*int64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]float64)
	fn := function.(func(int64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapInt64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]int64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(int64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]float64)
	fn := function.(func(*int64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrInt64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*int64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*int64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]string)
	fn := function.(func(uint) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*string)
	fn := function.(func(uint) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]string)
	fn := function.(func(*uint) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*string)
	fn := function.(func(*uint) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]bool)
	fn := function.(func(uint) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*bool)
	fn := function.(func(uint) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]bool)
	fn := function.(func(*uint) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*uint) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]int)
	fn := function.(func(uint) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*int)
	fn := function.(func(uint) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]int)
	fn := function.(func(*uint) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*int)
	fn := function.(func(*uint) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]int8)
	fn := function.(func(uint) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*int8)
	fn := function.(func(uint) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]int8)
	fn := function.(func(*uint) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*uint) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]int16)
	fn := function.(func(uint) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*int16)
	fn := function.(func(uint) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]int16)
	fn := function.(func(*uint) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*uint) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]int32)
	fn := function.(func(uint) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*int32)
	fn := function.(func(uint) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]int32)
	fn := function.(func(*uint) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*uint) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]int64)
	fn := function.(func(uint) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*int64)
	fn := function.(func(uint) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]int64)
	fn := function.(func(*uint) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*uint) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]uint)
	fn := function.(func(uint) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*uint)
	fn := function.(func(uint) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]uint)
	fn := function.(func(*uint) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*uint) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]uint8)
	fn := function.(func(uint) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(uint) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*uint) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*uint) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]uint16)
	fn := function.(func(uint) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(uint) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*uint) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*uint) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]uint32)
	fn := function.(func(uint) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(uint) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*uint) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*uint) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]uint64)
	fn := function.(func(uint) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(uint) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*uint) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*uint) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]float32)
	fn := function.(func(uint) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*float32)
	fn := function.(func(uint) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]float32)
	fn := function.(func(*uint) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*uint) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]float64)
	fn := function.(func(uint) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUintToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint)
	output := info.output.Interface().([]*float64)
	fn := function.(func(uint) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]float64)
	fn := function.(func(*uint) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUintToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*uint) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]string)
	fn := function.(func(uint8) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*string)
	fn := function.(func(uint8) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]string)
	fn := function.(func(*uint8) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*string)
	fn := function.(func(*uint8) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]bool)
	fn := function.(func(uint8) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*bool)
	fn := function.(func(uint8) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]bool)
	fn := function.(func(*uint8) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*uint8) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]int)
	fn := function.(func(uint8) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*int)
	fn := function.(func(uint8) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]int)
	fn := function.(func(*uint8) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*int)
	fn := function.(func(*uint8) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]int8)
	fn := function.(func(uint8) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*int8)
	fn := function.(func(uint8) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]int8)
	fn := function.(func(*uint8) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*uint8) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]int16)
	fn := function.(func(uint8) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*int16)
	fn := function.(func(uint8) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]int16)
	fn := function.(func(*uint8) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*uint8) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]int32)
	fn := function.(func(uint8) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*int32)
	fn := function.(func(uint8) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]int32)
	fn := function.(func(*uint8) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*uint8) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]int64)
	fn := function.(func(uint8) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*int64)
	fn := function.(func(uint8) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]int64)
	fn := function.(func(*uint8) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*uint8) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]uint)
	fn := function.(func(uint8) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*uint)
	fn := function.(func(uint8) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]uint)
	fn := function.(func(*uint8) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*uint8) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]uint8)
	fn := function.(func(uint8) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(uint8) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*uint8) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*uint8) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]uint16)
	fn := function.(func(uint8) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(uint8) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*uint8) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*uint8) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]uint32)
	fn := function.(func(uint8) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(uint8) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*uint8) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*uint8) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]uint64)
	fn := function.(func(uint8) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(uint8) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*uint8) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*uint8) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]float32)
	fn := function.(func(uint8) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*float32)
	fn := function.(func(uint8) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]float32)
	fn := function.(func(*uint8) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*uint8) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]float64)
	fn := function.(func(uint8) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint8ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint8)
	output := info.output.Interface().([]*float64)
	fn := function.(func(uint8) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]float64)
	fn := function.(func(*uint8) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint8ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint8)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*uint8) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]string)
	fn := function.(func(uint16) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*string)
	fn := function.(func(uint16) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]string)
	fn := function.(func(*uint16) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*string)
	fn := function.(func(*uint16) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]bool)
	fn := function.(func(uint16) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*bool)
	fn := function.(func(uint16) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]bool)
	fn := function.(func(*uint16) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*uint16) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]int)
	fn := function.(func(uint16) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*int)
	fn := function.(func(uint16) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]int)
	fn := function.(func(*uint16) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*int)
	fn := function.(func(*uint16) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]int8)
	fn := function.(func(uint16) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*int8)
	fn := function.(func(uint16) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]int8)
	fn := function.(func(*uint16) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*uint16) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]int16)
	fn := function.(func(uint16) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*int16)
	fn := function.(func(uint16) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]int16)
	fn := function.(func(*uint16) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*uint16) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]int32)
	fn := function.(func(uint16) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*int32)
	fn := function.(func(uint16) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]int32)
	fn := function.(func(*uint16) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*uint16) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]int64)
	fn := function.(func(uint16) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*int64)
	fn := function.(func(uint16) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]int64)
	fn := function.(func(*uint16) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*uint16) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]uint)
	fn := function.(func(uint16) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*uint)
	fn := function.(func(uint16) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]uint)
	fn := function.(func(*uint16) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*uint16) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]uint8)
	fn := function.(func(uint16) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(uint16) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*uint16) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*uint16) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]uint16)
	fn := function.(func(uint16) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(uint16) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*uint16) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*uint16) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]uint32)
	fn := function.(func(uint16) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(uint16) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*uint16) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*uint16) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]uint64)
	fn := function.(func(uint16) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(uint16) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*uint16) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*uint16) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]float32)
	fn := function.(func(uint16) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*float32)
	fn := function.(func(uint16) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]float32)
	fn := function.(func(*uint16) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*uint16) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]float64)
	fn := function.(func(uint16) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint16ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint16)
	output := info.output.Interface().([]*float64)
	fn := function.(func(uint16) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]float64)
	fn := function.(func(*uint16) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint16ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint16)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*uint16) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]string)
	fn := function.(func(uint32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*string)
	fn := function.(func(uint32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]string)
	fn := function.(func(*uint32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*string)
	fn := function.(func(*uint32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]bool)
	fn := function.(func(uint32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(uint32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]bool)
	fn := function.(func(*uint32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*uint32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]int)
	fn := function.(func(uint32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*int)
	fn := function.(func(uint32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]int)
	fn := function.(func(*uint32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*int)
	fn := function.(func(*uint32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]int8)
	fn := function.(func(uint32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(uint32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]int8)
	fn := function.(func(*uint32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*uint32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]int16)
	fn := function.(func(uint32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(uint32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]int16)
	fn := function.(func(*uint32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*uint32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]int32)
	fn := function.(func(uint32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(uint32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]int32)
	fn := function.(func(*uint32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*uint32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]int64)
	fn := function.(func(uint32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(uint32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]int64)
	fn := function.(func(*uint32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*uint32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]uint)
	fn := function.(func(uint32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(uint32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]uint)
	fn := function.(func(*uint32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*uint32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(uint32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(uint32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*uint32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*uint32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(uint32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(uint32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*uint32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*uint32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(uint32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(uint32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*uint32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*uint32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(uint32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(uint32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*uint32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*uint32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]float32)
	fn := function.(func(uint32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(uint32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]float32)
	fn := function.(func(*uint32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*uint32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]float64)
	fn := function.(func(uint32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(uint32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]float64)
	fn := function.(func(*uint32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*uint32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]string)
	fn := function.(func(uint64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*string)
	fn := function.(func(uint64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]string)
	fn := function.(func(*uint64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*string)
	fn := function.(func(*uint64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]bool)
	fn := function.(func(uint64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(uint64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]bool)
	fn := function.(func(*uint64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*uint64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]int)
	fn := function.(func(uint64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*int)
	fn := function.(func(uint64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]int)
	fn := function.(func(*uint64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*int)
	fn := function.(func(*uint64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]int8)
	fn := function.(func(uint64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(uint64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]int8)
	fn := function.(func(*uint64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*uint64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]int16)
	fn := function.(func(uint64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(uint64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]int16)
	fn := function.(func(*uint64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*uint64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]int32)
	fn := function.(func(uint64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(uint64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]int32)
	fn := function.(func(*uint64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*uint64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]int64)
	fn := function.(func(uint64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(uint64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]int64)
	fn := function.(func(*uint64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*uint64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]uint)
	fn := function.(func(uint64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(uint64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]uint)
	fn := function.(func(*uint64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*uint64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(uint64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(uint64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*uint64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*uint64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(uint64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(uint64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*uint64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*uint64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(uint64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(uint64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*uint64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*uint64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(uint64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(uint64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*uint64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*uint64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]float32)
	fn := function.(func(uint64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(uint64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]float32)
	fn := function.(func(*uint64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*uint64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]float64)
	fn := function.(func(uint64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapUint64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]uint64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(uint64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]float64)
	fn := function.(func(*uint64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrUint64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*uint64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*uint64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]string)
	fn := function.(func(float32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*string)
	fn := function.(func(float32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]string)
	fn := function.(func(*float32) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*string)
	fn := function.(func(*float32) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]bool)
	fn := function.(func(float32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(float32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]bool)
	fn := function.(func(*float32) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*float32) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]int)
	fn := function.(func(float32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*int)
	fn := function.(func(float32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]int)
	fn := function.(func(*float32) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*int)
	fn := function.(func(*float32) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]int8)
	fn := function.(func(float32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(float32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]int8)
	fn := function.(func(*float32) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*float32) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]int16)
	fn := function.(func(float32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(float32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]int16)
	fn := function.(func(*float32) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*float32) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]int32)
	fn := function.(func(float32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(float32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]int32)
	fn := function.(func(*float32) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*float32) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]int64)
	fn := function.(func(float32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(float32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]int64)
	fn := function.(func(*float32) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*float32) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]uint)
	fn := function.(func(float32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(float32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]uint)
	fn := function.(func(*float32) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*float32) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(float32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(float32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*float32) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*float32) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(float32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(float32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*float32) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*float32) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(float32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(float32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*float32) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*float32) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(float32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(float32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*float32) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*float32) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]float32)
	fn := function.(func(float32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(float32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]float32)
	fn := function.(func(*float32) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*float32) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]float64)
	fn := function.(func(float32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(float32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]float64)
	fn := function.(func(*float32) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat32ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float32)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*float32) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]string)
	fn := function.(func(float64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*string)
	fn := function.(func(float64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]string)
	fn := function.(func(*float64) string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrString(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*string)
	fn := function.(func(*float64) *string)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]bool)
	fn := function.(func(float64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(float64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]bool)
	fn := function.(func(*float64) bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrBool(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*bool)
	fn := function.(func(*float64) *bool)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]int)
	fn := function.(func(float64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*int)
	fn := function.(func(float64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]int)
	fn := function.(func(*float64) int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrInt(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*int)
	fn := function.(func(*float64) *int)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]int8)
	fn := function.(func(float64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(float64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]int8)
	fn := function.(func(*float64) int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrInt8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*int8)
	fn := function.(func(*float64) *int8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]int16)
	fn := function.(func(float64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(float64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]int16)
	fn := function.(func(*float64) int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrInt16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*int16)
	fn := function.(func(*float64) *int16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]int32)
	fn := function.(func(float64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(float64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]int32)
	fn := function.(func(*float64) int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrInt32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*int32)
	fn := function.(func(*float64) *int32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]int64)
	fn := function.(func(float64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(float64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]int64)
	fn := function.(func(*float64) int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrInt64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*int64)
	fn := function.(func(*float64) *int64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]uint)
	fn := function.(func(float64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(float64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]uint)
	fn := function.(func(*float64) uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrUint(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*uint)
	fn := function.(func(*float64) *uint)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(float64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(float64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]uint8)
	fn := function.(func(*float64) uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrUint8(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*uint8)
	fn := function.(func(*float64) *uint8)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(float64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(float64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]uint16)
	fn := function.(func(*float64) uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrUint16(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*uint16)
	fn := function.(func(*float64) *uint16)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(float64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(float64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]uint32)
	fn := function.(func(*float64) uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrUint32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*uint32)
	fn := function.(func(*float64) *uint32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(float64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(float64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]uint64)
	fn := function.(func(*float64) uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrUint64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*uint64)
	fn := function.(func(*float64) *uint64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]float32)
	fn := function.(func(float64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(float64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]float32)
	fn := function.(func(*float64) float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrFloat32(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*float32)
	fn := function.(func(*float64) *float32)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]float64)
	fn := function.(func(float64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}

func mapFloat64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]float64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(float64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]float64)
	fn := function.(func(*float64) float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}


func mapPtrFloat64ToPtrFloat64(itemsValue *reflect.Value, function interface{}, info *mapInfo) interface{}  {
	input := itemsValue.Interface().([]*float64)
	output := info.output.Interface().([]*float64)
	fn := function.(func(*float64) *float64)
	for i := 0; i < len(input); i++ {
		output = append(output,fn(input[i]))
	}
	return output
}
