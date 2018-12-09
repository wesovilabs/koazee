package maps

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_mapStringToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ string) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ string) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *string) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *string) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ string) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ string) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *string) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *string) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ string) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ string) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *string) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *string) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ string) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ string) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *string) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *string) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ string) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ string) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *string) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *string) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ string) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ string) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *string) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *string) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ string) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ string) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *string) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *string) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ string) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ string) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *string) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *string) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ string) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ string) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *string) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *string) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ string) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ string) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *string) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *string) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ string) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ string) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *string) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *string) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ string) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ string) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *string) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *string) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ string) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ string) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *string) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *string) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapStringToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ string) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapStringToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayString[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ string) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *string) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrStringToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *string) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ bool) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ bool) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *bool) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *bool) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ bool) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ bool) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *bool) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *bool) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ bool) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ bool) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *bool) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *bool) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ bool) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ bool) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *bool) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *bool) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ bool) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ bool) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *bool) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *bool) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ bool) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ bool) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *bool) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *bool) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ bool) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ bool) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *bool) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *bool) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ bool) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ bool) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *bool) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *bool) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ bool) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ bool) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *bool) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *bool) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ bool) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ bool) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *bool) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *bool) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ bool) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ bool) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *bool) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *bool) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ bool) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ bool) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *bool) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *bool) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ bool) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ bool) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *bool) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *bool) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ bool) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapBoolToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ bool) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *bool) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrBoolToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *bool) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapIntToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapIntToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrIntToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int8) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int8) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int8) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int8) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int8) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int8) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int8) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int8) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int8) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int8) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int8) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int8) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int8) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int8) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int8) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int8) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int8) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int8) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int8) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int8) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int8) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int8) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int8) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int8) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int8) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int8) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int8) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int8) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int8) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int8) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int8) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int8) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int8) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int8) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int8) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int8) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int8) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int8) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int8) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int8) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int8) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int8) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int8) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int8) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int8) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int8) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int8) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int8) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int8) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int8) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int8) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int8) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int8) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapInt8ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int8) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int8) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt8ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int8) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int16) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int16) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int16) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int16) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int16) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int16) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int16) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int16) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int16) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int16) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int16) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int16) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int16) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int16) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int16) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int16) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int16) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int16) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int16) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int16) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int16) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int16) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int16) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int16) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int16) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int16) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int16) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int16) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int16) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int16) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int16) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int16) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int16) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int16) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int16) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int16) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int16) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int16) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int16) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int16) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int16) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int16) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int16) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int16) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int16) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int16) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int16) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int16) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int16) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int16) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int16) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int16) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int16) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapInt16ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int16) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int16) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt16ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int16) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapInt32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapInt64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrInt64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUintToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapUintToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUintToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint8) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint8) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint8) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint8) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint8) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint8) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint8) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint8) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint8) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint8) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint8) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint8) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint8) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint8) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint8) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint8) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint8) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint8) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint8) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint8) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint8) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint8) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint8) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint8) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint8) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint8) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint8) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint8) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint8) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint8) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint8) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint8) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint8) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint8) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint8) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint8) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint8) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint8) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint8) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint8) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint8) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint8) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint8) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint8) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint8) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint8) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint8) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint8) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint8) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint8) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint8) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint8) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint8) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapUint8ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint8) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint8) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint8ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint8) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint16) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint16) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint16) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint16) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint16) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint16) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint16) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint16) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint16) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint16) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint16) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint16) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint16) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint16) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint16) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint16) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint16) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint16) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint16) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint16) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint16) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint16) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint16) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint16) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint16) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint16) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint16) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint16) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint16) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint16) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint16) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint16) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint16) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint16) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint16) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint16) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint16) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint16) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint16) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint16) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint16) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint16) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint16) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint16) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint16) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint16) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint16) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint16) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint16) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint16) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint16) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint16) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint16) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapUint16ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint16) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint16) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint16ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint16) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapUint32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapUint64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrUint64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ float32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ float32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *float32) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *float32) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ float32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ float32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *float32) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *float32) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ float32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ float32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *float32) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *float32) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ float32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ float32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *float32) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *float32) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ float32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ float32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *float32) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *float32) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ float32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ float32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *float32) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *float32) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ float32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ float32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *float32) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *float32) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ float32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ float32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *float32) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *float32) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ float32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ float32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *float32) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *float32) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ float32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ float32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *float32) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *float32) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ float32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ float32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *float32) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *float32) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ float32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ float32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *float32) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *float32) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ float32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ float32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *float32) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *float32) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ float32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapFloat32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ float32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *float32) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat32ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *float32) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ float64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ float64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *float64) string { return utils.ValueString }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrString(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *float64) *string { return utils.ValueStringPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ float64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ float64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *float64) bool { return utils.ValueBool }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrBool(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *float64) *bool { return utils.ValueBoolPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ float64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ float64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *float64) int { return utils.ValueInt }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrInt(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *float64) *int { return utils.ValueIntPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ float64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ float64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *float64) int8 { return utils.ValueInt8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrInt8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *float64) *int8 { return utils.ValueInt8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ float64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ float64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *float64) int16 { return utils.ValueInt16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrInt16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *float64) *int16 { return utils.ValueInt16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ float64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ float64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *float64) int32 { return utils.ValueInt32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrInt32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *float64) *int32 { return utils.ValueInt32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ float64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ float64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *float64) int64 { return utils.ValueInt64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrInt64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *float64) *int64 { return utils.ValueInt64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ float64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ float64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *float64) uint { return utils.ValueUint }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrUint(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *float64) *uint { return utils.ValueUintPtr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ float64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ float64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *float64) uint8 { return utils.ValueUint8 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrUint8(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *float64) *uint8 { return utils.ValueUint8Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ float64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ float64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *float64) uint16 { return utils.ValueUint16 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrUint16(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *float64) *uint16 { return utils.ValueUint16Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ float64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ float64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *float64) uint32 { return utils.ValueUint32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrUint32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *float64) *uint32 { return utils.ValueUint32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ float64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ float64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *float64) uint64 { return utils.ValueUint64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrUint64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *float64) *uint64 { return utils.ValueUint64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ float64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ float64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *float64) float32 { return utils.ValueFloat32 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrFloat32(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *float64) *float32 { return utils.ValueFloat32Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ float64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapFloat64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ float64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *float64) float64 { return utils.ValueFloat64 }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}

func Test_mapPtrFloat64ToPtrFloat64(t *testing.T) {
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	outputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *float64) *float64 { return utils.ValueFloat64Ptr }
	info := &mapInfo{fnInputType: inputType, fnOutputType: outputType}
	found, _ := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
}
