package pop

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_popString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayString), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayString[0], element)
	assert.Equal(t, utils.ArrayString[1:], output.([]string))
}

func Test_popStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayStringPtr[0], element)
	assert.Equal(t, utils.ArrayStringPtr[1:], output.([]*string))
}

func Test_popBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayBool), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayBool[0], element)
	assert.Equal(t, utils.ArrayBool[1:], output.([]bool))
}

func Test_popBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayBoolPtr[0], element)
	assert.Equal(t, utils.ArrayBoolPtr[1:], output.([]*bool))
}

func Test_popInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt[0], element)
	assert.Equal(t, utils.ArrayInt[1:], output.([]int))
}

func Test_popIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayIntPtr[0], element)
	assert.Equal(t, utils.ArrayIntPtr[1:], output.([]*int))
}

func Test_popInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt8), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt8[0], element)
	assert.Equal(t, utils.ArrayInt8[1:], output.([]int8))
}

func Test_popInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt8Ptr[0], element)
	assert.Equal(t, utils.ArrayInt8Ptr[1:], output.([]*int8))
}

func Test_popInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt16), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt16[0], element)
	assert.Equal(t, utils.ArrayInt16[1:], output.([]int16))
}

func Test_popInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt16Ptr[0], element)
	assert.Equal(t, utils.ArrayInt16Ptr[1:], output.([]*int16))
}

func Test_popInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt32), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt32[0], element)
	assert.Equal(t, utils.ArrayInt32[1:], output.([]int32))
}

func Test_popInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt32Ptr[0], element)
	assert.Equal(t, utils.ArrayInt32Ptr[1:], output.([]*int32))
}

func Test_popInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt64), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt64[0], element)
	assert.Equal(t, utils.ArrayInt64[1:], output.([]int64))
}

func Test_popInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayInt64Ptr[0], element)
	assert.Equal(t, utils.ArrayInt64Ptr[1:], output.([]*int64))
}

func Test_popUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint[0], element)
	assert.Equal(t, utils.ArrayUint[1:], output.([]uint))
}

func Test_popUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUintPtr[0], element)
	assert.Equal(t, utils.ArrayUintPtr[1:], output.([]*uint))
}

func Test_popUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint8), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint8[0], element)
	assert.Equal(t, utils.ArrayUint8[1:], output.([]uint8))
}

func Test_popUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint8Ptr[0], element)
	assert.Equal(t, utils.ArrayUint8Ptr[1:], output.([]*uint8))
}

func Test_popUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint16), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint16[0], element)
	assert.Equal(t, utils.ArrayUint16[1:], output.([]uint16))
}

func Test_popUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint16Ptr[0], element)
	assert.Equal(t, utils.ArrayUint16Ptr[1:], output.([]*uint16))
}

func Test_popUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint32), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint32[0], element)
	assert.Equal(t, utils.ArrayUint32[1:], output.([]uint32))
}

func Test_popUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint32Ptr[0], element)
	assert.Equal(t, utils.ArrayUint32Ptr[1:], output.([]*uint32))
}

func Test_popUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint64), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint64[0], element)
	assert.Equal(t, utils.ArrayUint64[1:], output.([]uint64))
}

func Test_popUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayUint64Ptr[0], element)
	assert.Equal(t, utils.ArrayUint64Ptr[1:], output.([]*uint64))
}

func Test_popFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayFloat32[0], element)
	assert.Equal(t, utils.ArrayFloat32[1:], output.([]float32))
}

func Test_popFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayFloat32Ptr[0], element)
	assert.Equal(t, utils.ArrayFloat32Ptr[1:], output.([]*float32))
}

func Test_popFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayFloat64[0], element)
	assert.Equal(t, utils.ArrayFloat64[1:], output.([]float64))
}

func Test_popFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	found, element, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, utils.ArrayFloat64Ptr[0], element)
	assert.Equal(t, utils.ArrayFloat64Ptr[1:], output.([]*float64))
}
