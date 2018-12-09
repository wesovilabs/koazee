package add

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_addString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_string[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_string), utils.Array_string[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_string[1], output.([]string)[5])
}

func Test_addStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_stringPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_stringPtr), utils.Array_stringPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_stringPtr[1], output.([]*string)[5])
}

func Test_addBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_bool[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_bool), utils.Array_bool[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_bool[1], output.([]bool)[5])
}

func Test_addBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_boolPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_boolPtr), utils.Array_boolPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_boolPtr[1], output.([]*bool)[5])
}

func Test_addInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int), utils.Array_int[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int[1], output.([]int)[5])
}

func Test_addIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_intPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_intPtr), utils.Array_intPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_intPtr[1], output.([]*int)[5])
}

func Test_addInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int8[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int8), utils.Array_int8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int8[1], output.([]int8)[5])
}

func Test_addInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int8Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int8Ptr), utils.Array_int8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int8Ptr[1], output.([]*int8)[5])
}

func Test_addInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int16[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int16), utils.Array_int16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int16[1], output.([]int16)[5])
}

func Test_addInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int16Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int16Ptr), utils.Array_int16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int16Ptr[1], output.([]*int16)[5])
}

func Test_addInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int32), utils.Array_int32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int32[1], output.([]int32)[5])
}

func Test_addInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int32Ptr), utils.Array_int32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int32Ptr[1], output.([]*int32)[5])
}

func Test_addInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int64), utils.Array_int64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int64[1], output.([]int64)[5])
}

func Test_addInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int64Ptr), utils.Array_int64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_int64Ptr[1], output.([]*int64)[5])
}

func Test_addUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint), utils.Array_uint[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint[1], output.([]uint)[5])
}

func Test_addUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uintPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uintPtr), utils.Array_uintPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uintPtr[1], output.([]*uint)[5])
}

func Test_addUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint8[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8), utils.Array_uint8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint8[1], output.([]uint8)[5])
}

func Test_addUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint8Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8Ptr), utils.Array_uint8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint8Ptr[1], output.([]*uint8)[5])
}

func Test_addUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint16[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16), utils.Array_uint16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint16[1], output.([]uint16)[5])
}

func Test_addUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint16Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16Ptr), utils.Array_uint16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint16Ptr[1], output.([]*uint16)[5])
}

func Test_addUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32), utils.Array_uint32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint32[1], output.([]uint32)[5])
}

func Test_addUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32Ptr), utils.Array_uint32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint32Ptr[1], output.([]*uint32)[5])
}

func Test_addUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64), utils.Array_uint64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint64[1], output.([]uint64)[5])
}

func Test_addUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64Ptr), utils.Array_uint64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_uint64Ptr[1], output.([]*uint64)[5])
}

func Test_addFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float32), utils.Array_float32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_float32[1], output.([]float32)[5])
}

func Test_addFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float32Ptr), utils.Array_float32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_float32Ptr[1], output.([]*float32)[5])
}

func Test_addFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float64), utils.Array_float64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_float64[1], output.([]float64)[5])
}

func Test_addFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float64Ptr), utils.Array_float64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.Array_float64Ptr[1], output.([]*float64)[5])
}
