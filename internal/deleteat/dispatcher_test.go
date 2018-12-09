package deleteat

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_deleteAtString(t *testing.T) {
	searched := utils.ArrayString[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtStringPtr(t *testing.T) {
	searched := utils.ArrayStringPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtBool(t *testing.T) {
	searched := utils.ArrayBool[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtBoolPtr(t *testing.T) {
	searched := utils.ArrayBoolPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt(t *testing.T) {
	searched := utils.ArrayInt[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtIntPtr(t *testing.T) {
	searched := utils.ArrayIntPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt8(t *testing.T) {
	searched := utils.ArrayInt8[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt8Ptr(t *testing.T) {
	searched := utils.ArrayInt8Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt16(t *testing.T) {
	searched := utils.ArrayInt16[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt16Ptr(t *testing.T) {
	searched := utils.ArrayInt16Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt32(t *testing.T) {
	searched := utils.ArrayInt32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt32Ptr(t *testing.T) {
	searched := utils.ArrayInt32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt64(t *testing.T) {
	searched := utils.ArrayInt64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtInt64Ptr(t *testing.T) {
	searched := utils.ArrayInt64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint(t *testing.T) {
	searched := utils.ArrayUint[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUintPtr(t *testing.T) {
	searched := utils.ArrayUintPtr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint8(t *testing.T) {
	searched := utils.ArrayUint8[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint8Ptr(t *testing.T) {
	searched := utils.ArrayUint8Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint16(t *testing.T) {
	searched := utils.ArrayUint16[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint16Ptr(t *testing.T) {
	searched := utils.ArrayUint16Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint32(t *testing.T) {
	searched := utils.ArrayUint32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint32Ptr(t *testing.T) {
	searched := utils.ArrayUint32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint64(t *testing.T) {
	searched := utils.ArrayUint64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtUint64Ptr(t *testing.T) {
	searched := utils.ArrayUint64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat32(t *testing.T) {
	searched := utils.ArrayFloat32[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat32Ptr(t *testing.T) {
	searched := utils.ArrayFloat32Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat64(t *testing.T) {
	searched := utils.ArrayFloat64[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}

func Test_deleteAtFloat64Ptr(t *testing.T) {
	searched := utils.ArrayFloat64Ptr[0]
	typeElement := reflect.TypeOf(searched)
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), &typeElement, 5, 0)
	assert.True(t, found)
	assert.Len(t, output, 4)
}
