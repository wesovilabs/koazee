package lastindexof

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_lastIndexOfString(t *testing.T) {
	searched := utils.ArrayString[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfStringPtr(t *testing.T) {
	searched := utils.ArrayStringPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfStringNotFound(t *testing.T) {
	searched := utils.ValueString
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfStringPtrNotFound(t *testing.T) {
	searched := utils.ValueStringPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfBool(t *testing.T) {
	searched := utils.ArrayBool[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayBool), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 4, index)
}

func Test_lastIndexOfBoolPtr(t *testing.T) {
	searched := utils.ArrayBoolPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 4, index)
}
func Test_lastIndexOfBoolNotFound(t *testing.T) {
	searched := false
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf([]bool{true, true}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfBoolPtrNotFound(t *testing.T) {
	searchedValue := false
	searched := &searchedValue
	trueValue := true
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf([]*bool{&trueValue, &trueValue}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfInt(t *testing.T) {
	searched := utils.ArrayInt[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfIntPtr(t *testing.T) {
	searched := utils.ArrayIntPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfIntNotFound(t *testing.T) {
	searched := utils.ValueInt
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfIntPtrNotFound(t *testing.T) {
	searched := utils.ValueIntPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfInt8(t *testing.T) {
	searched := utils.ArrayInt8[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfInt8Ptr(t *testing.T) {
	searched := utils.ArrayInt8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfInt8NotFound(t *testing.T) {
	searched := utils.ValueInt8
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfInt8PtrNotFound(t *testing.T) {
	searched := utils.ValueInt8Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfInt16(t *testing.T) {
	searched := utils.ArrayInt16[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfInt16Ptr(t *testing.T) {
	searched := utils.ArrayInt16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfInt16NotFound(t *testing.T) {
	searched := utils.ValueInt16
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfInt16PtrNotFound(t *testing.T) {
	searched := utils.ValueInt16Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfInt32(t *testing.T) {
	searched := utils.ArrayInt32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfInt32Ptr(t *testing.T) {
	searched := utils.ArrayInt32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfInt32NotFound(t *testing.T) {
	searched := utils.ValueInt32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfInt32PtrNotFound(t *testing.T) {
	searched := utils.ValueInt32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfInt64(t *testing.T) {
	searched := utils.ArrayInt64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfInt64Ptr(t *testing.T) {
	searched := utils.ArrayInt64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfInt64NotFound(t *testing.T) {
	searched := utils.ValueInt64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfInt64PtrNotFound(t *testing.T) {
	searched := utils.ValueInt64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfUint(t *testing.T) {
	searched := utils.ArrayUint[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfUintPtr(t *testing.T) {
	searched := utils.ArrayUintPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfUintNotFound(t *testing.T) {
	searched := utils.ValueUint
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfUintPtrNotFound(t *testing.T) {
	searched := utils.ValueUintPtr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfUint8(t *testing.T) {
	searched := utils.ArrayUint8[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfUint8Ptr(t *testing.T) {
	searched := utils.ArrayUint8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfUint8NotFound(t *testing.T) {
	searched := utils.ValueUint8
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfUint8PtrNotFound(t *testing.T) {
	searched := utils.ValueUint8Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfUint16(t *testing.T) {
	searched := utils.ArrayUint16[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfUint16Ptr(t *testing.T) {
	searched := utils.ArrayUint16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfUint16NotFound(t *testing.T) {
	searched := utils.ValueUint16
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfUint16PtrNotFound(t *testing.T) {
	searched := utils.ValueUint16Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfUint32(t *testing.T) {
	searched := utils.ArrayUint32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfUint32Ptr(t *testing.T) {
	searched := utils.ArrayUint32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfUint32NotFound(t *testing.T) {
	searched := utils.ValueUint32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfUint32PtrNotFound(t *testing.T) {
	searched := utils.ValueUint32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfUint64(t *testing.T) {
	searched := utils.ArrayUint64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfUint64Ptr(t *testing.T) {
	searched := utils.ArrayUint64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfUint64NotFound(t *testing.T) {
	searched := utils.ValueUint64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfUint64PtrNotFound(t *testing.T) {
	searched := utils.ValueUint64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfFloat32(t *testing.T) {
	searched := utils.ArrayFloat32[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfFloat32Ptr(t *testing.T) {
	searched := utils.ArrayFloat32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfFloat32NotFound(t *testing.T) {
	searched := utils.ValueFloat32
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfFloat32PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat32Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
func Test_lastIndexOfFloat64(t *testing.T) {
	searched := utils.ArrayFloat64[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}

func Test_lastIndexOfFloat64Ptr(t *testing.T) {
	searched := utils.ArrayFloat64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, 1, index)
}
func Test_lastIndexOfFloat64NotFound(t *testing.T) {
	searched := utils.ValueFloat64
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}

func Test_lastIndexOfFloat64PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat64Ptr
	typeElement := reflect.TypeOf(searched)
	found, index := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, InvalidIndex, index)
}
