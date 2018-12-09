package reduce

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_reduceStringToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ string, _ string) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceStringToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *string, _ string) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrStringToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ string, _ *string) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrStringToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *string, _ *string) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceStringToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ bool, _ string) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceStringToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *bool, _ string) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrStringToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ bool, _ *string) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrStringToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *bool, _ *string) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceStringToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int, _ string) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceStringToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int, _ string) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrStringToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int, _ *string) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrStringToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int, _ *string) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceStringToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int8, _ string) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceStringToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int8, _ string) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrStringToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int8, _ *string) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrStringToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int8, _ *string) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceStringToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int16, _ string) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceStringToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int16, _ string) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrStringToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int16, _ *string) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrStringToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int16, _ *string) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceStringToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int32, _ string) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceStringToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int32, _ string) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrStringToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int32, _ *string) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrStringToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int32, _ *string) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceStringToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ int64, _ string) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceStringToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *int64, _ string) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrStringToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ int64, _ *string) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrStringToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *int64, _ *string) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceStringToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint, _ string) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceStringToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint, _ string) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrStringToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint, _ *string) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrStringToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint, _ *string) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceStringToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint8, _ string) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceStringToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint8, _ string) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrStringToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint8, _ *string) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrStringToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint8, _ *string) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceStringToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint16, _ string) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceStringToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint16, _ string) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrStringToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint16, _ *string) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrStringToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint16, _ *string) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceStringToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint32, _ string) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceStringToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint32, _ string) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrStringToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint32, _ *string) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrStringToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint32, _ *string) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceStringToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ uint64, _ string) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceStringToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *uint64, _ string) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrStringToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ uint64, _ *string) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrStringToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *uint64, _ *string) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceStringToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ float32, _ string) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceStringToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *float32, _ string) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrStringToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ float32, _ *string) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrStringToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *float32, _ *string) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceStringToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ float64, _ string) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceStringToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ *float64, _ string) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrStringToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ float64, _ *string) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrStringToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *float64, _ *string) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceBoolToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ string, _ bool) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceBoolToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *string, _ bool) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrBoolToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ string, _ *bool) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrBoolToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *string, _ *bool) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceBoolToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ bool, _ bool) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceBoolToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *bool, _ bool) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrBoolToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ bool, _ *bool) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrBoolToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *bool, _ *bool) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceBoolToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int, _ bool) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceBoolToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int, _ bool) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrBoolToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int, _ *bool) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrBoolToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int, _ *bool) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceBoolToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int8, _ bool) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceBoolToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int8, _ bool) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrBoolToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int8, _ *bool) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrBoolToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int8, _ *bool) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceBoolToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int16, _ bool) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceBoolToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int16, _ bool) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrBoolToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int16, _ *bool) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrBoolToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int16, _ *bool) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceBoolToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int32, _ bool) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceBoolToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int32, _ bool) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrBoolToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int32, _ *bool) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrBoolToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int32, _ *bool) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceBoolToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ int64, _ bool) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceBoolToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *int64, _ bool) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrBoolToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ int64, _ *bool) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrBoolToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *int64, _ *bool) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceBoolToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint, _ bool) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceBoolToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint, _ bool) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrBoolToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint, _ *bool) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrBoolToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint, _ *bool) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceBoolToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint8, _ bool) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceBoolToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint8, _ bool) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrBoolToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint8, _ *bool) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrBoolToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint8, _ *bool) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceBoolToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint16, _ bool) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceBoolToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint16, _ bool) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrBoolToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint16, _ *bool) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrBoolToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint16, _ *bool) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceBoolToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint32, _ bool) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceBoolToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint32, _ bool) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrBoolToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint32, _ *bool) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrBoolToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint32, _ *bool) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceBoolToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ uint64, _ bool) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceBoolToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *uint64, _ bool) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrBoolToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ uint64, _ *bool) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrBoolToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *uint64, _ *bool) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceBoolToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ float32, _ bool) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceBoolToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *float32, _ bool) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrBoolToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ float32, _ *bool) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrBoolToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *float32, _ *bool) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceBoolToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ float64, _ bool) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceBoolToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ *float64, _ bool) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrBoolToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ float64, _ *bool) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrBoolToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *float64, _ *bool) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceIntToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ string, _ int) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceIntToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *string, _ int) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrIntToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ string, _ *int) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrIntToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *string, _ *int) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceIntToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ bool, _ int) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceIntToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *bool, _ int) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrIntToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ bool, _ *int) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrIntToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *bool, _ *int) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceIntToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int, _ int) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceIntToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int, _ int) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrIntToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int, _ *int) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrIntToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int, _ *int) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceIntToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int8, _ int) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceIntToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int8, _ int) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrIntToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int8, _ *int) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrIntToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int8, _ *int) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceIntToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int16, _ int) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceIntToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int16, _ int) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrIntToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int16, _ *int) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrIntToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int16, _ *int) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceIntToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int32, _ int) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceIntToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int32, _ int) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrIntToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int32, _ *int) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrIntToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int32, _ *int) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceIntToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int64, _ int) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceIntToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *int64, _ int) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrIntToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ int64, _ *int) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrIntToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int64, _ *int) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceIntToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint, _ int) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceIntToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint, _ int) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrIntToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint, _ *int) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrIntToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint, _ *int) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceIntToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint8, _ int) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceIntToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint8, _ int) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrIntToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint8, _ *int) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrIntToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint8, _ *int) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceIntToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint16, _ int) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceIntToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint16, _ int) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrIntToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint16, _ *int) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrIntToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint16, _ *int) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceIntToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint32, _ int) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceIntToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint32, _ int) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrIntToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint32, _ *int) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrIntToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint32, _ *int) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceIntToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ uint64, _ int) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceIntToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *uint64, _ int) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrIntToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ uint64, _ *int) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrIntToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *uint64, _ *int) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceIntToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ float32, _ int) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceIntToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *float32, _ int) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrIntToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ float32, _ *int) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrIntToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *float32, _ *int) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceIntToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ float64, _ int) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceIntToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ *float64, _ int) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrIntToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ float64, _ *int) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrIntToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *float64, _ *int) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceInt8ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ string, _ int8) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceInt8ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *string, _ int8) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrInt8ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ string, _ *int8) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrInt8ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *string, _ *int8) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceInt8ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ bool, _ int8) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceInt8ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *bool, _ int8) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrInt8ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ bool, _ *int8) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrInt8ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *bool, _ *int8) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceInt8ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int, _ int8) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceInt8ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int, _ int8) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrInt8ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int, _ *int8) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrInt8ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int, _ *int8) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceInt8ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int8, _ int8) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceInt8ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int8, _ int8) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrInt8ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int8, _ *int8) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrInt8ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int8, _ *int8) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceInt8ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int16, _ int8) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceInt8ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int16, _ int8) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrInt8ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int16, _ *int8) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrInt8ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int16, _ *int8) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceInt8ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int32, _ int8) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceInt8ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int32, _ int8) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrInt8ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int32, _ *int8) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrInt8ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int32, _ *int8) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceInt8ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int64, _ int8) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceInt8ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *int64, _ int8) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrInt8ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ int64, _ *int8) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrInt8ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int64, _ *int8) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceInt8ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint, _ int8) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceInt8ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint, _ int8) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrInt8ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint, _ *int8) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrInt8ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint, _ *int8) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceInt8ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint8, _ int8) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceInt8ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint8, _ int8) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrInt8ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint8, _ *int8) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrInt8ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint8, _ *int8) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceInt8ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint16, _ int8) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceInt8ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint16, _ int8) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrInt8ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint16, _ *int8) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrInt8ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint16, _ *int8) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceInt8ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint32, _ int8) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceInt8ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint32, _ int8) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrInt8ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint32, _ *int8) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrInt8ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint32, _ *int8) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceInt8ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ uint64, _ int8) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceInt8ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *uint64, _ int8) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrInt8ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ uint64, _ *int8) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrInt8ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *uint64, _ *int8) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceInt8ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ float32, _ int8) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceInt8ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *float32, _ int8) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrInt8ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ float32, _ *int8) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrInt8ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *float32, _ *int8) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceInt8ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ float64, _ int8) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceInt8ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ *float64, _ int8) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrInt8ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ float64, _ *int8) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrInt8ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *float64, _ *int8) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceInt16ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ string, _ int16) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceInt16ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *string, _ int16) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrInt16ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ string, _ *int16) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrInt16ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *string, _ *int16) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceInt16ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ bool, _ int16) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceInt16ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *bool, _ int16) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrInt16ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ bool, _ *int16) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrInt16ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *bool, _ *int16) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceInt16ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int, _ int16) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceInt16ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int, _ int16) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrInt16ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int, _ *int16) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrInt16ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int, _ *int16) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceInt16ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int8, _ int16) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceInt16ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int8, _ int16) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrInt16ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int8, _ *int16) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrInt16ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int8, _ *int16) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceInt16ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int16, _ int16) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceInt16ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int16, _ int16) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrInt16ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int16, _ *int16) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrInt16ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int16, _ *int16) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceInt16ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int32, _ int16) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceInt16ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int32, _ int16) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrInt16ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int32, _ *int16) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrInt16ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int32, _ *int16) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceInt16ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int64, _ int16) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceInt16ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *int64, _ int16) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrInt16ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ int64, _ *int16) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrInt16ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int64, _ *int16) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceInt16ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint, _ int16) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceInt16ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint, _ int16) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrInt16ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint, _ *int16) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrInt16ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint, _ *int16) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceInt16ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint8, _ int16) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceInt16ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint8, _ int16) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrInt16ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint8, _ *int16) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrInt16ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint8, _ *int16) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceInt16ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint16, _ int16) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceInt16ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint16, _ int16) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrInt16ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint16, _ *int16) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrInt16ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint16, _ *int16) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceInt16ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint32, _ int16) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceInt16ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint32, _ int16) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrInt16ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint32, _ *int16) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrInt16ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint32, _ *int16) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceInt16ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ uint64, _ int16) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceInt16ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *uint64, _ int16) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrInt16ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ uint64, _ *int16) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrInt16ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *uint64, _ *int16) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceInt16ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ float32, _ int16) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceInt16ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *float32, _ int16) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrInt16ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ float32, _ *int16) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrInt16ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *float32, _ *int16) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceInt16ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ float64, _ int16) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceInt16ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ *float64, _ int16) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrInt16ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ float64, _ *int16) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrInt16ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *float64, _ *int16) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceInt32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ string, _ int32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceInt32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *string, _ int32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrInt32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ string, _ *int32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrInt32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *string, _ *int32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceInt32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ bool, _ int32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceInt32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *bool, _ int32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrInt32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ bool, _ *int32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrInt32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *bool, _ *int32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceInt32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int, _ int32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceInt32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int, _ int32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrInt32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int, _ *int32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrInt32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int, _ *int32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceInt32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int8, _ int32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceInt32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int8, _ int32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrInt32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int8, _ *int32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrInt32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int8, _ *int32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceInt32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int16, _ int32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceInt32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int16, _ int32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrInt32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int16, _ *int32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrInt32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int16, _ *int32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceInt32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int32, _ int32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceInt32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int32, _ int32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrInt32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int32, _ *int32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrInt32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int32, _ *int32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceInt32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int64, _ int32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceInt32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *int64, _ int32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrInt32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ int64, _ *int32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrInt32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int64, _ *int32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceInt32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint, _ int32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceInt32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint, _ int32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrInt32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint, _ *int32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrInt32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint, _ *int32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceInt32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint8, _ int32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceInt32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint8, _ int32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrInt32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint8, _ *int32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrInt32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint8, _ *int32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceInt32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint16, _ int32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceInt32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint16, _ int32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrInt32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint16, _ *int32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrInt32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint16, _ *int32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceInt32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint32, _ int32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceInt32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint32, _ int32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrInt32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint32, _ *int32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrInt32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint32, _ *int32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceInt32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ uint64, _ int32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceInt32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *uint64, _ int32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrInt32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ uint64, _ *int32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrInt32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *uint64, _ *int32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceInt32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ float32, _ int32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceInt32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *float32, _ int32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrInt32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ float32, _ *int32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrInt32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *float32, _ *int32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceInt32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ float64, _ int32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceInt32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ *float64, _ int32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrInt32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ float64, _ *int32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrInt32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *float64, _ *int32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceInt64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ string, _ int64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceInt64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *string, _ int64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrInt64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ string, _ *int64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrInt64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *string, _ *int64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceInt64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ bool, _ int64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceInt64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *bool, _ int64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrInt64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ bool, _ *int64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrInt64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *bool, _ *int64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceInt64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int, _ int64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceInt64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int, _ int64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrInt64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int, _ *int64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrInt64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int, _ *int64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceInt64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int8, _ int64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceInt64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int8, _ int64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrInt64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int8, _ *int64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrInt64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int8, _ *int64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceInt64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int16, _ int64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceInt64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int16, _ int64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrInt64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int16, _ *int64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrInt64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int16, _ *int64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceInt64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int32, _ int64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceInt64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int32, _ int64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrInt64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int32, _ *int64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrInt64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int32, _ *int64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceInt64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int64, _ int64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceInt64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *int64, _ int64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrInt64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ int64, _ *int64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrInt64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int64, _ *int64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceInt64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint, _ int64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceInt64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint, _ int64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrInt64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint, _ *int64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrInt64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint, _ *int64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceInt64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint8, _ int64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceInt64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint8, _ int64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrInt64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint8, _ *int64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrInt64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint8, _ *int64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceInt64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint16, _ int64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceInt64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint16, _ int64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrInt64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint16, _ *int64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrInt64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint16, _ *int64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceInt64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint32, _ int64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceInt64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint32, _ int64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrInt64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint32, _ *int64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrInt64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint32, _ *int64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceInt64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ uint64, _ int64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceInt64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *uint64, _ int64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrInt64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ uint64, _ *int64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrInt64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *uint64, _ *int64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceInt64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ float32, _ int64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceInt64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *float32, _ int64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrInt64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ float32, _ *int64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrInt64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *float32, _ *int64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceInt64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ float64, _ int64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceInt64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ *float64, _ int64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrInt64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ float64, _ *int64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrInt64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *float64, _ *int64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceUintToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ string, _ uint) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceUintToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *string, _ uint) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrUintToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ string, _ *uint) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrUintToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *string, _ *uint) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceUintToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ bool, _ uint) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceUintToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *bool, _ uint) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrUintToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ bool, _ *uint) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrUintToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *bool, _ *uint) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceUintToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int, _ uint) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceUintToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int, _ uint) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrUintToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int, _ *uint) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrUintToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int, _ *uint) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceUintToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int8, _ uint) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceUintToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int8, _ uint) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrUintToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int8, _ *uint) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrUintToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int8, _ *uint) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceUintToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int16, _ uint) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceUintToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int16, _ uint) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrUintToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int16, _ *uint) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrUintToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int16, _ *uint) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceUintToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int32, _ uint) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceUintToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int32, _ uint) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrUintToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int32, _ *uint) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrUintToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int32, _ *uint) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceUintToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ int64, _ uint) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceUintToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *int64, _ uint) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrUintToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ int64, _ *uint) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrUintToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *int64, _ *uint) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceUintToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint, _ uint) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceUintToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint, _ uint) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrUintToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint, _ *uint) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrUintToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint, _ *uint) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceUintToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint8, _ uint) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceUintToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint8, _ uint) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrUintToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint8, _ *uint) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrUintToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint8, _ *uint) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceUintToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint16, _ uint) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceUintToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint16, _ uint) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrUintToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint16, _ *uint) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrUintToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint16, _ *uint) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceUintToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint32, _ uint) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceUintToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint32, _ uint) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrUintToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint32, _ *uint) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrUintToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint32, _ *uint) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceUintToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint64, _ uint) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceUintToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *uint64, _ uint) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrUintToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ uint64, _ *uint) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrUintToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint64, _ *uint) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceUintToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ float32, _ uint) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceUintToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *float32, _ uint) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrUintToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ float32, _ *uint) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrUintToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *float32, _ *uint) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceUintToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ float64, _ uint) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceUintToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ *float64, _ uint) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrUintToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ float64, _ *uint) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrUintToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *float64, _ *uint) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceUint8ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ string, _ uint8) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceUint8ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *string, _ uint8) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrUint8ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ string, _ *uint8) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrUint8ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *string, _ *uint8) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceUint8ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ bool, _ uint8) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceUint8ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *bool, _ uint8) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrUint8ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ bool, _ *uint8) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrUint8ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *bool, _ *uint8) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceUint8ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int, _ uint8) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceUint8ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int, _ uint8) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrUint8ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int, _ *uint8) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrUint8ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int, _ *uint8) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceUint8ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int8, _ uint8) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceUint8ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int8, _ uint8) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrUint8ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int8, _ *uint8) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrUint8ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int8, _ *uint8) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceUint8ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int16, _ uint8) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceUint8ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int16, _ uint8) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrUint8ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int16, _ *uint8) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrUint8ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int16, _ *uint8) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceUint8ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int32, _ uint8) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceUint8ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int32, _ uint8) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrUint8ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int32, _ *uint8) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrUint8ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int32, _ *uint8) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceUint8ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ int64, _ uint8) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceUint8ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *int64, _ uint8) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrUint8ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ int64, _ *uint8) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrUint8ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *int64, _ *uint8) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceUint8ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint, _ uint8) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceUint8ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint, _ uint8) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrUint8ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint, _ *uint8) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrUint8ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint, _ *uint8) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceUint8ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint8, _ uint8) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceUint8ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint8, _ uint8) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrUint8ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint8, _ *uint8) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrUint8ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint8, _ *uint8) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceUint8ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint16, _ uint8) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceUint8ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint16, _ uint8) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrUint8ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint16, _ *uint8) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrUint8ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint16, _ *uint8) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceUint8ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint32, _ uint8) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceUint8ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint32, _ uint8) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrUint8ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint32, _ *uint8) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrUint8ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint32, _ *uint8) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceUint8ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint64, _ uint8) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceUint8ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *uint64, _ uint8) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrUint8ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ uint64, _ *uint8) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrUint8ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint64, _ *uint8) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceUint8ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ float32, _ uint8) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceUint8ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *float32, _ uint8) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrUint8ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ float32, _ *uint8) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrUint8ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *float32, _ *uint8) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceUint8ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ float64, _ uint8) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceUint8ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ *float64, _ uint8) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrUint8ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ float64, _ *uint8) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrUint8ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *float64, _ *uint8) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceUint16ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ string, _ uint16) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceUint16ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *string, _ uint16) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrUint16ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ string, _ *uint16) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrUint16ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *string, _ *uint16) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceUint16ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ bool, _ uint16) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceUint16ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *bool, _ uint16) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrUint16ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ bool, _ *uint16) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrUint16ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *bool, _ *uint16) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceUint16ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int, _ uint16) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceUint16ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int, _ uint16) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrUint16ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int, _ *uint16) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrUint16ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int, _ *uint16) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceUint16ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int8, _ uint16) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceUint16ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int8, _ uint16) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrUint16ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int8, _ *uint16) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrUint16ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int8, _ *uint16) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceUint16ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int16, _ uint16) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceUint16ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int16, _ uint16) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrUint16ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int16, _ *uint16) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrUint16ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int16, _ *uint16) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceUint16ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int32, _ uint16) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceUint16ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int32, _ uint16) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrUint16ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int32, _ *uint16) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrUint16ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int32, _ *uint16) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceUint16ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ int64, _ uint16) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceUint16ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *int64, _ uint16) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrUint16ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ int64, _ *uint16) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrUint16ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *int64, _ *uint16) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceUint16ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint, _ uint16) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceUint16ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint, _ uint16) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrUint16ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint, _ *uint16) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrUint16ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint, _ *uint16) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceUint16ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint8, _ uint16) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceUint16ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint8, _ uint16) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrUint16ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint8, _ *uint16) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrUint16ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint8, _ *uint16) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceUint16ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint16, _ uint16) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceUint16ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint16, _ uint16) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrUint16ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint16, _ *uint16) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrUint16ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint16, _ *uint16) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceUint16ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint32, _ uint16) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceUint16ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint32, _ uint16) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrUint16ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint32, _ *uint16) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrUint16ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint32, _ *uint16) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceUint16ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint64, _ uint16) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceUint16ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *uint64, _ uint16) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrUint16ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ uint64, _ *uint16) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrUint16ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint64, _ *uint16) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceUint16ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ float32, _ uint16) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceUint16ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *float32, _ uint16) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrUint16ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ float32, _ *uint16) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrUint16ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *float32, _ *uint16) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceUint16ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ float64, _ uint16) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceUint16ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ *float64, _ uint16) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrUint16ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ float64, _ *uint16) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrUint16ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *float64, _ *uint16) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceUint32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ string, _ uint32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceUint32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *string, _ uint32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrUint32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ string, _ *uint32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrUint32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *string, _ *uint32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceUint32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ bool, _ uint32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceUint32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *bool, _ uint32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrUint32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ bool, _ *uint32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrUint32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *bool, _ *uint32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceUint32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int, _ uint32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceUint32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int, _ uint32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrUint32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int, _ *uint32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrUint32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int, _ *uint32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceUint32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int8, _ uint32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceUint32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int8, _ uint32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrUint32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int8, _ *uint32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrUint32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int8, _ *uint32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceUint32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int16, _ uint32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceUint32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int16, _ uint32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrUint32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int16, _ *uint32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrUint32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int16, _ *uint32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceUint32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int32, _ uint32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceUint32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int32, _ uint32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrUint32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int32, _ *uint32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrUint32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int32, _ *uint32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceUint32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ int64, _ uint32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceUint32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *int64, _ uint32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrUint32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ int64, _ *uint32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrUint32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *int64, _ *uint32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceUint32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint, _ uint32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceUint32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint, _ uint32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrUint32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint, _ *uint32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrUint32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint, _ *uint32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceUint32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint8, _ uint32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceUint32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint8, _ uint32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrUint32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint8, _ *uint32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrUint32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint8, _ *uint32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceUint32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint16, _ uint32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceUint32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint16, _ uint32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrUint32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint16, _ *uint32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrUint32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint16, _ *uint32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceUint32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint32, _ uint32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceUint32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint32, _ uint32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrUint32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint32, _ *uint32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrUint32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint32, _ *uint32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceUint32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint64, _ uint32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceUint32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *uint64, _ uint32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrUint32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ uint64, _ *uint32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrUint32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint64, _ *uint32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceUint32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ float32, _ uint32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceUint32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *float32, _ uint32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrUint32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ float32, _ *uint32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrUint32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *float32, _ *uint32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceUint32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ float64, _ uint32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceUint32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ *float64, _ uint32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrUint32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ float64, _ *uint32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrUint32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *float64, _ *uint32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceUint64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ string, _ uint64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceUint64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *string, _ uint64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrUint64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ string, _ *uint64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrUint64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *string, _ *uint64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceUint64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ bool, _ uint64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceUint64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *bool, _ uint64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrUint64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ bool, _ *uint64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrUint64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *bool, _ *uint64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceUint64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int, _ uint64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceUint64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int, _ uint64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrUint64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int, _ *uint64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrUint64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int, _ *uint64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceUint64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int8, _ uint64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceUint64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int8, _ uint64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrUint64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int8, _ *uint64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrUint64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int8, _ *uint64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceUint64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int16, _ uint64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceUint64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int16, _ uint64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrUint64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int16, _ *uint64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrUint64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int16, _ *uint64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceUint64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int32, _ uint64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceUint64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int32, _ uint64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrUint64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int32, _ *uint64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrUint64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int32, _ *uint64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceUint64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ int64, _ uint64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceUint64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *int64, _ uint64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrUint64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ int64, _ *uint64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrUint64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *int64, _ *uint64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceUint64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint, _ uint64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceUint64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint, _ uint64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrUint64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint, _ *uint64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrUint64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint, _ *uint64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceUint64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint8, _ uint64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceUint64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint8, _ uint64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrUint64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint8, _ *uint64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrUint64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint8, _ *uint64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceUint64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint16, _ uint64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceUint64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint16, _ uint64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrUint64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint16, _ *uint64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrUint64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint16, _ *uint64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceUint64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint32, _ uint64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceUint64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint32, _ uint64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrUint64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint32, _ *uint64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrUint64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint32, _ *uint64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceUint64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint64, _ uint64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceUint64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *uint64, _ uint64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrUint64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ uint64, _ *uint64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrUint64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint64, _ *uint64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceUint64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ float32, _ uint64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceUint64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *float32, _ uint64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrUint64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ float32, _ *uint64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrUint64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *float32, _ *uint64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceUint64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ float64, _ uint64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceUint64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ *float64, _ uint64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrUint64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ float64, _ *uint64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrUint64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *float64, _ *uint64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceFloat32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ string, _ float32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceFloat32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *string, _ float32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrFloat32ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ string, _ *float32) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrFloat32ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *string, _ *float32) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceFloat32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ bool, _ float32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceFloat32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *bool, _ float32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrFloat32ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ bool, _ *float32) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrFloat32ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *bool, _ *float32) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceFloat32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int, _ float32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceFloat32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int, _ float32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrFloat32ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int, _ *float32) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrFloat32ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int, _ *float32) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceFloat32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int8, _ float32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceFloat32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int8, _ float32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrFloat32ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int8, _ *float32) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrFloat32ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int8, _ *float32) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceFloat32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int16, _ float32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceFloat32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int16, _ float32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrFloat32ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int16, _ *float32) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrFloat32ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int16, _ *float32) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceFloat32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int32, _ float32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceFloat32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int32, _ float32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrFloat32ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int32, _ *float32) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrFloat32ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int32, _ *float32) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceFloat32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ int64, _ float32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceFloat32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *int64, _ float32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrFloat32ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ int64, _ *float32) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrFloat32ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *int64, _ *float32) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceFloat32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint, _ float32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceFloat32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint, _ float32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrFloat32ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint, _ *float32) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrFloat32ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint, _ *float32) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceFloat32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint8, _ float32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceFloat32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint8, _ float32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrFloat32ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint8, _ *float32) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrFloat32ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint8, _ *float32) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceFloat32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint16, _ float32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceFloat32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint16, _ float32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrFloat32ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint16, _ *float32) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrFloat32ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint16, _ *float32) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceFloat32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint32, _ float32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceFloat32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint32, _ float32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrFloat32ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint32, _ *float32) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrFloat32ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint32, _ *float32) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceFloat32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ uint64, _ float32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceFloat32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *uint64, _ float32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrFloat32ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ uint64, _ *float32) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrFloat32ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *uint64, _ *float32) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceFloat32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ float32, _ float32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceFloat32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *float32, _ float32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrFloat32ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ float32, _ *float32) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrFloat32ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *float32, _ *float32) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceFloat32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ float64, _ float32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceFloat32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ *float64, _ float32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrFloat32ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ float64, _ *float32) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrFloat32ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *float64, _ *float32) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reduceFloat64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ string, _ float64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reduceFloat64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *string, _ float64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reducePtrFloat64ToString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayString[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ string, _ *float64) string { return utils.ValueString }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueString, output)
}

func Test_reducePtrFloat64ToPtrString(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayStringPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *string, _ *float64) *string { return utils.ValueStringPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueStringPtr, output)
}

func Test_reduceFloat64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ bool, _ float64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reduceFloat64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *bool, _ float64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reducePtrFloat64ToBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBool[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ bool, _ *float64) bool { return utils.ValueBool }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBool, output)
}

func Test_reducePtrFloat64ToPtrBool(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *bool, _ *float64) *bool { return utils.ValueBoolPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueBoolPtr, output)
}

func Test_reduceFloat64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int, _ float64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reduceFloat64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int, _ float64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reducePtrFloat64ToInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int, _ *float64) int { return utils.ValueInt }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt, output)
}

func Test_reducePtrFloat64ToPtrInt(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayIntPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int, _ *float64) *int { return utils.ValueIntPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueIntPtr, output)
}

func Test_reduceFloat64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int8, _ float64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reduceFloat64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int8, _ float64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reducePtrFloat64ToInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int8, _ *float64) int8 { return utils.ValueInt8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8, output)
}

func Test_reducePtrFloat64ToPtrInt8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int8, _ *float64) *int8 { return utils.ValueInt8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt8Ptr, output)
}

func Test_reduceFloat64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int16, _ float64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reduceFloat64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int16, _ float64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reducePtrFloat64ToInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int16, _ *float64) int16 { return utils.ValueInt16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16, output)
}

func Test_reducePtrFloat64ToPtrInt16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int16, _ *float64) *int16 { return utils.ValueInt16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt16Ptr, output)
}

func Test_reduceFloat64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int32, _ float64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reduceFloat64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int32, _ float64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reducePtrFloat64ToInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int32, _ *float64) int32 { return utils.ValueInt32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32, output)
}

func Test_reducePtrFloat64ToPtrInt32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int32, _ *float64) *int32 { return utils.ValueInt32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt32Ptr, output)
}

func Test_reduceFloat64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ int64, _ float64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reduceFloat64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *int64, _ float64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reducePtrFloat64ToInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ int64, _ *float64) int64 { return utils.ValueInt64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64, output)
}

func Test_reducePtrFloat64ToPtrInt64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *int64, _ *float64) *int64 { return utils.ValueInt64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueInt64Ptr, output)
}

func Test_reduceFloat64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint, _ float64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reduceFloat64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint, _ float64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reducePtrFloat64ToUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint, _ *float64) uint { return utils.ValueUint }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint, output)
}

func Test_reducePtrFloat64ToPtrUint(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUintPtr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint, _ *float64) *uint { return utils.ValueUintPtr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUintPtr, output)
}

func Test_reduceFloat64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint8, _ float64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reduceFloat64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint8, _ float64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reducePtrFloat64ToUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint8, _ *float64) uint8 { return utils.ValueUint8 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8, output)
}

func Test_reducePtrFloat64ToPtrUint8(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint8, _ *float64) *uint8 { return utils.ValueUint8Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint8Ptr, output)
}

func Test_reduceFloat64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint16, _ float64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reduceFloat64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint16, _ float64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reducePtrFloat64ToUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint16, _ *float64) uint16 { return utils.ValueUint16 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16, output)
}

func Test_reducePtrFloat64ToPtrUint16(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint16, _ *float64) *uint16 { return utils.ValueUint16Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint16Ptr, output)
}

func Test_reduceFloat64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint32, _ float64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reduceFloat64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint32, _ float64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reducePtrFloat64ToUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint32, _ *float64) uint32 { return utils.ValueUint32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32, output)
}

func Test_reducePtrFloat64ToPtrUint32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint32, _ *float64) *uint32 { return utils.ValueUint32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint32Ptr, output)
}

func Test_reduceFloat64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ uint64, _ float64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reduceFloat64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *uint64, _ float64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reducePtrFloat64ToUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ uint64, _ *float64) uint64 { return utils.ValueUint64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64, output)
}

func Test_reducePtrFloat64ToPtrUint64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *uint64, _ *float64) *uint64 { return utils.ValueUint64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueUint64Ptr, output)
}

func Test_reduceFloat64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ float32, _ float64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reduceFloat64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *float32, _ float64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reducePtrFloat64ToFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ float32, _ *float64) float32 { return utils.ValueFloat32 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32, output)
}

func Test_reducePtrFloat64ToPtrFloat32(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *float32, _ *float64) *float32 { return utils.ValueFloat32Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat32Ptr, output)
}

func Test_reduceFloat64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ float64, _ float64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reduceFloat64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ *float64, _ float64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}

func Test_reducePtrFloat64ToFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ float64, _ *float64) float64 { return utils.ValueFloat64 }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64, output)
}

func Test_reducePtrFloat64ToPtrFloat64(t *testing.T) {
	fnIn1Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fnIn2Type := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *float64, _ *float64) *float64 { return utils.ValueFloat64Ptr }
	info := &reduceInfo{fnIn1Type: fnIn1Type, fnIn2Type: fnIn2Type}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, utils.ValueFloat64Ptr, output)
}
