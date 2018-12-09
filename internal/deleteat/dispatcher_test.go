package deleteat

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_deleteAtString(t *testing.T) {
	searched := utils.Array_string[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_string), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtStringPtr(t *testing.T) {
	searched := utils.Array_stringPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_stringPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtBool(t *testing.T) {
	searched := utils.Array_bool[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_bool), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtBoolPtr(t *testing.T) {
	searched := utils.Array_boolPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_boolPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt(t *testing.T) {
	searched := utils.Array_int[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtIntPtr(t *testing.T) {
	searched := utils.Array_intPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_intPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt8(t *testing.T) {
	searched := utils.Array_int8[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int8), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt8Ptr(t *testing.T) {
	searched := utils.Array_int8Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int8Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt16(t *testing.T) {
	searched := utils.Array_int16[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int16), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt16Ptr(t *testing.T) {
	searched := utils.Array_int16Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int16Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt32(t *testing.T) {
	searched := utils.Array_int32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt32Ptr(t *testing.T) {
	searched := utils.Array_int32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt64(t *testing.T) {
	searched := utils.Array_int64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt64Ptr(t *testing.T) {
	searched := utils.Array_int64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_int64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint(t *testing.T) {
	searched := utils.Array_uint[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUintPtr(t *testing.T) {
	searched := utils.Array_uintPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uintPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint8(t *testing.T) {
	searched := utils.Array_uint8[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint8Ptr(t *testing.T) {
	searched := utils.Array_uint8Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint8Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint16(t *testing.T) {
	searched := utils.Array_uint16[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint16Ptr(t *testing.T) {
	searched := utils.Array_uint16Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint16Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint32(t *testing.T) {
	searched := utils.Array_uint32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint32Ptr(t *testing.T) {
	searched := utils.Array_uint32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint64(t *testing.T) {
	searched := utils.Array_uint64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint64Ptr(t *testing.T) {
	searched := utils.Array_uint64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_uint64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat32(t *testing.T) {
	searched := utils.Array_float32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_float32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat32Ptr(t *testing.T) {
	searched := utils.Array_float32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_float32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat64(t *testing.T) {
	searched := utils.Array_float64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_float64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat64Ptr(t *testing.T) {
	searched := utils.Array_float64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.Array_float64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}
