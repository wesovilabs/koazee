package contains

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_containsString(t *testing.T) {
	searched := utils.Array_string[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_string), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsStringPtr(t *testing.T) {
	searched := utils.Array_stringPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_stringPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsStringNotFound(t *testing.T) {
	searched := utils.Value_string
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_string), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsStringPtrNotFound(t *testing.T) {
	searched := utils.Value_stringPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_stringPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsBool(t *testing.T) {
	searched := utils.Array_bool[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_bool), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsBoolPtr(t *testing.T) {
	searched := utils.Array_boolPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_boolPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsBoolNotFound(t *testing.T) {
	searched := utils.Value_bool
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_bool), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsBoolPtrNotFound(t *testing.T) {
	searched := utils.Value_boolPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_boolPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt(t *testing.T) {
	searched := utils.Array_int[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsIntPtr(t *testing.T) {
	searched := utils.Array_intPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_intPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsIntNotFound(t *testing.T) {
	searched := utils.Value_int
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsIntPtrNotFound(t *testing.T) {
	searched := utils.Value_intPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_intPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt8(t *testing.T) {
	searched := utils.Array_int8[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int8), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt8Ptr(t *testing.T) {
	searched := utils.Array_int8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt8NotFound(t *testing.T) {
	searched := utils.Value_int8
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int8), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt8PtrNotFound(t *testing.T) {
	searched := utils.Value_int8Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt16(t *testing.T) {
	searched := utils.Array_int16[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int16), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt16Ptr(t *testing.T) {
	searched := utils.Array_int16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt16NotFound(t *testing.T) {
	searched := utils.Value_int16
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int16), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt16PtrNotFound(t *testing.T) {
	searched := utils.Value_int16Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt32(t *testing.T) {
	searched := utils.Array_int32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt32Ptr(t *testing.T) {
	searched := utils.Array_int32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt32NotFound(t *testing.T) {
	searched := utils.Value_int32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt32PtrNotFound(t *testing.T) {
	searched := utils.Value_int32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsInt64(t *testing.T) {
	searched := utils.Array_int64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsInt64Ptr(t *testing.T) {
	searched := utils.Array_int64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsInt64NotFound(t *testing.T) {
	searched := utils.Value_int64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsInt64PtrNotFound(t *testing.T) {
	searched := utils.Value_int64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_int64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint(t *testing.T) {
	searched := utils.Array_uint[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUintPtr(t *testing.T) {
	searched := utils.Array_uintPtr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uintPtr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUintNotFound(t *testing.T) {
	searched := utils.Value_uint
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUintPtrNotFound(t *testing.T) {
	searched := utils.Value_uintPtr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uintPtr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint8(t *testing.T) {
	searched := utils.Array_uint8[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint8), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint8Ptr(t *testing.T) {
	searched := utils.Array_uint8Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint8NotFound(t *testing.T) {
	searched := utils.Value_uint8
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint8), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint8PtrNotFound(t *testing.T) {
	searched := utils.Value_uint8Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint8Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint16(t *testing.T) {
	searched := utils.Array_uint16[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint16), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint16Ptr(t *testing.T) {
	searched := utils.Array_uint16Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint16NotFound(t *testing.T) {
	searched := utils.Value_uint16
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint16), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint16PtrNotFound(t *testing.T) {
	searched := utils.Value_uint16Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint16Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint32(t *testing.T) {
	searched := utils.Array_uint32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint32Ptr(t *testing.T) {
	searched := utils.Array_uint32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint32NotFound(t *testing.T) {
	searched := utils.Value_uint32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint32PtrNotFound(t *testing.T) {
	searched := utils.Value_uint32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsUint64(t *testing.T) {
	searched := utils.Array_uint64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsUint64Ptr(t *testing.T) {
	searched := utils.Array_uint64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsUint64NotFound(t *testing.T) {
	searched := utils.Value_uint64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsUint64PtrNotFound(t *testing.T) {
	searched := utils.Value_uint64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_uint64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsFloat32(t *testing.T) {
	searched := utils.Array_float32[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float32), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsFloat32Ptr(t *testing.T) {
	searched := utils.Array_float32Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsFloat32NotFound(t *testing.T) {
	searched := utils.Value_float32
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float32), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsFloat32PtrNotFound(t *testing.T) {
	searched := utils.Value_float32Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float32Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
func Test_containsFloat64(t *testing.T) {
	searched := utils.Array_float64[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float64), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}

func Test_containsFloat64Ptr(t *testing.T) {
	searched := utils.Array_float64Ptr[1]
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.True(t, contained)
}
func Test_containsFloat64NotFound(t *testing.T) {
	searched := utils.Value_float64
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float64), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}

func Test_containsFloat64PtrNotFound(t *testing.T) {
	searched := utils.Value_float64Ptr
	typeElement := reflect.TypeOf(searched)
	found, contained := dispatch(reflect.ValueOf(utils.Array_float64Ptr), searched, typeElement)
	assert.True(t, found)
	assert.False(t, contained)
}
