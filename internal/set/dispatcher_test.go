package set

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_setString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), utils.ArrayString[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayString[1], output.([]string)[0])
}

func Test_setStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), utils.ArrayStringPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayStringPtr[1], output.([]*string)[0])
}

func Test_setBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), utils.ArrayBool[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayBool[1], output.([]bool)[0])
}

func Test_setBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), utils.ArrayBoolPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayBoolPtr[1], output.([]*bool)[0])
}

func Test_setInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), utils.ArrayInt[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt[1], output.([]int)[0])
}

func Test_setIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), utils.ArrayIntPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayIntPtr[1], output.([]*int)[0])
}

func Test_setInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), utils.ArrayInt8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt8[1], output.([]int8)[0])
}

func Test_setInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), utils.ArrayInt8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt8Ptr[1], output.([]*int8)[0])
}

func Test_setInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), utils.ArrayInt16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt16[1], output.([]int16)[0])
}

func Test_setInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), utils.ArrayInt16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt16Ptr[1], output.([]*int16)[0])
}

func Test_setInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), utils.ArrayInt32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt32[1], output.([]int32)[0])
}

func Test_setInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), utils.ArrayInt32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt32Ptr[1], output.([]*int32)[0])
}

func Test_setInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), utils.ArrayInt64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt64[1], output.([]int64)[0])
}

func Test_setInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), utils.ArrayInt64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayInt64Ptr[1], output.([]*int64)[0])
}

func Test_setUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), utils.ArrayUint[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint[1], output.([]uint)[0])
}

func Test_setUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), utils.ArrayUintPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUintPtr[1], output.([]*uint)[0])
}

func Test_setUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), utils.ArrayUint8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint8[1], output.([]uint8)[0])
}

func Test_setUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), utils.ArrayUint8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint8Ptr[1], output.([]*uint8)[0])
}

func Test_setUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), utils.ArrayUint16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint16[1], output.([]uint16)[0])
}

func Test_setUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), utils.ArrayUint16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint16Ptr[1], output.([]*uint16)[0])
}

func Test_setUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), utils.ArrayUint32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint32[1], output.([]uint32)[0])
}

func Test_setUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), utils.ArrayUint32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint32Ptr[1], output.([]*uint32)[0])
}

func Test_setUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), utils.ArrayUint64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint64[1], output.([]uint64)[0])
}

func Test_setUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), utils.ArrayUint64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayUint64Ptr[1], output.([]*uint64)[0])
}

func Test_setFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), utils.ArrayFloat32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayFloat32[1], output.([]float32)[0])
}

func Test_setFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), utils.ArrayFloat32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayFloat32Ptr[1], output.([]*float32)[0])
}

func Test_setFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), utils.ArrayFloat64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayFloat64[1], output.([]float64)[0])
}

func Test_setFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	info := &setInfo{itemType: &typeElement}
	info.index = 0
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), utils.ArrayFloat64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 5)
	assert.Equal(t, utils.ArrayFloat64Ptr[1], output.([]*float64)[0])
}
