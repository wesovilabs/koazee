package drop

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_dropString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_string[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_string), utils.Array_string[1], info)
	assert.True(t, found)
	for index := range output.([]string) {
		assert.NotEqual(t, utils.Array_string[1], output.([]string)[index])
	}
}

func Test_dropStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_stringPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_stringPtr), utils.Array_stringPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*string) {
		assert.NotEqual(t, *utils.Array_stringPtr[1], *output.([]*string)[index])
	}
}

func Test_dropBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_bool[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_bool), utils.Array_bool[1], info)
	assert.True(t, found)
	for index := range output.([]bool) {
		assert.NotEqual(t, utils.Array_bool[1], output.([]bool)[index])
	}
}

func Test_dropBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_boolPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_boolPtr), utils.Array_boolPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*bool) {
		assert.NotEqual(t, *utils.Array_boolPtr[1], *output.([]*bool)[index])
	}
}

func Test_dropInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int), utils.Array_int[1], info)
	assert.True(t, found)
	for index := range output.([]int) {
		assert.NotEqual(t, utils.Array_int[1], output.([]int)[index])
	}
}

func Test_dropIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_intPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_intPtr), utils.Array_intPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*int) {
		assert.NotEqual(t, *utils.Array_intPtr[1], *output.([]*int)[index])
	}
}

func Test_dropInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int8[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int8), utils.Array_int8[1], info)
	assert.True(t, found)
	for index := range output.([]int8) {
		assert.NotEqual(t, utils.Array_int8[1], output.([]int8)[index])
	}
}

func Test_dropInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int8Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int8Ptr), utils.Array_int8Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int8) {
		assert.NotEqual(t, *utils.Array_int8Ptr[1], *output.([]*int8)[index])
	}
}

func Test_dropInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int16[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int16), utils.Array_int16[1], info)
	assert.True(t, found)
	for index := range output.([]int16) {
		assert.NotEqual(t, utils.Array_int16[1], output.([]int16)[index])
	}
}

func Test_dropInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int16Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int16Ptr), utils.Array_int16Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int16) {
		assert.NotEqual(t, *utils.Array_int16Ptr[1], *output.([]*int16)[index])
	}
}

func Test_dropInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int32), utils.Array_int32[1], info)
	assert.True(t, found)
	for index := range output.([]int32) {
		assert.NotEqual(t, utils.Array_int32[1], output.([]int32)[index])
	}
}

func Test_dropInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int32Ptr), utils.Array_int32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int32) {
		assert.NotEqual(t, *utils.Array_int32Ptr[1], *output.([]*int32)[index])
	}
}

func Test_dropInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int64), utils.Array_int64[1], info)
	assert.True(t, found)
	for index := range output.([]int64) {
		assert.NotEqual(t, utils.Array_int64[1], output.([]int64)[index])
	}
}

func Test_dropInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_int64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_int64Ptr), utils.Array_int64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int64) {
		assert.NotEqual(t, *utils.Array_int64Ptr[1], *output.([]*int64)[index])
	}
}

func Test_dropUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint), utils.Array_uint[1], info)
	assert.True(t, found)
	for index := range output.([]uint) {
		assert.NotEqual(t, utils.Array_uint[1], output.([]uint)[index])
	}
}

func Test_dropUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uintPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uintPtr), utils.Array_uintPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint) {
		assert.NotEqual(t, *utils.Array_uintPtr[1], *output.([]*uint)[index])
	}
}

func Test_dropUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint8[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8), utils.Array_uint8[1], info)
	assert.True(t, found)
	for index := range output.([]uint8) {
		assert.NotEqual(t, utils.Array_uint8[1], output.([]uint8)[index])
	}
}

func Test_dropUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint8Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8Ptr), utils.Array_uint8Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint8) {
		assert.NotEqual(t, *utils.Array_uint8Ptr[1], *output.([]*uint8)[index])
	}
}

func Test_dropUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint16[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16), utils.Array_uint16[1], info)
	assert.True(t, found)
	for index := range output.([]uint16) {
		assert.NotEqual(t, utils.Array_uint16[1], output.([]uint16)[index])
	}
}

func Test_dropUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint16Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16Ptr), utils.Array_uint16Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint16) {
		assert.NotEqual(t, *utils.Array_uint16Ptr[1], *output.([]*uint16)[index])
	}
}

func Test_dropUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32), utils.Array_uint32[1], info)
	assert.True(t, found)
	for index := range output.([]uint32) {
		assert.NotEqual(t, utils.Array_uint32[1], output.([]uint32)[index])
	}
}

func Test_dropUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32Ptr), utils.Array_uint32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint32) {
		assert.NotEqual(t, *utils.Array_uint32Ptr[1], *output.([]*uint32)[index])
	}
}

func Test_dropUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64), utils.Array_uint64[1], info)
	assert.True(t, found)
	for index := range output.([]uint64) {
		assert.NotEqual(t, utils.Array_uint64[1], output.([]uint64)[index])
	}
}

func Test_dropUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_uint64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64Ptr), utils.Array_uint64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint64) {
		assert.NotEqual(t, *utils.Array_uint64Ptr[1], *output.([]*uint64)[index])
	}
}

func Test_dropFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float32), utils.Array_float32[1], info)
	assert.True(t, found)
	for index := range output.([]float32) {
		assert.NotEqual(t, utils.Array_float32[1], output.([]float32)[index])
	}
}

func Test_dropFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float32Ptr), utils.Array_float32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*float32) {
		assert.NotEqual(t, *utils.Array_float32Ptr[1], *output.([]*float32)[index])
	}
}

func Test_dropFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float64), utils.Array_float64[1], info)
	assert.True(t, found)
	for index := range output.([]float64) {
		assert.NotEqual(t, utils.Array_float64[1], output.([]float64)[index])
	}
}

func Test_dropFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.Array_float64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.Array_float64Ptr), utils.Array_float64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*float64) {
		assert.NotEqual(t, *utils.Array_float64Ptr[1], *output.([]*float64)[index])
	}
}
