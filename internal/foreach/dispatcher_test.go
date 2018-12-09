package foreach

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func Test_forEachString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	fn := func(_ string) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayString), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	fn := func(_ *string) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayStringPtr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	fn := func(_ bool) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayBool), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	fn := func(_ *bool) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	fn := func(_ int) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	fn := func(_ *int) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayIntPtr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	fn := func(_ int8) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt8), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	fn := func(_ *int8) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	fn := func(_ int16) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt16), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	fn := func(_ *int16) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	fn := func(_ int32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt32), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	fn := func(_ *int32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	fn := func(_ int64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt64), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	fn := func(_ *int64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	fn := func(_ uint) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	fn := func(_ *uint) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUintPtr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	fn := func(_ uint8) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint8), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	fn := func(_ *uint8) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	fn := func(_ uint16) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint16), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	fn := func(_ *uint16) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	fn := func(_ uint32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint32), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	fn := func(_ *uint32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	fn := func(_ uint64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint64), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	fn := func(_ *uint64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	fn := func(_ float32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat32), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	fn := func(_ *float32) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), typeElement, info)
	assert.True(t, found)
}

func Test_forEachFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	fn := func(_ float64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat64), typeElement, info)
	assert.True(t, found)
}

func Test_forEachPtrFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	fn := func(_ *float64) {}
	info := &forEachInfo{fnValue: reflect.ValueOf(fn)}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), typeElement, info)
	assert.True(t, found)
}
