package add

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_addString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), utils.ArrayString[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayString[1], output.([]string)[5])
}

func Test_addStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), utils.ArrayStringPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayStringPtr[1], output.([]*string)[5])
}

func Test_addBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), utils.ArrayBool[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayBool[1], output.([]bool)[5])
}

func Test_addBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), utils.ArrayBoolPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayBoolPtr[1], output.([]*bool)[5])
}

func Test_addInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), utils.ArrayInt[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt[1], output.([]int)[5])
}

func Test_addIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), utils.ArrayIntPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayIntPtr[1], output.([]*int)[5])
}

func Test_addInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), utils.ArrayInt8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt8[1], output.([]int8)[5])
}

func Test_addInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), utils.ArrayInt8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt8Ptr[1], output.([]*int8)[5])
}

func Test_addInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), utils.ArrayInt16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt16[1], output.([]int16)[5])
}

func Test_addInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), utils.ArrayInt16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt16Ptr[1], output.([]*int16)[5])
}

func Test_addInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), utils.ArrayInt32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt32[1], output.([]int32)[5])
}

func Test_addInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), utils.ArrayInt32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt32Ptr[1], output.([]*int32)[5])
}

func Test_addInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), utils.ArrayInt64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt64[1], output.([]int64)[5])
}

func Test_addInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), utils.ArrayInt64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayInt64Ptr[1], output.([]*int64)[5])
}

func Test_addUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), utils.ArrayUint[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint[1], output.([]uint)[5])
}

func Test_addUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), utils.ArrayUintPtr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUintPtr[1], output.([]*uint)[5])
}

func Test_addUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), utils.ArrayUint8[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint8[1], output.([]uint8)[5])
}

func Test_addUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), utils.ArrayUint8Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint8Ptr[1], output.([]*uint8)[5])
}

func Test_addUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), utils.ArrayUint16[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint16[1], output.([]uint16)[5])
}

func Test_addUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), utils.ArrayUint16Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint16Ptr[1], output.([]*uint16)[5])
}

func Test_addUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), utils.ArrayUint32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint32[1], output.([]uint32)[5])
}

func Test_addUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), utils.ArrayUint32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint32Ptr[1], output.([]*uint32)[5])
}

func Test_addUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), utils.ArrayUint64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint64[1], output.([]uint64)[5])
}

func Test_addUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), utils.ArrayUint64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayUint64Ptr[1], output.([]*uint64)[5])
}

func Test_addFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), utils.ArrayFloat32[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayFloat32[1], output.([]float32)[5])
}

func Test_addFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), utils.ArrayFloat32Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayFloat32Ptr[1], output.([]*float32)[5])
}

func Test_addFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), utils.ArrayFloat64[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayFloat64[1], output.([]float64)[5])
}

func Test_addFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	info := &addInfo{itemType: &typeElement}
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), utils.ArrayFloat64Ptr[1], info)
	assert.True(t, found)
	assert.Len(t, output, 6)
	assert.Equal(t, utils.ArrayFloat64Ptr[1], output.([]*float64)[5])
}
