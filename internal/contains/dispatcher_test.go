package contains

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_containsString(t *testing.T) {
	searched := utils.ArrayString[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsStringPtr(t *testing.T) {
	searched := utils.ArrayStringPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsStringNotFound(t *testing.T) {
	searched := utils.ValueString
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayString), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsStringPtrNotFound(t *testing.T) {
	searched := utils.ValueStringPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayStringPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsBool(t *testing.T) {
	searched := utils.ArrayBool[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayBool), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsBoolPtr(t *testing.T) {
	searched := utils.ArrayBoolPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsBoolNotFound(t *testing.T) {
	searched := false
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf([]bool{true, true}), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsBoolPtrNotFound(t *testing.T) {
	searchedValue := false
	searched := &searchedValue
	trueValue := true
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf([]*bool{&trueValue, &trueValue}), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt(t *testing.T) {
	searched := utils.ArrayInt[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsIntPtr(t *testing.T) {
	searched := utils.ArrayIntPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsIntNotFound(t *testing.T) {
	searched := utils.ValueInt
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsIntPtrNotFound(t *testing.T) {
	searched := utils.ValueIntPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayIntPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt8(t *testing.T) {
	searched := utils.ArrayInt8[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt8Ptr(t *testing.T) {
	searched := utils.ArrayInt8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt8NotFound(t *testing.T) {
	searched := utils.ValueInt8
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt8), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt8PtrNotFound(t *testing.T) {
	searched := utils.ValueInt8Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt16(t *testing.T) {
	searched := utils.ArrayInt16[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt16Ptr(t *testing.T) {
	searched := utils.ArrayInt16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt16NotFound(t *testing.T) {
	searched := utils.ValueInt16
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt16), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt16PtrNotFound(t *testing.T) {
	searched := utils.ValueInt16Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt32(t *testing.T) {
	searched := utils.ArrayInt32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt32Ptr(t *testing.T) {
	searched := utils.ArrayInt32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt32NotFound(t *testing.T) {
	searched := utils.ValueInt32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt32PtrNotFound(t *testing.T) {
	searched := utils.ValueInt32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt64(t *testing.T) {
	searched := utils.ArrayInt64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt64Ptr(t *testing.T) {
	searched := utils.ArrayInt64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt64NotFound(t *testing.T) {
	searched := utils.ValueInt64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt64PtrNotFound(t *testing.T) {
	searched := utils.ValueInt64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint(t *testing.T) {
	searched := utils.ArrayUint[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUintPtr(t *testing.T) {
	searched := utils.ArrayUintPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUintNotFound(t *testing.T) {
	searched := utils.ValueUint
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUintPtrNotFound(t *testing.T) {
	searched := utils.ValueUintPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUintPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint8(t *testing.T) {
	searched := utils.ArrayUint8[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint8Ptr(t *testing.T) {
	searched := utils.ArrayUint8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint8NotFound(t *testing.T) {
	searched := utils.ValueUint8
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint8), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint8PtrNotFound(t *testing.T) {
	searched := utils.ValueUint8Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint16(t *testing.T) {
	searched := utils.ArrayUint16[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint16Ptr(t *testing.T) {
	searched := utils.ArrayUint16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint16NotFound(t *testing.T) {
	searched := utils.ValueUint16
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint16), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint16PtrNotFound(t *testing.T) {
	searched := utils.ValueUint16Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint32(t *testing.T) {
	searched := utils.ArrayUint32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint32Ptr(t *testing.T) {
	searched := utils.ArrayUint32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint32NotFound(t *testing.T) {
	searched := utils.ValueUint32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint32PtrNotFound(t *testing.T) {
	searched := utils.ValueUint32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint64(t *testing.T) {
	searched := utils.ArrayUint64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint64Ptr(t *testing.T) {
	searched := utils.ArrayUint64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint64NotFound(t *testing.T) {
	searched := utils.ValueUint64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint64PtrNotFound(t *testing.T) {
	searched := utils.ValueUint64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsFloat32(t *testing.T) {
	searched := utils.ArrayFloat32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsFloat32Ptr(t *testing.T) {
	searched := utils.ArrayFloat32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsFloat32NotFound(t *testing.T) {
	searched := utils.ValueFloat32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsFloat32PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsFloat64(t *testing.T) {
	searched := utils.ArrayFloat64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsFloat64Ptr(t *testing.T) {
	searched := utils.ArrayFloat64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsFloat64NotFound(t *testing.T) {
	searched := utils.ValueFloat64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsFloat64PtrNotFound(t *testing.T) {
	searched := utils.ValueFloat64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
