package drop

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_dropString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), utils.ArrayString[1], info)
	assert.True(t, found)
	for index := range output.([]string) {
		assert.NotEqual(t, utils.ArrayString[1], output.([]string)[index])
	}
}

func Test_dropStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), utils.ArrayStringPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*string) {
		assert.NotEqual(t, *utils.ArrayStringPtr[1], *output.([]*string)[index])
	}
}

func Test_dropBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), utils.ArrayBool[1], info)
	assert.True(t, found)
	for index := range output.([]bool) {
		assert.NotEqual(t, utils.ArrayBool[1], output.([]bool)[index])
	}
}

func Test_dropBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), utils.ArrayBoolPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*bool) {
		assert.NotEqual(t, *utils.ArrayBoolPtr[1], *output.([]*bool)[index])
	}
}

func Test_dropInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), utils.ArrayInt[1], info)
	assert.True(t, found)
	for index := range output.([]int) {
		assert.NotEqual(t, utils.ArrayInt[1], output.([]int)[index])
	}
}

func Test_dropIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), utils.ArrayIntPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*int) {
		assert.NotEqual(t, *utils.ArrayIntPtr[1], *output.([]*int)[index])
	}
}

func Test_dropInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), utils.ArrayInt8[1], info)
	assert.True(t, found)
	for index := range output.([]int8) {
		assert.NotEqual(t, utils.ArrayInt8[1], output.([]int8)[index])
	}
}

func Test_dropInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), utils.ArrayInt8Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int8) {
		assert.NotEqual(t, *utils.ArrayInt8Ptr[1], *output.([]*int8)[index])
	}
}

func Test_dropInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), utils.ArrayInt16[1], info)
	assert.True(t, found)
	for index := range output.([]int16) {
		assert.NotEqual(t, utils.ArrayInt16[1], output.([]int16)[index])
	}
}

func Test_dropInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), utils.ArrayInt16Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int16) {
		assert.NotEqual(t, *utils.ArrayInt16Ptr[1], *output.([]*int16)[index])
	}
}

func Test_dropInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), utils.ArrayInt32[1], info)
	assert.True(t, found)
	for index := range output.([]int32) {
		assert.NotEqual(t, utils.ArrayInt32[1], output.([]int32)[index])
	}
}

func Test_dropInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), utils.ArrayInt32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int32) {
		assert.NotEqual(t, *utils.ArrayInt32Ptr[1], *output.([]*int32)[index])
	}
}

func Test_dropInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), utils.ArrayInt64[1], info)
	assert.True(t, found)
	for index := range output.([]int64) {
		assert.NotEqual(t, utils.ArrayInt64[1], output.([]int64)[index])
	}
}

func Test_dropInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), utils.ArrayInt64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*int64) {
		assert.NotEqual(t, *utils.ArrayInt64Ptr[1], *output.([]*int64)[index])
	}
}

func Test_dropUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), utils.ArrayUint[1], info)
	assert.True(t, found)
	for index := range output.([]uint) {
		assert.NotEqual(t, utils.ArrayUint[1], output.([]uint)[index])
	}
}

func Test_dropUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), utils.ArrayUintPtr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint) {
		assert.NotEqual(t, *utils.ArrayUintPtr[1], *output.([]*uint)[index])
	}
}

func Test_dropUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), utils.ArrayUint8[1], info)
	assert.True(t, found)
	for index := range output.([]uint8) {
		assert.NotEqual(t, utils.ArrayUint8[1], output.([]uint8)[index])
	}
}

func Test_dropUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), utils.ArrayUint8Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint8) {
		assert.NotEqual(t, *utils.ArrayUint8Ptr[1], *output.([]*uint8)[index])
	}
}

func Test_dropUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), utils.ArrayUint16[1], info)
	assert.True(t, found)
	for index := range output.([]uint16) {
		assert.NotEqual(t, utils.ArrayUint16[1], output.([]uint16)[index])
	}
}

func Test_dropUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), utils.ArrayUint16Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint16) {
		assert.NotEqual(t, *utils.ArrayUint16Ptr[1], *output.([]*uint16)[index])
	}
}

func Test_dropUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), utils.ArrayUint32[1], info)
	assert.True(t, found)
	for index := range output.([]uint32) {
		assert.NotEqual(t, utils.ArrayUint32[1], output.([]uint32)[index])
	}
}

func Test_dropUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), utils.ArrayUint32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint32) {
		assert.NotEqual(t, *utils.ArrayUint32Ptr[1], *output.([]*uint32)[index])
	}
}

func Test_dropUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), utils.ArrayUint64[1], info)
	assert.True(t, found)
	for index := range output.([]uint64) {
		assert.NotEqual(t, utils.ArrayUint64[1], output.([]uint64)[index])
	}
}

func Test_dropUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), utils.ArrayUint64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*uint64) {
		assert.NotEqual(t, *utils.ArrayUint64Ptr[1], *output.([]*uint64)[index])
	}
}

func Test_dropFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), utils.ArrayFloat32[1], info)
	assert.True(t, found)
	for index := range output.([]float32) {
		assert.NotEqual(t, utils.ArrayFloat32[1], output.([]float32)[index])
	}
}

func Test_dropFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), utils.ArrayFloat32Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*float32) {
		assert.NotEqual(t, *utils.ArrayFloat32Ptr[1], *output.([]*float32)[index])
	}
}

func Test_dropFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), utils.ArrayFloat64[1], info)
	assert.True(t, found)
	for index := range output.([]float64) {
		assert.NotEqual(t, utils.ArrayFloat64[1], output.([]float64)[index])
	}
}

func Test_dropFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	info := &dropInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), utils.ArrayFloat64Ptr[1], info)
	assert.True(t, found)
	for index := range output.([]*float64) {
		assert.NotEqual(t, *utils.ArrayFloat64Ptr[1], *output.([]*float64)[index])
	}
}
