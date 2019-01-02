package indexesof

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_indexesOfString(t *testing.T) {
	searched := utils.ArrayString[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfStringPtr(t *testing.T) {
	searched := utils.ArrayStringPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfStringNotFound(t *testing.T) {
	searched := utils.ValueString
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfStringPtrNotFound(t *testing.T) {
	searched := utils.ValueStringPtr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfBool(t *testing.T) {
	searched := utils.ArrayBool[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1, 3, 4}, output)
}

func Test_indexesOfBoolPtr(t *testing.T) {
	searched := utils.ArrayBoolPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1, 3, 4}, output)
}
func Test_indexesOfBoolNotFound(t *testing.T) {
	searched := false
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf([]bool{true, true}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfBoolPtrNotFound(t *testing.T) {
	searchedValue := false
	searched := &searchedValue
	trueValue := true
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf([]*bool{&trueValue, &trueValue}), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfInt(t *testing.T) {
	searched := utils.ArrayInt[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfIntPtr(t *testing.T) {
	searched := utils.ArrayIntPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfIntNotFound(t *testing.T) {
	searched := utils.ValueInt
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfIntPtrNotFound(t *testing.T) {
	searched := utils.ValueIntPtr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfInt8(t *testing.T) {
	searched := utils.ArrayInt8[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfInt8Ptr(t *testing.T) {
	searched := utils.ArrayInt8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfInt8NotFound(t *testing.T) {
	searched := utils.ValueInt8
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfInt8PtrNotFound(t *testing.T) {
	searched := utils.ValueInt8Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfInt16(t *testing.T) {
	searched := utils.ArrayInt16[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfInt16Ptr(t *testing.T) {
	searched := utils.ArrayInt16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfInt16NotFound(t *testing.T) {
	searched := utils.ValueInt16
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfInt16PtrNotFound(t *testing.T) {
	searched := utils.ValueInt16Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfInt32(t *testing.T) {
	searched := utils.ArrayInt32[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfInt32Ptr(t *testing.T) {
	searched := utils.ArrayInt32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfInt32NotFound(t *testing.T) {
	searched := utils.ValueInt32
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfInt32PtrNotFound(t *testing.T) {
	searched := utils.ValueInt32Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfInt64(t *testing.T) {
	searched := utils.ArrayInt64[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfInt64Ptr(t *testing.T) {
	searched := utils.ArrayInt64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfInt64NotFound(t *testing.T) {
	searched := utils.ValueInt64
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfInt64PtrNotFound(t *testing.T) {
	searched := utils.ValueInt64Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfUint(t *testing.T) {
	searched := utils.ArrayUint[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfUintPtr(t *testing.T) {
	searched := utils.ArrayUintPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfUintNotFound(t *testing.T) {
	searched := utils.ValueUint
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfUintPtrNotFound(t *testing.T) {
	searched := utils.ValueUintPtr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfUint8(t *testing.T) {
	searched := utils.ArrayUint8[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfUint8Ptr(t *testing.T) {
	searched := utils.ArrayUint8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfUint8NotFound(t *testing.T) {
	searched := utils.ValueUint8
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfUint8PtrNotFound(t *testing.T) {
	searched := utils.ValueUint8Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfUint16(t *testing.T) {
	searched := utils.ArrayUint16[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfUint16Ptr(t *testing.T) {
	searched := utils.ArrayUint16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfUint16NotFound(t *testing.T) {
	searched := utils.ValueUint16
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfUint16PtrNotFound(t *testing.T) {
	searched := utils.ValueUint16Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfUint32(t *testing.T) {
	searched := utils.ArrayUint32[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfUint32Ptr(t *testing.T) {
	searched := utils.ArrayUint32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfUint32NotFound(t *testing.T) {
	searched := utils.ValueUint32
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfUint32PtrNotFound(t *testing.T) {
	searched := utils.ValueUint32Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfUint64(t *testing.T) {
	searched := utils.ArrayUint64[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfUint64Ptr(t *testing.T) {
	searched := utils.ArrayUint64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfUint64NotFound(t *testing.T) {
	searched := utils.ValueUint64
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfUint64PtrNotFound(t *testing.T) {
	searched := utils.ValueUint64Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfFloat32(t *testing.T) {
	searched := utils.ArrayFloat32[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfFloat32Ptr(t *testing.T) {
	searched := utils.ArrayFloat32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfFloat32NotFound(t *testing.T) {
	searched := utils.ValueFloat32
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfFloat32PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat32Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
func Test_indexesOfFloat64(t *testing.T) {
	searched := utils.ArrayFloat64[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}

func Test_indexesOfFloat64Ptr(t *testing.T) {
	searched := utils.ArrayFloat64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{1}, output)
}
func Test_indexesOfFloat64NotFound(t *testing.T) {
	searched := utils.ValueFloat64
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}

func Test_indexesOfFloat64PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat64Ptr
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.Equal(t, []int{}, output)
}
