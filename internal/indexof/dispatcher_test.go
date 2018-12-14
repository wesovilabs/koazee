package indexof

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_indexOfString(t *testing.T) {
	searched := utils.ArrayString[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfStringPtr(t *testing.T) {
	searched := utils.ArrayStringPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfStringNotFound(t *testing.T) {
	searched := utils.ValueString
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfStringPtrNotFound(t *testing.T) {
	searched := utils.ValueStringPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfBool(t *testing.T) {
	searched := utils.ArrayBool[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayBool), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfBoolPtr(t *testing.T) {
	searched := utils.ArrayBoolPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfBoolNotFound(t *testing.T) {
	searched := false
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf([]bool{true, true}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfBoolPtrNotFound(t *testing.T) {
	searchedValue := false
	searched := &searchedValue
	trueValue := true
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf([]*bool{&trueValue, &trueValue}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfInt(t *testing.T) {
	searched := utils.ArrayInt[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfIntPtr(t *testing.T) {
	searched := utils.ArrayIntPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfIntNotFound(t *testing.T) {
	searched := utils.ValueInt
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfIntPtrNotFound(t *testing.T) {
	searched := utils.ValueIntPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfInt8(t *testing.T) {
	searched := utils.ArrayInt8[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfInt8Ptr(t *testing.T) {
	searched := utils.ArrayInt8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfInt8NotFound(t *testing.T) {
	searched := utils.ValueInt8
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfInt8PtrNotFound(t *testing.T) {
	searched := utils.ValueInt8Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfInt16(t *testing.T) {
	searched := utils.ArrayInt16[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfInt16Ptr(t *testing.T) {
	searched := utils.ArrayInt16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfInt16NotFound(t *testing.T) {
	searched := utils.ValueInt16
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfInt16PtrNotFound(t *testing.T) {
	searched := utils.ValueInt16Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfInt32(t *testing.T) {
	searched := utils.ArrayInt32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfInt32Ptr(t *testing.T) {
	searched := utils.ArrayInt32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfInt32NotFound(t *testing.T) {
	searched := utils.ValueInt32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfInt32PtrNotFound(t *testing.T) {
	searched := utils.ValueInt32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfInt64(t *testing.T) {
	searched := utils.ArrayInt64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfInt64Ptr(t *testing.T) {
	searched := utils.ArrayInt64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfInt64NotFound(t *testing.T) {
	searched := utils.ValueInt64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfInt64PtrNotFound(t *testing.T) {
	searched := utils.ValueInt64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfUint(t *testing.T) {
	searched := utils.ArrayUint[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfUintPtr(t *testing.T) {
	searched := utils.ArrayUintPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfUintNotFound(t *testing.T) {
	searched := utils.ValueUint
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfUintPtrNotFound(t *testing.T) {
	searched := utils.ValueUintPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfUint8(t *testing.T) {
	searched := utils.ArrayUint8[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfUint8Ptr(t *testing.T) {
	searched := utils.ArrayUint8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfUint8NotFound(t *testing.T) {
	searched := utils.ValueUint8
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfUint8PtrNotFound(t *testing.T) {
	searched := utils.ValueUint8Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfUint16(t *testing.T) {
	searched := utils.ArrayUint16[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfUint16Ptr(t *testing.T) {
	searched := utils.ArrayUint16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfUint16NotFound(t *testing.T) {
	searched := utils.ValueUint16
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfUint16PtrNotFound(t *testing.T) {
	searched := utils.ValueUint16Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfUint32(t *testing.T) {
	searched := utils.ArrayUint32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfUint32Ptr(t *testing.T) {
	searched := utils.ArrayUint32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfUint32NotFound(t *testing.T) {
	searched := utils.ValueUint32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfUint32PtrNotFound(t *testing.T) {
	searched := utils.ValueUint32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfUint64(t *testing.T) {
	searched := utils.ArrayUint64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfUint64Ptr(t *testing.T) {
	searched := utils.ArrayUint64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfUint64NotFound(t *testing.T) {
	searched := utils.ValueUint64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfUint64PtrNotFound(t *testing.T) {
	searched := utils.ValueUint64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfFloat32(t *testing.T) {
	searched := utils.ArrayFloat32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfFloat32Ptr(t *testing.T) {
	searched := utils.ArrayFloat32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfFloat32NotFound(t *testing.T) {
	searched := utils.ValueFloat32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfFloat32PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_indexOfFloat64(t *testing.T) {
	searched := utils.ArrayFloat64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_indexOfFloat64Ptr(t *testing.T) {
	searched := utils.ArrayFloat64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_indexOfFloat64NotFound(t *testing.T) {
	searched := utils.ValueFloat64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_indexOfFloat64PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
