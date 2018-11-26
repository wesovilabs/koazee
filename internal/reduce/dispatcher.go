package reduce

import "reflect"

type dispatchFunction func(items reflect.Value, fn interface{}) interface{}

var dispatcher = map[string]map[string]dispatchFunction{
	"string": {
		"string":   reduceStringToString,
		"bool":     reduceStringToBool,
		"int":      reduceStringToInt,
		"int8":     reduceStringToInt8,
		"int16":    reduceStringToInt16,
		"int32":    reduceStringToInt32,
		"int64":    reduceStringToInt64,
		"uint":     reduceStringToUint,
		"uint8":    reduceStringToUint8,
		"uint16":   reduceStringToUint16,
		"uint32":   reduceStringToUint32,
		"uint64":   reduceStringToUint64,
		"float32":  reduceStringToFloat32,
		"float64":  reduceStringToFloat64,
		"*string":  reduceStringToPtrString,
		"*bool":    reduceStringToPtrBool,
		"*int":     reduceStringToPtrInt,
		"*int8":    reduceStringToPtrInt8,
		"*int16":   reduceStringToPtrInt16,
		"*int32":   reduceStringToPtrInt32,
		"*int64":   reduceStringToPtrInt64,
		"*uint":    reduceStringToPtrUint,
		"*uint8":   reduceStringToPtrUint8,
		"*uint16":  reduceStringToPtrUint16,
		"*uint32":  reduceStringToPtrUint32,
		"*uint64":  reduceStringToPtrUint64,
		"*float32": reduceStringToPtrFloat32,
		"*float64": reduceStringToPtrFloat64,
	},
	"*string": {
		"string":   reducePtrStringToString,
		"bool":     reducePtrStringToBool,
		"int":      reducePtrStringToInt,
		"int8":     reducePtrStringToInt8,
		"int16":    reducePtrStringToInt16,
		"int32":    reducePtrStringToInt32,
		"int64":    reducePtrStringToInt64,
		"uint":     reducePtrStringToUint,
		"uint8":    reducePtrStringToUint8,
		"uint16":   reducePtrStringToUint16,
		"uint32":   reducePtrStringToUint32,
		"uint64":   reducePtrStringToUint64,
		"float32":  reducePtrStringToFloat32,
		"float64":  reducePtrStringToFloat64,
		"*string":  reducePtrStringToPtrString,
		"*bool":    reducePtrStringToPtrBool,
		"*int":     reducePtrStringToPtrInt,
		"*int8":    reducePtrStringToPtrInt8,
		"*int16":   reducePtrStringToPtrInt16,
		"*int32":   reducePtrStringToPtrInt32,
		"*int64":   reducePtrStringToPtrInt64,
		"*uint":    reducePtrStringToPtrUint,
		"*uint8":   reducePtrStringToPtrUint8,
		"*uint16":  reducePtrStringToPtrUint16,
		"*uint32":  reducePtrStringToPtrUint32,
		"*uint64":  reducePtrStringToPtrUint64,
		"*float32": reducePtrStringToPtrFloat32,
		"*float64": reducePtrStringToPtrFloat64,
	},
	"bool": {
		"string":   reduceBoolToString,
		"bool":     reduceBoolToBool,
		"int":      reduceBoolToInt,
		"int8":     reduceBoolToInt8,
		"int16":    reduceBoolToInt16,
		"int32":    reduceBoolToInt32,
		"int64":    reduceBoolToInt64,
		"uint":     reduceBoolToUint,
		"uint8":    reduceBoolToUint8,
		"uint16":   reduceBoolToUint16,
		"uint32":   reduceBoolToUint32,
		"uint64":   reduceBoolToUint64,
		"float32":  reduceBoolToFloat32,
		"float64":  reduceBoolToFloat64,
		"*string":  reduceBoolToPtrString,
		"*bool":    reduceBoolToPtrBool,
		"*int":     reduceBoolToPtrInt,
		"*int8":    reduceBoolToPtrInt8,
		"*int16":   reduceBoolToPtrInt16,
		"*int32":   reduceBoolToPtrInt32,
		"*int64":   reduceBoolToPtrInt64,
		"*uint":    reduceBoolToPtrUint,
		"*uint8":   reduceBoolToPtrUint8,
		"*uint16":  reduceBoolToPtrUint16,
		"*uint32":  reduceBoolToPtrUint32,
		"*uint64":  reduceBoolToPtrUint64,
		"*float32": reduceBoolToPtrFloat32,
		"*float64": reduceBoolToPtrFloat64,
	},
	"*bool": {
		"string":   reducePtrBoolToString,
		"bool":     reducePtrBoolToBool,
		"int":      reducePtrBoolToInt,
		"int8":     reducePtrBoolToInt8,
		"int16":    reducePtrBoolToInt16,
		"int32":    reducePtrBoolToInt32,
		"int64":    reducePtrBoolToInt64,
		"uint":     reducePtrBoolToUint,
		"uint8":    reducePtrBoolToUint8,
		"uint16":   reducePtrBoolToUint16,
		"uint32":   reducePtrBoolToUint32,
		"uint64":   reducePtrBoolToUint64,
		"float32":  reducePtrBoolToFloat32,
		"float64":  reducePtrBoolToFloat64,
		"*string":  reducePtrBoolToPtrString,
		"*bool":    reducePtrBoolToPtrBool,
		"*int":     reducePtrBoolToPtrInt,
		"*int8":    reducePtrBoolToPtrInt8,
		"*int16":   reducePtrBoolToPtrInt16,
		"*int32":   reducePtrBoolToPtrInt32,
		"*int64":   reducePtrBoolToPtrInt64,
		"*uint":    reducePtrBoolToPtrUint,
		"*uint8":   reducePtrBoolToPtrUint8,
		"*uint16":  reducePtrBoolToPtrUint16,
		"*uint32":  reducePtrBoolToPtrUint32,
		"*uint64":  reducePtrBoolToPtrUint64,
		"*float32": reducePtrBoolToPtrFloat32,
		"*float64": reducePtrBoolToPtrFloat64,
	},
	"int": {
		"string":   reduceIntToString,
		"bool":     reduceIntToBool,
		"int":      reduceIntToInt,
		"int8":     reduceIntToInt8,
		"int16":    reduceIntToInt16,
		"int32":    reduceIntToInt32,
		"int64":    reduceIntToInt64,
		"uint":     reduceIntToUint,
		"uint8":    reduceIntToUint8,
		"uint16":   reduceIntToUint16,
		"uint32":   reduceIntToUint32,
		"uint64":   reduceIntToUint64,
		"float32":  reduceIntToFloat32,
		"float64":  reduceIntToFloat64,
		"*string":  reduceIntToPtrString,
		"*bool":    reduceIntToPtrBool,
		"*int":     reduceIntToPtrInt,
		"*int8":    reduceIntToPtrInt8,
		"*int16":   reduceIntToPtrInt16,
		"*int32":   reduceIntToPtrInt32,
		"*int64":   reduceIntToPtrInt64,
		"*uint":    reduceIntToPtrUint,
		"*uint8":   reduceIntToPtrUint8,
		"*uint16":  reduceIntToPtrUint16,
		"*uint32":  reduceIntToPtrUint32,
		"*uint64":  reduceIntToPtrUint64,
		"*float32": reduceIntToPtrFloat32,
		"*float64": reduceIntToPtrFloat64,
	},
	"*int": {
		"string":   reducePtrIntToString,
		"bool":     reducePtrIntToBool,
		"int":      reducePtrIntToInt,
		"int8":     reducePtrIntToInt8,
		"int16":    reducePtrIntToInt16,
		"int32":    reducePtrIntToInt32,
		"int64":    reducePtrIntToInt64,
		"uint":     reducePtrIntToUint,
		"uint8":    reducePtrIntToUint8,
		"uint16":   reducePtrIntToUint16,
		"uint32":   reducePtrIntToUint32,
		"uint64":   reducePtrIntToUint64,
		"float32":  reducePtrIntToFloat32,
		"float64":  reducePtrIntToFloat64,
		"*string":  reducePtrIntToPtrString,
		"*bool":    reducePtrIntToPtrBool,
		"*int":     reducePtrIntToPtrInt,
		"*int8":    reducePtrIntToPtrInt8,
		"*int16":   reducePtrIntToPtrInt16,
		"*int32":   reducePtrIntToPtrInt32,
		"*int64":   reducePtrIntToPtrInt64,
		"*uint":    reducePtrIntToPtrUint,
		"*uint8":   reducePtrIntToPtrUint8,
		"*uint16":  reducePtrIntToPtrUint16,
		"*uint32":  reducePtrIntToPtrUint32,
		"*uint64":  reducePtrIntToPtrUint64,
		"*float32": reducePtrIntToPtrFloat32,
		"*float64": reducePtrIntToPtrFloat64,
	},
	"int8": {
		"string":   reduceInt8ToString,
		"bool":     reduceInt8ToBool,
		"int":      reduceInt8ToInt,
		"int8":     reduceInt8ToInt8,
		"int16":    reduceInt8ToInt16,
		"int32":    reduceInt8ToInt32,
		"int64":    reduceInt8ToInt64,
		"uint":     reduceInt8ToUint,
		"uint8":    reduceInt8ToUint8,
		"uint16":   reduceInt8ToUint16,
		"uint32":   reduceInt8ToUint32,
		"uint64":   reduceInt8ToUint64,
		"float32":  reduceInt8ToFloat32,
		"float64":  reduceInt8ToFloat64,
		"*string":  reduceInt8ToPtrString,
		"*bool":    reduceInt8ToPtrBool,
		"*int":     reduceInt8ToPtrInt,
		"*int8":    reduceInt8ToPtrInt8,
		"*int16":   reduceInt8ToPtrInt16,
		"*int32":   reduceInt8ToPtrInt32,
		"*int64":   reduceInt8ToPtrInt64,
		"*uint":    reduceInt8ToPtrUint,
		"*uint8":   reduceInt8ToPtrUint8,
		"*uint16":  reduceInt8ToPtrUint16,
		"*uint32":  reduceInt8ToPtrUint32,
		"*uint64":  reduceInt8ToPtrUint64,
		"*float32": reduceInt8ToPtrFloat32,
		"*float64": reduceInt8ToPtrFloat64,
	},
	"*int8": {
		"string":   reducePtrInt8ToString,
		"bool":     reducePtrInt8ToBool,
		"int":      reducePtrInt8ToInt,
		"int8":     reducePtrInt8ToInt8,
		"int16":    reducePtrInt8ToInt16,
		"int32":    reducePtrInt8ToInt32,
		"int64":    reducePtrInt8ToInt64,
		"uint":     reducePtrInt8ToUint,
		"uint8":    reducePtrInt8ToUint8,
		"uint16":   reducePtrInt8ToUint16,
		"uint32":   reducePtrInt8ToUint32,
		"uint64":   reducePtrInt8ToUint64,
		"float32":  reducePtrInt8ToFloat32,
		"float64":  reducePtrInt8ToFloat64,
		"*string":  reducePtrInt8ToPtrString,
		"*bool":    reducePtrInt8ToPtrBool,
		"*int":     reducePtrInt8ToPtrInt,
		"*int8":    reducePtrInt8ToPtrInt8,
		"*int16":   reducePtrInt8ToPtrInt16,
		"*int32":   reducePtrInt8ToPtrInt32,
		"*int64":   reducePtrInt8ToPtrInt64,
		"*uint":    reducePtrInt8ToPtrUint,
		"*uint8":   reducePtrInt8ToPtrUint8,
		"*uint16":  reducePtrInt8ToPtrUint16,
		"*uint32":  reducePtrInt8ToPtrUint32,
		"*uint64":  reducePtrInt8ToPtrUint64,
		"*float32": reducePtrInt8ToPtrFloat32,
		"*float64": reducePtrInt8ToPtrFloat64,
	},
	"int16": {
		"string":   reduceInt16ToString,
		"bool":     reduceInt16ToBool,
		"int":      reduceInt16ToInt,
		"int8":     reduceInt16ToInt8,
		"int16":    reduceInt16ToInt16,
		"int32":    reduceInt16ToInt32,
		"int64":    reduceInt16ToInt64,
		"uint":     reduceInt16ToUint,
		"uint8":    reduceInt16ToUint8,
		"uint16":   reduceInt16ToUint16,
		"uint32":   reduceInt16ToUint32,
		"uint64":   reduceInt16ToUint64,
		"float32":  reduceInt16ToFloat32,
		"float64":  reduceInt16ToFloat64,
		"*string":  reduceInt16ToPtrString,
		"*bool":    reduceInt16ToPtrBool,
		"*int":     reduceInt16ToPtrInt,
		"*int8":    reduceInt16ToPtrInt8,
		"*int16":   reduceInt16ToPtrInt16,
		"*int32":   reduceInt16ToPtrInt32,
		"*int64":   reduceInt16ToPtrInt64,
		"*uint":    reduceInt16ToPtrUint,
		"*uint8":   reduceInt16ToPtrUint8,
		"*uint16":  reduceInt16ToPtrUint16,
		"*uint32":  reduceInt16ToPtrUint32,
		"*uint64":  reduceInt16ToPtrUint64,
		"*float32": reduceInt16ToPtrFloat32,
		"*float64": reduceInt16ToPtrFloat64,
	},
	"*int16": {
		"string":   reducePtrInt16ToString,
		"bool":     reducePtrInt16ToBool,
		"int":      reducePtrInt16ToInt,
		"int8":     reducePtrInt16ToInt8,
		"int16":    reducePtrInt16ToInt16,
		"int32":    reducePtrInt16ToInt32,
		"int64":    reducePtrInt16ToInt64,
		"uint":     reducePtrInt16ToUint,
		"uint8":    reducePtrInt16ToUint8,
		"uint16":   reducePtrInt16ToUint16,
		"uint32":   reducePtrInt16ToUint32,
		"uint64":   reducePtrInt16ToUint64,
		"float32":  reducePtrInt16ToFloat32,
		"float64":  reducePtrInt16ToFloat64,
		"*string":  reducePtrInt16ToPtrString,
		"*bool":    reducePtrInt16ToPtrBool,
		"*int":     reducePtrInt16ToPtrInt,
		"*int8":    reducePtrInt16ToPtrInt8,
		"*int16":   reducePtrInt16ToPtrInt16,
		"*int32":   reducePtrInt16ToPtrInt32,
		"*int64":   reducePtrInt16ToPtrInt64,
		"*uint":    reducePtrInt16ToPtrUint,
		"*uint8":   reducePtrInt16ToPtrUint8,
		"*uint16":  reducePtrInt16ToPtrUint16,
		"*uint32":  reducePtrInt16ToPtrUint32,
		"*uint64":  reducePtrInt16ToPtrUint64,
		"*float32": reducePtrInt16ToPtrFloat32,
		"*float64": reducePtrInt16ToPtrFloat64,
	},
	"int32": {
		"string":   reduceInt32ToString,
		"bool":     reduceInt32ToBool,
		"int":      reduceInt32ToInt,
		"int8":     reduceInt32ToInt8,
		"int16":    reduceInt32ToInt16,
		"int32":    reduceInt32ToInt32,
		"int64":    reduceInt32ToInt64,
		"uint":     reduceInt32ToUint,
		"uint8":    reduceInt32ToUint8,
		"uint16":   reduceInt32ToUint16,
		"uint32":   reduceInt32ToUint32,
		"uint64":   reduceInt32ToUint64,
		"float32":  reduceInt32ToFloat32,
		"float64":  reduceInt32ToFloat64,
		"*string":  reduceInt32ToPtrString,
		"*bool":    reduceInt32ToPtrBool,
		"*int":     reduceInt32ToPtrInt,
		"*int8":    reduceInt32ToPtrInt8,
		"*int16":   reduceInt32ToPtrInt16,
		"*int32":   reduceInt32ToPtrInt32,
		"*int64":   reduceInt32ToPtrInt64,
		"*uint":    reduceInt32ToPtrUint,
		"*uint8":   reduceInt32ToPtrUint8,
		"*uint16":  reduceInt32ToPtrUint16,
		"*uint32":  reduceInt32ToPtrUint32,
		"*uint64":  reduceInt32ToPtrUint64,
		"*float32": reduceInt32ToPtrFloat32,
		"*float64": reduceInt32ToPtrFloat64,
	},
	"*int32": {
		"string":   reducePtrInt32ToString,
		"bool":     reducePtrInt32ToBool,
		"int":      reducePtrInt32ToInt,
		"int8":     reducePtrInt32ToInt8,
		"int16":    reducePtrInt32ToInt16,
		"int32":    reducePtrInt32ToInt32,
		"int64":    reducePtrInt32ToInt64,
		"uint":     reducePtrInt32ToUint,
		"uint8":    reducePtrInt32ToUint8,
		"uint16":   reducePtrInt32ToUint16,
		"uint32":   reducePtrInt32ToUint32,
		"uint64":   reducePtrInt32ToUint64,
		"float32":  reducePtrInt32ToFloat32,
		"float64":  reducePtrInt32ToFloat64,
		"*string":  reducePtrInt32ToPtrString,
		"*bool":    reducePtrInt32ToPtrBool,
		"*int":     reducePtrInt32ToPtrInt,
		"*int8":    reducePtrInt32ToPtrInt8,
		"*int16":   reducePtrInt32ToPtrInt16,
		"*int32":   reducePtrInt32ToPtrInt32,
		"*int64":   reducePtrInt32ToPtrInt64,
		"*uint":    reducePtrInt32ToPtrUint,
		"*uint8":   reducePtrInt32ToPtrUint8,
		"*uint16":  reducePtrInt32ToPtrUint16,
		"*uint32":  reducePtrInt32ToPtrUint32,
		"*uint64":  reducePtrInt32ToPtrUint64,
		"*float32": reducePtrInt32ToPtrFloat32,
		"*float64": reducePtrInt32ToPtrFloat64,
	},
	"int64": {
		"string":   reduceInt64ToString,
		"bool":     reduceInt64ToBool,
		"int":      reduceInt64ToInt,
		"int8":     reduceInt64ToInt8,
		"int16":    reduceInt64ToInt16,
		"int32":    reduceInt64ToInt32,
		"int64":    reduceInt64ToInt64,
		"uint":     reduceInt64ToUint,
		"uint8":    reduceInt64ToUint8,
		"uint16":   reduceInt64ToUint16,
		"uint32":   reduceInt64ToUint32,
		"uint64":   reduceInt64ToUint64,
		"float32":  reduceInt64ToFloat32,
		"float64":  reduceInt64ToFloat64,
		"*string":  reduceInt64ToPtrString,
		"*bool":    reduceInt64ToPtrBool,
		"*int":     reduceInt64ToPtrInt,
		"*int8":    reduceInt64ToPtrInt8,
		"*int16":   reduceInt64ToPtrInt16,
		"*int32":   reduceInt64ToPtrInt32,
		"*int64":   reduceInt64ToPtrInt64,
		"*uint":    reduceInt64ToPtrUint,
		"*uint8":   reduceInt64ToPtrUint8,
		"*uint16":  reduceInt64ToPtrUint16,
		"*uint32":  reduceInt64ToPtrUint32,
		"*uint64":  reduceInt64ToPtrUint64,
		"*float32": reduceInt64ToPtrFloat32,
		"*float64": reduceInt64ToPtrFloat64,
	},
	"*int64": {
		"string":   reducePtrInt64ToString,
		"bool":     reducePtrInt64ToBool,
		"int":      reducePtrInt64ToInt,
		"int8":     reducePtrInt64ToInt8,
		"int16":    reducePtrInt64ToInt16,
		"int32":    reducePtrInt64ToInt32,
		"int64":    reducePtrInt64ToInt64,
		"uint":     reducePtrInt64ToUint,
		"uint8":    reducePtrInt64ToUint8,
		"uint16":   reducePtrInt64ToUint16,
		"uint32":   reducePtrInt64ToUint32,
		"uint64":   reducePtrInt64ToUint64,
		"float32":  reducePtrInt64ToFloat32,
		"float64":  reducePtrInt64ToFloat64,
		"*string":  reducePtrInt64ToPtrString,
		"*bool":    reducePtrInt64ToPtrBool,
		"*int":     reducePtrInt64ToPtrInt,
		"*int8":    reducePtrInt64ToPtrInt8,
		"*int16":   reducePtrInt64ToPtrInt16,
		"*int32":   reducePtrInt64ToPtrInt32,
		"*int64":   reducePtrInt64ToPtrInt64,
		"*uint":    reducePtrInt64ToPtrUint,
		"*uint8":   reducePtrInt64ToPtrUint8,
		"*uint16":  reducePtrInt64ToPtrUint16,
		"*uint32":  reducePtrInt64ToPtrUint32,
		"*uint64":  reducePtrInt64ToPtrUint64,
		"*float32": reducePtrInt64ToPtrFloat32,
		"*float64": reducePtrInt64ToPtrFloat64,
	},
	"uint": {
		"string":   reduceUintToString,
		"bool":     reduceUintToBool,
		"int":      reduceUintToInt,
		"int8":     reduceUintToInt8,
		"int16":    reduceUintToInt16,
		"int32":    reduceUintToInt32,
		"int64":    reduceUintToInt64,
		"uint":     reduceUintToUint,
		"uint8":    reduceUintToUint8,
		"uint16":   reduceUintToUint16,
		"uint32":   reduceUintToUint32,
		"uint64":   reduceUintToUint64,
		"float32":  reduceUintToFloat32,
		"float64":  reduceUintToFloat64,
		"*string":  reduceUintToPtrString,
		"*bool":    reduceUintToPtrBool,
		"*int":     reduceUintToPtrInt,
		"*int8":    reduceUintToPtrInt8,
		"*int16":   reduceUintToPtrInt16,
		"*int32":   reduceUintToPtrInt32,
		"*int64":   reduceUintToPtrInt64,
		"*uint":    reduceUintToPtrUint,
		"*uint8":   reduceUintToPtrUint8,
		"*uint16":  reduceUintToPtrUint16,
		"*uint32":  reduceUintToPtrUint32,
		"*uint64":  reduceUintToPtrUint64,
		"*float32": reduceUintToPtrFloat32,
		"*float64": reduceUintToPtrFloat64,
	},
	"*uint": {
		"string":   reducePtrUintToString,
		"bool":     reducePtrUintToBool,
		"int":      reducePtrUintToInt,
		"int8":     reducePtrUintToInt8,
		"int16":    reducePtrUintToInt16,
		"int32":    reducePtrUintToInt32,
		"int64":    reducePtrUintToInt64,
		"uint":     reducePtrUintToUint,
		"uint8":    reducePtrUintToUint8,
		"uint16":   reducePtrUintToUint16,
		"uint32":   reducePtrUintToUint32,
		"uint64":   reducePtrUintToUint64,
		"float32":  reducePtrUintToFloat32,
		"float64":  reducePtrUintToFloat64,
		"*string":  reducePtrUintToPtrString,
		"*bool":    reducePtrUintToPtrBool,
		"*int":     reducePtrUintToPtrInt,
		"*int8":    reducePtrUintToPtrInt8,
		"*int16":   reducePtrUintToPtrInt16,
		"*int32":   reducePtrUintToPtrInt32,
		"*int64":   reducePtrUintToPtrInt64,
		"*uint":    reducePtrUintToPtrUint,
		"*uint8":   reducePtrUintToPtrUint8,
		"*uint16":  reducePtrUintToPtrUint16,
		"*uint32":  reducePtrUintToPtrUint32,
		"*uint64":  reducePtrUintToPtrUint64,
		"*float32": reducePtrUintToPtrFloat32,
		"*float64": reducePtrUintToPtrFloat64,
	},
	"uint8": {
		"string":   reduceUint8ToString,
		"bool":     reduceUint8ToBool,
		"int":      reduceUint8ToInt,
		"int8":     reduceUint8ToInt8,
		"int16":    reduceUint8ToInt16,
		"int32":    reduceUint8ToInt32,
		"int64":    reduceUint8ToInt64,
		"uint":     reduceUint8ToUint,
		"uint8":    reduceUint8ToUint8,
		"uint16":   reduceUint8ToUint16,
		"uint32":   reduceUint8ToUint32,
		"uint64":   reduceUint8ToUint64,
		"float32":  reduceUint8ToFloat32,
		"float64":  reduceUint8ToFloat64,
		"*string":  reduceUint8ToPtrString,
		"*bool":    reduceUint8ToPtrBool,
		"*int":     reduceUint8ToPtrInt,
		"*int8":    reduceUint8ToPtrInt8,
		"*int16":   reduceUint8ToPtrInt16,
		"*int32":   reduceUint8ToPtrInt32,
		"*int64":   reduceUint8ToPtrInt64,
		"*uint":    reduceUint8ToPtrUint,
		"*uint8":   reduceUint8ToPtrUint8,
		"*uint16":  reduceUint8ToPtrUint16,
		"*uint32":  reduceUint8ToPtrUint32,
		"*uint64":  reduceUint8ToPtrUint64,
		"*float32": reduceUint8ToPtrFloat32,
		"*float64": reduceUint8ToPtrFloat64,
	},
	"*uint8": {
		"string":   reducePtrUint8ToString,
		"bool":     reducePtrUint8ToBool,
		"int":      reducePtrUint8ToInt,
		"int8":     reducePtrUint8ToInt8,
		"int16":    reducePtrUint8ToInt16,
		"int32":    reducePtrUint8ToInt32,
		"int64":    reducePtrUint8ToInt64,
		"uint":     reducePtrUint8ToUint,
		"uint8":    reducePtrUint8ToUint8,
		"uint16":   reducePtrUint8ToUint16,
		"uint32":   reducePtrUint8ToUint32,
		"uint64":   reducePtrUint8ToUint64,
		"float32":  reducePtrUint8ToFloat32,
		"float64":  reducePtrUint8ToFloat64,
		"*string":  reducePtrUint8ToPtrString,
		"*bool":    reducePtrUint8ToPtrBool,
		"*int":     reducePtrUint8ToPtrInt,
		"*int8":    reducePtrUint8ToPtrInt8,
		"*int16":   reducePtrUint8ToPtrInt16,
		"*int32":   reducePtrUint8ToPtrInt32,
		"*int64":   reducePtrUint8ToPtrInt64,
		"*uint":    reducePtrUint8ToPtrUint,
		"*uint8":   reducePtrUint8ToPtrUint8,
		"*uint16":  reducePtrUint8ToPtrUint16,
		"*uint32":  reducePtrUint8ToPtrUint32,
		"*uint64":  reducePtrUint8ToPtrUint64,
		"*float32": reducePtrUint8ToPtrFloat32,
		"*float64": reducePtrUint8ToPtrFloat64,
	},
	"uint16": {
		"string":   reduceUint16ToString,
		"bool":     reduceUint16ToBool,
		"int":      reduceUint16ToInt,
		"int8":     reduceUint16ToInt8,
		"int16":    reduceUint16ToInt16,
		"int32":    reduceUint16ToInt32,
		"int64":    reduceUint16ToInt64,
		"uint":     reduceUint16ToUint,
		"uint8":    reduceUint16ToUint8,
		"uint16":   reduceUint16ToUint16,
		"uint32":   reduceUint16ToUint32,
		"uint64":   reduceUint16ToUint64,
		"float32":  reduceUint16ToFloat32,
		"float64":  reduceUint16ToFloat64,
		"*string":  reduceUint16ToPtrString,
		"*bool":    reduceUint16ToPtrBool,
		"*int":     reduceUint16ToPtrInt,
		"*int8":    reduceUint16ToPtrInt8,
		"*int16":   reduceUint16ToPtrInt16,
		"*int32":   reduceUint16ToPtrInt32,
		"*int64":   reduceUint16ToPtrInt64,
		"*uint":    reduceUint16ToPtrUint,
		"*uint8":   reduceUint16ToPtrUint8,
		"*uint16":  reduceUint16ToPtrUint16,
		"*uint32":  reduceUint16ToPtrUint32,
		"*uint64":  reduceUint16ToPtrUint64,
		"*float32": reduceUint16ToPtrFloat32,
		"*float64": reduceUint16ToPtrFloat64,
	},
	"*uint16": {
		"string":   reducePtrUint16ToString,
		"bool":     reducePtrUint16ToBool,
		"int":      reducePtrUint16ToInt,
		"int8":     reducePtrUint16ToInt8,
		"int16":    reducePtrUint16ToInt16,
		"int32":    reducePtrUint16ToInt32,
		"int64":    reducePtrUint16ToInt64,
		"uint":     reducePtrUint16ToUint,
		"uint8":    reducePtrUint16ToUint8,
		"uint16":   reducePtrUint16ToUint16,
		"uint32":   reducePtrUint16ToUint32,
		"uint64":   reducePtrUint16ToUint64,
		"float32":  reducePtrUint16ToFloat32,
		"float64":  reducePtrUint16ToFloat64,
		"*string":  reducePtrUint16ToPtrString,
		"*bool":    reducePtrUint16ToPtrBool,
		"*int":     reducePtrUint16ToPtrInt,
		"*int8":    reducePtrUint16ToPtrInt8,
		"*int16":   reducePtrUint16ToPtrInt16,
		"*int32":   reducePtrUint16ToPtrInt32,
		"*int64":   reducePtrUint16ToPtrInt64,
		"*uint":    reducePtrUint16ToPtrUint,
		"*uint8":   reducePtrUint16ToPtrUint8,
		"*uint16":  reducePtrUint16ToPtrUint16,
		"*uint32":  reducePtrUint16ToPtrUint32,
		"*uint64":  reducePtrUint16ToPtrUint64,
		"*float32": reducePtrUint16ToPtrFloat32,
		"*float64": reducePtrUint16ToPtrFloat64,
	},
	"uint32": {
		"string":   reduceUint32ToString,
		"bool":     reduceUint32ToBool,
		"int":      reduceUint32ToInt,
		"int8":     reduceUint32ToInt8,
		"int16":    reduceUint32ToInt16,
		"int32":    reduceUint32ToInt32,
		"int64":    reduceUint32ToInt64,
		"uint":     reduceUint32ToUint,
		"uint8":    reduceUint32ToUint8,
		"uint16":   reduceUint32ToUint16,
		"uint32":   reduceUint32ToUint32,
		"uint64":   reduceUint32ToUint64,
		"float32":  reduceUint32ToFloat32,
		"float64":  reduceUint32ToFloat64,
		"*string":  reduceUint32ToPtrString,
		"*bool":    reduceUint32ToPtrBool,
		"*int":     reduceUint32ToPtrInt,
		"*int8":    reduceUint32ToPtrInt8,
		"*int16":   reduceUint32ToPtrInt16,
		"*int32":   reduceUint32ToPtrInt32,
		"*int64":   reduceUint32ToPtrInt64,
		"*uint":    reduceUint32ToPtrUint,
		"*uint8":   reduceUint32ToPtrUint8,
		"*uint16":  reduceUint32ToPtrUint16,
		"*uint32":  reduceUint32ToPtrUint32,
		"*uint64":  reduceUint32ToPtrUint64,
		"*float32": reduceUint32ToPtrFloat32,
		"*float64": reduceUint32ToPtrFloat64,
	},
	"*uint32": {
		"string":   reducePtrUint32ToString,
		"bool":     reducePtrUint32ToBool,
		"int":      reducePtrUint32ToInt,
		"int8":     reducePtrUint32ToInt8,
		"int16":    reducePtrUint32ToInt16,
		"int32":    reducePtrUint32ToInt32,
		"int64":    reducePtrUint32ToInt64,
		"uint":     reducePtrUint32ToUint,
		"uint8":    reducePtrUint32ToUint8,
		"uint16":   reducePtrUint32ToUint16,
		"uint32":   reducePtrUint32ToUint32,
		"uint64":   reducePtrUint32ToUint64,
		"float32":  reducePtrUint32ToFloat32,
		"float64":  reducePtrUint32ToFloat64,
		"*string":  reducePtrUint32ToPtrString,
		"*bool":    reducePtrUint32ToPtrBool,
		"*int":     reducePtrUint32ToPtrInt,
		"*int8":    reducePtrUint32ToPtrInt8,
		"*int16":   reducePtrUint32ToPtrInt16,
		"*int32":   reducePtrUint32ToPtrInt32,
		"*int64":   reducePtrUint32ToPtrInt64,
		"*uint":    reducePtrUint32ToPtrUint,
		"*uint8":   reducePtrUint32ToPtrUint8,
		"*uint16":  reducePtrUint32ToPtrUint16,
		"*uint32":  reducePtrUint32ToPtrUint32,
		"*uint64":  reducePtrUint32ToPtrUint64,
		"*float32": reducePtrUint32ToPtrFloat32,
		"*float64": reducePtrUint32ToPtrFloat64,
	},
	"uint64": {
		"string":   reduceUint64ToString,
		"bool":     reduceUint64ToBool,
		"int":      reduceUint64ToInt,
		"int8":     reduceUint64ToInt8,
		"int16":    reduceUint64ToInt16,
		"int32":    reduceUint64ToInt32,
		"int64":    reduceUint64ToInt64,
		"uint":     reduceUint64ToUint,
		"uint8":    reduceUint64ToUint8,
		"uint16":   reduceUint64ToUint16,
		"uint32":   reduceUint64ToUint32,
		"uint64":   reduceUint64ToUint64,
		"float32":  reduceUint64ToFloat32,
		"float64":  reduceUint64ToFloat64,
		"*string":  reduceUint64ToPtrString,
		"*bool":    reduceUint64ToPtrBool,
		"*int":     reduceUint64ToPtrInt,
		"*int8":    reduceUint64ToPtrInt8,
		"*int16":   reduceUint64ToPtrInt16,
		"*int32":   reduceUint64ToPtrInt32,
		"*int64":   reduceUint64ToPtrInt64,
		"*uint":    reduceUint64ToPtrUint,
		"*uint8":   reduceUint64ToPtrUint8,
		"*uint16":  reduceUint64ToPtrUint16,
		"*uint32":  reduceUint64ToPtrUint32,
		"*uint64":  reduceUint64ToPtrUint64,
		"*float32": reduceUint64ToPtrFloat32,
		"*float64": reduceUint64ToPtrFloat64,
	},
	"*uint64": {
		"string":   reducePtrUint64ToString,
		"bool":     reducePtrUint64ToBool,
		"int":      reducePtrUint64ToInt,
		"int8":     reducePtrUint64ToInt8,
		"int16":    reducePtrUint64ToInt16,
		"int32":    reducePtrUint64ToInt32,
		"int64":    reducePtrUint64ToInt64,
		"uint":     reducePtrUint64ToUint,
		"uint8":    reducePtrUint64ToUint8,
		"uint16":   reducePtrUint64ToUint16,
		"uint32":   reducePtrUint64ToUint32,
		"uint64":   reducePtrUint64ToUint64,
		"float32":  reducePtrUint64ToFloat32,
		"float64":  reducePtrUint64ToFloat64,
		"*string":  reducePtrUint64ToPtrString,
		"*bool":    reducePtrUint64ToPtrBool,
		"*int":     reducePtrUint64ToPtrInt,
		"*int8":    reducePtrUint64ToPtrInt8,
		"*int16":   reducePtrUint64ToPtrInt16,
		"*int32":   reducePtrUint64ToPtrInt32,
		"*int64":   reducePtrUint64ToPtrInt64,
		"*uint":    reducePtrUint64ToPtrUint,
		"*uint8":   reducePtrUint64ToPtrUint8,
		"*uint16":  reducePtrUint64ToPtrUint16,
		"*uint32":  reducePtrUint64ToPtrUint32,
		"*uint64":  reducePtrUint64ToPtrUint64,
		"*float32": reducePtrUint64ToPtrFloat32,
		"*float64": reducePtrUint64ToPtrFloat64,
	},
	"float32": {
		"string":   reduceFloat32ToString,
		"bool":     reduceFloat32ToBool,
		"int":      reduceFloat32ToInt,
		"int8":     reduceFloat32ToInt8,
		"int16":    reduceFloat32ToInt16,
		"int32":    reduceFloat32ToInt32,
		"int64":    reduceFloat32ToInt64,
		"uint":     reduceFloat32ToUint,
		"uint8":    reduceFloat32ToUint8,
		"uint16":   reduceFloat32ToUint16,
		"uint32":   reduceFloat32ToUint32,
		"uint64":   reduceFloat32ToUint64,
		"float32":  reduceFloat32ToFloat32,
		"float64":  reduceFloat32ToFloat64,
		"*string":  reduceFloat32ToPtrString,
		"*bool":    reduceFloat32ToPtrBool,
		"*int":     reduceFloat32ToPtrInt,
		"*int8":    reduceFloat32ToPtrInt8,
		"*int16":   reduceFloat32ToPtrInt16,
		"*int32":   reduceFloat32ToPtrInt32,
		"*int64":   reduceFloat32ToPtrInt64,
		"*uint":    reduceFloat32ToPtrUint,
		"*uint8":   reduceFloat32ToPtrUint8,
		"*uint16":  reduceFloat32ToPtrUint16,
		"*uint32":  reduceFloat32ToPtrUint32,
		"*uint64":  reduceFloat32ToPtrUint64,
		"*float32": reduceFloat32ToPtrFloat32,
		"*float64": reduceFloat32ToPtrFloat64,
	},
	"*float32": {
		"string":   reducePtrFloat32ToString,
		"bool":     reducePtrFloat32ToBool,
		"int":      reducePtrFloat32ToInt,
		"int8":     reducePtrFloat32ToInt8,
		"int16":    reducePtrFloat32ToInt16,
		"int32":    reducePtrFloat32ToInt32,
		"int64":    reducePtrFloat32ToInt64,
		"uint":     reducePtrFloat32ToUint,
		"uint8":    reducePtrFloat32ToUint8,
		"uint16":   reducePtrFloat32ToUint16,
		"uint32":   reducePtrFloat32ToUint32,
		"uint64":   reducePtrFloat32ToUint64,
		"float32":  reducePtrFloat32ToFloat32,
		"float64":  reducePtrFloat32ToFloat64,
		"*string":  reducePtrFloat32ToPtrString,
		"*bool":    reducePtrFloat32ToPtrBool,
		"*int":     reducePtrFloat32ToPtrInt,
		"*int8":    reducePtrFloat32ToPtrInt8,
		"*int16":   reducePtrFloat32ToPtrInt16,
		"*int32":   reducePtrFloat32ToPtrInt32,
		"*int64":   reducePtrFloat32ToPtrInt64,
		"*uint":    reducePtrFloat32ToPtrUint,
		"*uint8":   reducePtrFloat32ToPtrUint8,
		"*uint16":  reducePtrFloat32ToPtrUint16,
		"*uint32":  reducePtrFloat32ToPtrUint32,
		"*uint64":  reducePtrFloat32ToPtrUint64,
		"*float32": reducePtrFloat32ToPtrFloat32,
		"*float64": reducePtrFloat32ToPtrFloat64,
	},
	"float64": {
		"string":   reduceFloat64ToString,
		"bool":     reduceFloat64ToBool,
		"int":      reduceFloat64ToInt,
		"int8":     reduceFloat64ToInt8,
		"int16":    reduceFloat64ToInt16,
		"int32":    reduceFloat64ToInt32,
		"int64":    reduceFloat64ToInt64,
		"uint":     reduceFloat64ToUint,
		"uint8":    reduceFloat64ToUint8,
		"uint16":   reduceFloat64ToUint16,
		"uint32":   reduceFloat64ToUint32,
		"uint64":   reduceFloat64ToUint64,
		"float32":  reduceFloat64ToFloat32,
		"float64":  reduceFloat64ToFloat64,
		"*string":  reduceFloat64ToPtrString,
		"*bool":    reduceFloat64ToPtrBool,
		"*int":     reduceFloat64ToPtrInt,
		"*int8":    reduceFloat64ToPtrInt8,
		"*int16":   reduceFloat64ToPtrInt16,
		"*int32":   reduceFloat64ToPtrInt32,
		"*int64":   reduceFloat64ToPtrInt64,
		"*uint":    reduceFloat64ToPtrUint,
		"*uint8":   reduceFloat64ToPtrUint8,
		"*uint16":  reduceFloat64ToPtrUint16,
		"*uint32":  reduceFloat64ToPtrUint32,
		"*uint64":  reduceFloat64ToPtrUint64,
		"*float32": reduceFloat64ToPtrFloat32,
		"*float64": reduceFloat64ToPtrFloat64,
	},
	"*float64": {
		"string":   reducePtrFloat64ToString,
		"bool":     reducePtrFloat64ToBool,
		"int":      reducePtrFloat64ToInt,
		"int8":     reducePtrFloat64ToInt8,
		"int16":    reducePtrFloat64ToInt16,
		"int32":    reducePtrFloat64ToInt32,
		"int64":    reducePtrFloat64ToInt64,
		"uint":     reducePtrFloat64ToUint,
		"uint8":    reducePtrFloat64ToUint8,
		"uint16":   reducePtrFloat64ToUint16,
		"uint32":   reducePtrFloat64ToUint32,
		"uint64":   reducePtrFloat64ToUint64,
		"float32":  reducePtrFloat64ToFloat32,
		"float64":  reducePtrFloat64ToFloat64,
		"*string":  reducePtrFloat64ToPtrString,
		"*bool":    reducePtrFloat64ToPtrBool,
		"*int":     reducePtrFloat64ToPtrInt,
		"*int8":    reducePtrFloat64ToPtrInt8,
		"*int16":   reducePtrFloat64ToPtrInt16,
		"*int32":   reducePtrFloat64ToPtrInt32,
		"*int64":   reducePtrFloat64ToPtrInt64,
		"*uint":    reducePtrFloat64ToPtrUint,
		"*uint8":   reducePtrFloat64ToPtrUint8,
		"*uint16":  reducePtrFloat64ToPtrUint16,
		"*uint32":  reducePtrFloat64ToPtrUint32,
		"*uint64":  reducePtrFloat64ToPtrUint64,
		"*float32": reducePtrFloat64ToPtrFloat32,
		"*float64": reducePtrFloat64ToPtrFloat64,
	},
}

func dispatch(items reflect.Value, function interface{}, info *reduceInfo) (bool, interface{}) {
	output := info.fnIn1Type.String()
	input := info.fnIn2Type.String()
	if inputVal, ok := dispatcher[input]; ok {
		if outputVal, ok := inputVal[output]; ok {
			return true, outputVal(items, function)
		}
	}
	return false, nil
}

func reduceStringToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(string, string) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*string, string) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*string, *string) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(string, *string) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(bool, string) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*bool, string) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*bool, *string) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(bool, *string) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(int, string) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*int, string) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*int, *string) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(int, *string) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(int8, string) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*int8, string) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*int8, *string) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(int8, *string) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(int16, string) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*int16, string) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*int16, *string) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(int16, *string) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(int32, string) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*int32, string) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*int32, *string) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(int32, *string) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(int64, string) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*int64, string) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*int64, *string) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(int64, *string) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(uint, string) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*uint, string) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*uint, *string) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(uint, *string) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(uint8, string) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*uint8, string) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*uint8, *string) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(uint8, *string) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(uint16, string) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*uint16, string) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*uint16, *string) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(uint16, *string) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(uint32, string) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*uint32, string) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*uint32, *string) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(uint32, *string) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(uint64, string) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*uint64, string) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*uint64, *string) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(uint64, *string) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(float32, string) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*float32, string) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*float32, *string) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(float32, *string) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(float64, string) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceStringToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]string)
	fn := function.(func(*float64, string) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(*float64, *string) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrStringToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*string)
	fn := function.(func(float64, *string) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(string, bool) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*string, bool) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*string, *bool) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(string, *bool) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(bool, bool) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*bool, bool) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*bool, *bool) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(bool, *bool) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(int, bool) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*int, bool) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*int, *bool) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(int, *bool) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(int8, bool) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*int8, bool) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*int8, *bool) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(int8, *bool) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(int16, bool) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*int16, bool) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*int16, *bool) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(int16, *bool) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(int32, bool) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*int32, bool) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*int32, *bool) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(int32, *bool) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(int64, bool) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*int64, bool) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*int64, *bool) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(int64, *bool) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(uint, bool) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*uint, bool) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*uint, *bool) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(uint, *bool) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(uint8, bool) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*uint8, bool) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*uint8, *bool) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(uint8, *bool) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(uint16, bool) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*uint16, bool) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*uint16, *bool) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(uint16, *bool) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(uint32, bool) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*uint32, bool) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*uint32, *bool) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(uint32, *bool) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(uint64, bool) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*uint64, bool) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*uint64, *bool) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(uint64, *bool) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(float32, bool) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*float32, bool) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*float32, *bool) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(float32, *bool) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(float64, bool) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceBoolToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]bool)
	fn := function.(func(*float64, bool) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(*float64, *bool) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrBoolToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*bool)
	fn := function.(func(float64, *bool) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(string, int) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*string, int) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*string, *int) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(string, *int) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(bool, int) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*bool, int) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*bool, *int) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(bool, *int) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(int, int) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*int, int) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*int, *int) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(int, *int) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(int8, int) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*int8, int) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*int8, *int) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(int8, *int) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(int16, int) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*int16, int) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*int16, *int) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(int16, *int) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(int32, int) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*int32, int) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*int32, *int) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(int32, *int) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(int64, int) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*int64, int) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*int64, *int) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(int64, *int) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(uint, int) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*uint, int) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*uint, *int) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(uint, *int) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(uint8, int) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*uint8, int) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*uint8, *int) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(uint8, *int) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(uint16, int) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*uint16, int) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*uint16, *int) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(uint16, *int) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(uint32, int) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*uint32, int) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*uint32, *int) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(uint32, *int) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(uint64, int) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*uint64, int) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*uint64, *int) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(uint64, *int) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(float32, int) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*float32, int) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*float32, *int) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(float32, *int) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(float64, int) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceIntToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int)
	fn := function.(func(*float64, int) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(*float64, *int) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrIntToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int)
	fn := function.(func(float64, *int) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(string, int8) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*string, int8) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*string, *int8) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(string, *int8) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(bool, int8) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*bool, int8) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*bool, *int8) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(bool, *int8) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(int, int8) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*int, int8) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*int, *int8) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(int, *int8) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(int8, int8) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*int8, int8) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*int8, *int8) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(int8, *int8) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(int16, int8) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*int16, int8) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*int16, *int8) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(int16, *int8) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(int32, int8) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*int32, int8) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*int32, *int8) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(int32, *int8) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(int64, int8) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*int64, int8) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*int64, *int8) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(int64, *int8) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(uint, int8) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*uint, int8) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*uint, *int8) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(uint, *int8) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(uint8, int8) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*uint8, int8) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*uint8, *int8) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(uint8, *int8) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(uint16, int8) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*uint16, int8) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*uint16, *int8) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(uint16, *int8) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(uint32, int8) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*uint32, int8) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*uint32, *int8) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(uint32, *int8) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(uint64, int8) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*uint64, int8) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*uint64, *int8) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(uint64, *int8) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(float32, int8) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*float32, int8) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*float32, *int8) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(float32, *int8) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(float64, int8) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt8ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int8)
	fn := function.(func(*float64, int8) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(*float64, *int8) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt8ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int8)
	fn := function.(func(float64, *int8) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(string, int16) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*string, int16) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*string, *int16) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(string, *int16) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(bool, int16) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*bool, int16) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*bool, *int16) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(bool, *int16) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(int, int16) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*int, int16) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*int, *int16) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(int, *int16) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(int8, int16) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*int8, int16) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*int8, *int16) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(int8, *int16) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(int16, int16) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*int16, int16) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*int16, *int16) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(int16, *int16) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(int32, int16) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*int32, int16) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*int32, *int16) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(int32, *int16) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(int64, int16) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*int64, int16) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*int64, *int16) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(int64, *int16) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(uint, int16) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*uint, int16) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*uint, *int16) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(uint, *int16) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(uint8, int16) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*uint8, int16) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*uint8, *int16) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(uint8, *int16) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(uint16, int16) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*uint16, int16) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*uint16, *int16) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(uint16, *int16) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(uint32, int16) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*uint32, int16) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*uint32, *int16) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(uint32, *int16) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(uint64, int16) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*uint64, int16) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*uint64, *int16) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(uint64, *int16) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(float32, int16) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*float32, int16) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*float32, *int16) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(float32, *int16) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(float64, int16) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt16ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int16)
	fn := function.(func(*float64, int16) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(*float64, *int16) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt16ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int16)
	fn := function.(func(float64, *int16) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(string, int32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*string, int32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*string, *int32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(string, *int32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(bool, int32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*bool, int32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*bool, *int32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(bool, *int32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(int, int32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*int, int32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*int, *int32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(int, *int32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(int8, int32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*int8, int32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*int8, *int32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(int8, *int32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(int16, int32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*int16, int32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*int16, *int32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(int16, *int32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(int32, int32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*int32, int32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*int32, *int32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(int32, *int32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(int64, int32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*int64, int32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*int64, *int32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(int64, *int32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(uint, int32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*uint, int32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*uint, *int32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(uint, *int32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(uint8, int32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*uint8, int32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*uint8, *int32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(uint8, *int32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(uint16, int32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*uint16, int32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*uint16, *int32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(uint16, *int32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(uint32, int32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*uint32, int32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*uint32, *int32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(uint32, *int32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(uint64, int32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*uint64, int32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*uint64, *int32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(uint64, *int32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(float32, int32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*float32, int32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*float32, *int32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(float32, *int32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(float64, int32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int32)
	fn := function.(func(*float64, int32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(*float64, *int32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int32)
	fn := function.(func(float64, *int32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(string, int64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*string, int64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*string, *int64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(string, *int64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(bool, int64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*bool, int64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*bool, *int64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(bool, *int64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(int, int64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*int, int64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*int, *int64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(int, *int64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(int8, int64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*int8, int64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*int8, *int64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(int8, *int64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(int16, int64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*int16, int64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*int16, *int64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(int16, *int64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(int32, int64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*int32, int64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*int32, *int64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(int32, *int64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(int64, int64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*int64, int64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*int64, *int64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(int64, *int64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(uint, int64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*uint, int64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*uint, *int64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(uint, *int64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(uint8, int64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*uint8, int64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*uint8, *int64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(uint8, *int64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(uint16, int64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*uint16, int64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*uint16, *int64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(uint16, *int64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(uint32, int64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*uint32, int64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*uint32, *int64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(uint32, *int64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(uint64, int64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*uint64, int64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*uint64, *int64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(uint64, *int64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(float32, int64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*float32, int64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*float32, *int64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(float32, *int64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(float64, int64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceInt64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]int64)
	fn := function.(func(*float64, int64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(*float64, *int64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrInt64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*int64)
	fn := function.(func(float64, *int64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(string, uint) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*string, uint) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*string, *uint) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(string, *uint) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(bool, uint) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*bool, uint) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*bool, *uint) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(bool, *uint) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(int, uint) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*int, uint) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*int, *uint) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(int, *uint) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(int8, uint) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*int8, uint) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*int8, *uint) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(int8, *uint) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(int16, uint) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*int16, uint) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*int16, *uint) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(int16, *uint) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(int32, uint) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*int32, uint) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*int32, *uint) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(int32, *uint) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(int64, uint) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*int64, uint) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*int64, *uint) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(int64, *uint) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(uint, uint) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*uint, uint) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint, *uint) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(uint, *uint) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(uint8, uint) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*uint8, uint) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint8, *uint) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(uint8, *uint) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(uint16, uint) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*uint16, uint) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint16, *uint) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(uint16, *uint) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(uint32, uint) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*uint32, uint) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint32, *uint) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(uint32, *uint) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(uint64, uint) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*uint64, uint) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*uint64, *uint) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(uint64, *uint) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(float32, uint) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*float32, uint) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*float32, *uint) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(float32, *uint) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(float64, uint) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUintToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint)
	fn := function.(func(*float64, uint) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(*float64, *uint) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUintToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint)
	fn := function.(func(float64, *uint) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(string, uint8) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*string, uint8) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*string, *uint8) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(string, *uint8) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(bool, uint8) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*bool, uint8) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*bool, *uint8) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(bool, *uint8) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(int, uint8) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*int, uint8) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*int, *uint8) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(int, *uint8) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(int8, uint8) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*int8, uint8) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*int8, *uint8) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(int8, *uint8) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(int16, uint8) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*int16, uint8) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*int16, *uint8) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(int16, *uint8) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(int32, uint8) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*int32, uint8) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*int32, *uint8) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(int32, *uint8) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(int64, uint8) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*int64, uint8) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*int64, *uint8) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(int64, *uint8) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(uint, uint8) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*uint, uint8) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint, *uint8) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(uint, *uint8) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(uint8, uint8) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*uint8, uint8) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint8, *uint8) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(uint8, *uint8) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(uint16, uint8) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*uint16, uint8) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint16, *uint8) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(uint16, *uint8) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(uint32, uint8) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*uint32, uint8) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint32, *uint8) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(uint32, *uint8) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(uint64, uint8) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*uint64, uint8) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*uint64, *uint8) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(uint64, *uint8) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(float32, uint8) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*float32, uint8) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*float32, *uint8) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(float32, *uint8) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(float64, uint8) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint8ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint8)
	fn := function.(func(*float64, uint8) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(*float64, *uint8) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint8ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint8)
	fn := function.(func(float64, *uint8) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(string, uint16) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*string, uint16) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*string, *uint16) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(string, *uint16) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(bool, uint16) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*bool, uint16) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*bool, *uint16) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(bool, *uint16) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(int, uint16) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*int, uint16) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*int, *uint16) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(int, *uint16) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(int8, uint16) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*int8, uint16) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*int8, *uint16) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(int8, *uint16) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(int16, uint16) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*int16, uint16) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*int16, *uint16) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(int16, *uint16) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(int32, uint16) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*int32, uint16) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*int32, *uint16) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(int32, *uint16) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(int64, uint16) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*int64, uint16) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*int64, *uint16) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(int64, *uint16) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(uint, uint16) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*uint, uint16) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint, *uint16) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(uint, *uint16) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(uint8, uint16) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*uint8, uint16) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint8, *uint16) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(uint8, *uint16) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(uint16, uint16) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*uint16, uint16) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint16, *uint16) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(uint16, *uint16) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(uint32, uint16) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*uint32, uint16) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint32, *uint16) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(uint32, *uint16) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(uint64, uint16) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*uint64, uint16) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*uint64, *uint16) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(uint64, *uint16) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(float32, uint16) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*float32, uint16) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*float32, *uint16) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(float32, *uint16) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(float64, uint16) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint16ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint16)
	fn := function.(func(*float64, uint16) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(*float64, *uint16) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint16ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint16)
	fn := function.(func(float64, *uint16) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(string, uint32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*string, uint32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*string, *uint32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(string, *uint32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(bool, uint32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*bool, uint32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*bool, *uint32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(bool, *uint32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(int, uint32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*int, uint32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*int, *uint32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(int, *uint32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(int8, uint32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*int8, uint32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*int8, *uint32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(int8, *uint32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(int16, uint32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*int16, uint32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*int16, *uint32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(int16, *uint32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(int32, uint32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*int32, uint32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*int32, *uint32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(int32, *uint32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(int64, uint32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*int64, uint32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*int64, *uint32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(int64, *uint32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(uint, uint32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*uint, uint32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint, *uint32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(uint, *uint32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(uint8, uint32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*uint8, uint32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint8, *uint32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(uint8, *uint32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(uint16, uint32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*uint16, uint32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint16, *uint32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(uint16, *uint32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(uint32, uint32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*uint32, uint32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint32, *uint32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(uint32, *uint32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(uint64, uint32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*uint64, uint32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*uint64, *uint32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(uint64, *uint32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(float32, uint32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*float32, uint32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*float32, *uint32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(float32, *uint32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(float64, uint32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint32)
	fn := function.(func(*float64, uint32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(*float64, *uint32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint32)
	fn := function.(func(float64, *uint32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(string, uint64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*string, uint64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*string, *uint64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(string, *uint64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(bool, uint64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*bool, uint64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*bool, *uint64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(bool, *uint64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(int, uint64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*int, uint64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*int, *uint64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(int, *uint64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(int8, uint64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*int8, uint64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*int8, *uint64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(int8, *uint64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(int16, uint64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*int16, uint64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*int16, *uint64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(int16, *uint64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(int32, uint64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*int32, uint64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*int32, *uint64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(int32, *uint64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(int64, uint64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*int64, uint64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*int64, *uint64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(int64, *uint64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(uint, uint64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*uint, uint64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint, *uint64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(uint, *uint64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(uint8, uint64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*uint8, uint64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint8, *uint64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(uint8, *uint64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(uint16, uint64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*uint16, uint64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint16, *uint64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(uint16, *uint64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(uint32, uint64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*uint32, uint64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint32, *uint64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(uint32, *uint64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(uint64, uint64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*uint64, uint64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*uint64, *uint64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(uint64, *uint64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(float32, uint64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*float32, uint64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*float32, *uint64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(float32, *uint64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(float64, uint64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceUint64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]uint64)
	fn := function.(func(*float64, uint64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(*float64, *uint64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrUint64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*uint64)
	fn := function.(func(float64, *uint64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(string, float32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*string, float32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*string, *float32) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(string, *float32) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(bool, float32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*bool, float32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*bool, *float32) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(bool, *float32) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(int, float32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*int, float32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*int, *float32) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(int, *float32) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(int8, float32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*int8, float32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*int8, *float32) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(int8, *float32) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(int16, float32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*int16, float32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*int16, *float32) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(int16, *float32) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(int32, float32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*int32, float32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*int32, *float32) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(int32, *float32) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(int64, float32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*int64, float32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*int64, *float32) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(int64, *float32) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(uint, float32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*uint, float32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*uint, *float32) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(uint, *float32) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(uint8, float32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*uint8, float32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*uint8, *float32) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(uint8, *float32) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(uint16, float32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*uint16, float32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*uint16, *float32) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(uint16, *float32) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(uint32, float32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*uint32, float32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*uint32, *float32) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(uint32, *float32) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(uint64, float32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*uint64, float32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*uint64, *float32) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(uint64, *float32) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(float32, float32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*float32, float32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*float32, *float32) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(float32, *float32) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(float64, float32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float32)
	fn := function.(func(*float64, float32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(*float64, *float32) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat32ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float32)
	fn := function.(func(float64, *float32) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(string, float64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*string, float64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*string, *float64) *string)
	acc := ""
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToString(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(string, *float64) string)
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(bool, float64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*bool, float64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*bool, *float64) *bool)
	acc := false
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToBool(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(bool, *float64) bool)
	acc := false
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(int, float64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*int, float64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*int, *float64) *int)
	acc := 0
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToInt(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(int, *float64) int)
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(int8, float64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*int8, float64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*int8, *float64) *int8)
	acc := int8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToInt8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(int8, *float64) int8)
	acc := int8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(int16, float64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*int16, float64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*int16, *float64) *int16)
	acc := int16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToInt16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(int16, *float64) int16)
	acc := int16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(int32, float64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*int32, float64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*int32, *float64) *int32)
	acc := int32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToInt32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(int32, *float64) int32)
	acc := int32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(int64, float64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*int64, float64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*int64, *float64) *int64)
	acc := int64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToInt64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(int64, *float64) int64)
	acc := int64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(uint, float64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*uint, float64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*uint, *float64) *uint)
	acc := uint(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToUint(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(uint, *float64) uint)
	acc := uint(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(uint8, float64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*uint8, float64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*uint8, *float64) *uint8)
	acc := uint8(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToUint8(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(uint8, *float64) uint8)
	acc := uint8(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(uint16, float64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*uint16, float64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*uint16, *float64) *uint16)
	acc := uint16(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToUint16(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(uint16, *float64) uint16)
	acc := uint16(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(uint32, float64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*uint32, float64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*uint32, *float64) *uint32)
	acc := uint32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToUint32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(uint32, *float64) uint32)
	acc := uint32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(uint64, float64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*uint64, float64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*uint64, *float64) *uint64)
	acc := uint64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToUint64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(uint64, *float64) uint64)
	acc := uint64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(float32, float64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*float32, float64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*float32, *float64) *float32)
	acc := float32(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToFloat32(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(float32, *float64) float32)
	acc := float32(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(float64, float64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

func reduceFloat64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]float64)
	fn := function.(func(*float64, float64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToPtrFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(*float64, *float64) *float64)
	acc := float64(0)
	accPtr := &acc
	for _, item := range items {
		accPtr = fn(accPtr, item)
	}
	return accPtr
}

func reducePtrFloat64ToFloat64(itemsValue reflect.Value, function interface{}) interface{} {
	items := itemsValue.Interface().([]*float64)
	fn := function.(func(float64, *float64) float64)
	acc := float64(0)
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}
