package concat

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_concatString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), reflect.ValueOf(utils.ArrayString), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), reflect.ValueOf(utils.ArrayStringPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), reflect.ValueOf(utils.ArrayBool), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), reflect.ValueOf(utils.ArrayBoolPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), reflect.ValueOf(utils.ArrayInt), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), reflect.ValueOf(utils.ArrayIntPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), reflect.ValueOf(utils.ArrayInt8), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), reflect.ValueOf(utils.ArrayInt8Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), reflect.ValueOf(utils.ArrayInt16), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), reflect.ValueOf(utils.ArrayInt16Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), reflect.ValueOf(utils.ArrayInt32), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), reflect.ValueOf(utils.ArrayInt32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), reflect.ValueOf(utils.ArrayInt64), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), reflect.ValueOf(utils.ArrayInt64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), reflect.ValueOf(utils.ArrayUint), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), reflect.ValueOf(utils.ArrayUintPtr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), reflect.ValueOf(utils.ArrayUint8), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), reflect.ValueOf(utils.ArrayUint8Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), reflect.ValueOf(utils.ArrayUint16), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), reflect.ValueOf(utils.ArrayUint16Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), reflect.ValueOf(utils.ArrayUint32), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), reflect.ValueOf(utils.ArrayUint32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), reflect.ValueOf(utils.ArrayUint64), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), reflect.ValueOf(utils.ArrayUint64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), reflect.ValueOf(utils.ArrayFloat32), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), reflect.ValueOf(utils.ArrayFloat32Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), reflect.ValueOf(utils.ArrayFloat64), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}

func Test_concatPtrFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), reflect.ValueOf(utils.ArrayFloat64Ptr), typeElement)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 10)
}
