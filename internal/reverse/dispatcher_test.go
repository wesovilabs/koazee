package reverse

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), typeElement)
	assert.True(t, found)
	for index := range output.([]string) {
		assert.Equal(t, output.([]string)[index], utils.ArrayString[4-index])
	}
}

func Test_reverseStringPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), typeElement)
	assert.True(t, found)
	for index := range output.([]*string) {
		assert.Equal(t, *output.([]*string)[index], *utils.ArrayStringPtr[4-index])
	}
}

func Test_reverseBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), typeElement)
	assert.True(t, found)
	for index := range output.([]bool) {
		assert.Equal(t, output.([]bool)[index], utils.ArrayBool[4-index])
	}
}

func Test_reverseBoolPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), typeElement)
	assert.True(t, found)
	for index := range output.([]*bool) {
		assert.Equal(t, *output.([]*bool)[index], *utils.ArrayBoolPtr[4-index])
	}
}

func Test_reverseInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), typeElement)
	assert.True(t, found)
	for index := range output.([]int) {
		assert.Equal(t, output.([]int)[index], utils.ArrayInt[4-index])
	}
}

func Test_reverseIntPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), typeElement)
	assert.True(t, found)
	for index := range output.([]*int) {
		assert.Equal(t, *output.([]*int)[index], *utils.ArrayIntPtr[4-index])
	}
}

func Test_reverseInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), typeElement)
	assert.True(t, found)
	for index := range output.([]int8) {
		assert.Equal(t, output.([]int8)[index], utils.ArrayInt8[4-index])
	}
}

func Test_reverseInt8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*int8) {
		assert.Equal(t, *output.([]*int8)[index], *utils.ArrayInt8Ptr[4-index])
	}
}

func Test_reverseInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), typeElement)
	assert.True(t, found)
	for index := range output.([]int16) {
		assert.Equal(t, output.([]int16)[index], utils.ArrayInt16[4-index])
	}
}

func Test_reverseInt16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*int16) {
		assert.Equal(t, *output.([]*int16)[index], *utils.ArrayInt16Ptr[4-index])
	}
}

func Test_reverseInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), typeElement)
	assert.True(t, found)
	for index := range output.([]int32) {
		assert.Equal(t, output.([]int32)[index], utils.ArrayInt32[4-index])
	}
}

func Test_reverseInt32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*int32) {
		assert.Equal(t, *output.([]*int32)[index], *utils.ArrayInt32Ptr[4-index])
	}
}

func Test_reverseInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), typeElement)
	assert.True(t, found)
	for index := range output.([]int64) {
		assert.Equal(t, output.([]int64)[index], utils.ArrayInt64[4-index])
	}
}

func Test_reverseInt64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*int64) {
		assert.Equal(t, *output.([]*int64)[index], *utils.ArrayInt64Ptr[4-index])
	}
}

func Test_reverseUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), typeElement)
	assert.True(t, found)
	for index := range output.([]uint) {
		assert.Equal(t, output.([]uint)[index], utils.ArrayUint[4-index])
	}
}

func Test_reverseUintPtr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), typeElement)
	assert.True(t, found)
	for index := range output.([]*uint) {
		assert.Equal(t, *output.([]*uint)[index], *utils.ArrayUintPtr[4-index])
	}
}

func Test_reverseUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), typeElement)
	assert.True(t, found)
	for index := range output.([]uint8) {
		assert.Equal(t, output.([]uint8)[index], utils.ArrayUint8[4-index])
	}
}

func Test_reverseUint8Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*uint8) {
		assert.Equal(t, *output.([]*uint8)[index], *utils.ArrayUint8Ptr[4-index])
	}
}

func Test_reverseUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), typeElement)
	assert.True(t, found)
	for index := range output.([]uint16) {
		assert.Equal(t, output.([]uint16)[index], utils.ArrayUint16[4-index])
	}
}

func Test_reverseUint16Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*uint16) {
		assert.Equal(t, *output.([]*uint16)[index], *utils.ArrayUint16Ptr[4-index])
	}
}

func Test_reverseUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), typeElement)
	assert.True(t, found)
	for index := range output.([]uint32) {
		assert.Equal(t, output.([]uint32)[index], utils.ArrayUint32[4-index])
	}
}

func Test_reverseUint32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*uint32) {
		assert.Equal(t, *output.([]*uint32)[index], *utils.ArrayUint32Ptr[4-index])
	}
}

func Test_reverseUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), typeElement)
	assert.True(t, found)
	for index := range output.([]uint64) {
		assert.Equal(t, output.([]uint64)[index], utils.ArrayUint64[4-index])
	}
}

func Test_reverseUint64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*uint64) {
		assert.Equal(t, *output.([]*uint64)[index], *utils.ArrayUint64Ptr[4-index])
	}
}

func Test_reverseFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), typeElement)
	assert.True(t, found)
	for index := range output.([]float32) {
		assert.Equal(t, output.([]float32)[index], utils.ArrayFloat32[4-index])
	}
}

func Test_reverseFloat32Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*float32) {
		assert.Equal(t, *output.([]*float32)[index], *utils.ArrayFloat32Ptr[4-index])
	}
}

func Test_reverseFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), typeElement)
	assert.True(t, found)
	for index := range output.([]float64) {
		assert.Equal(t, output.([]float64)[index], utils.ArrayFloat64[4-index])
	}
}

func Test_reverseFloat64Ptr(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), typeElement)
	assert.True(t, found)
	for index := range output.([]*float64) {
		assert.Equal(t, *output.([]*float64)[index], *utils.ArrayFloat64Ptr[4-index])
	}
}
