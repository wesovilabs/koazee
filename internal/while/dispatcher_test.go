package while

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_whileString(t *testing.T) {
	sfn := func(_ string) bool { return true }
	dfn := func(v string) {}
	inputType := reflect.TypeOf(utils.ArrayString[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayString), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrString(t *testing.T) {
	sfn := func(_ *string) bool { return true }
	dfn := func(v *string) {}
	inputType := reflect.TypeOf(utils.ArrayStringPtr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayStringPtr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileBool(t *testing.T) {
	sfn := func(_ bool) bool { return true }
	dfn := func(v bool) {}
	inputType := reflect.TypeOf(utils.ArrayBool[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayBool), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrBool(t *testing.T) {
	sfn := func(_ *bool) bool { return true }
	dfn := func(v *bool) {}
	inputType := reflect.TypeOf(utils.ArrayBoolPtr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileInt(t *testing.T) {
	sfn := func(_ int) bool { return true }
	dfn := func(v int) {}
	inputType := reflect.TypeOf(utils.ArrayInt[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrInt(t *testing.T) {
	sfn := func(_ *int) bool { return true }
	dfn := func(v *int) {}
	inputType := reflect.TypeOf(utils.ArrayIntPtr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayIntPtr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileInt8(t *testing.T) {
	sfn := func(_ int8) bool { return true }
	dfn := func(v int8) {}
	inputType := reflect.TypeOf(utils.ArrayInt8[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt8), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrInt8(t *testing.T) {
	sfn := func(_ *int8) bool { return true }
	dfn := func(v *int8) {}
	inputType := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileInt16(t *testing.T) {
	sfn := func(_ int16) bool { return true }
	dfn := func(v int16) {}
	inputType := reflect.TypeOf(utils.ArrayInt16[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt16), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrInt16(t *testing.T) {
	sfn := func(_ *int16) bool { return true }
	dfn := func(v *int16) {}
	inputType := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileInt32(t *testing.T) {
	sfn := func(_ int32) bool { return true }
	dfn := func(v int32) {}
	inputType := reflect.TypeOf(utils.ArrayInt32[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt32), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrInt32(t *testing.T) {
	sfn := func(_ *int32) bool { return true }
	dfn := func(v *int32) {}
	inputType := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileInt64(t *testing.T) {
	sfn := func(_ int64) bool { return true }
	dfn := func(v int64) {}
	inputType := reflect.TypeOf(utils.ArrayInt64[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt64), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrInt64(t *testing.T) {
	sfn := func(_ *int64) bool { return true }
	dfn := func(v *int64) {}
	inputType := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileUint(t *testing.T) {
	sfn := func(_ uint) bool { return true }
	dfn := func(v uint) {}
	inputType := reflect.TypeOf(utils.ArrayUint[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrUint(t *testing.T) {
	sfn := func(_ *uint) bool { return true }
	dfn := func(v *uint) {}
	inputType := reflect.TypeOf(utils.ArrayUintPtr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUintPtr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileUint8(t *testing.T) {
	sfn := func(_ uint8) bool { return true }
	dfn := func(v uint8) {}
	inputType := reflect.TypeOf(utils.ArrayUint8[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint8), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrUint8(t *testing.T) {
	sfn := func(_ *uint8) bool { return true }
	dfn := func(v *uint8) {}
	inputType := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileUint16(t *testing.T) {
	sfn := func(_ uint16) bool { return true }
	dfn := func(v uint16) {}
	inputType := reflect.TypeOf(utils.ArrayUint16[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint16), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrUint16(t *testing.T) {
	sfn := func(_ *uint16) bool { return true }
	dfn := func(v *uint16) {}
	inputType := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileUint32(t *testing.T) {
	sfn := func(_ uint32) bool { return true }
	dfn := func(v uint32) {}
	inputType := reflect.TypeOf(utils.ArrayUint32[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint32), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrUint32(t *testing.T) {
	sfn := func(_ *uint32) bool { return true }
	dfn := func(v *uint32) {}
	inputType := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileUint64(t *testing.T) {
	sfn := func(_ uint64) bool { return true }
	dfn := func(v uint64) {}
	inputType := reflect.TypeOf(utils.ArrayUint64[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint64), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrUint64(t *testing.T) {
	sfn := func(_ *uint64) bool { return true }
	dfn := func(v *uint64) {}
	inputType := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileFloat32(t *testing.T) {
	sfn := func(_ float32) bool { return true }
	dfn := func(v float32) {}
	inputType := reflect.TypeOf(utils.ArrayFloat32[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat32), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrFloat32(t *testing.T) {
	sfn := func(_ *float32) bool { return true }
	dfn := func(v *float32) {}
	inputType := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whileFloat64(t *testing.T) {
	sfn := func(_ float64) bool { return true }
	dfn := func(v float64) {}
	inputType := reflect.TypeOf(utils.ArrayFloat64[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat64), sfn, dfn, info)
	assert.True(t, found)
}

func Test_whilePtrFloat64(t *testing.T) {
	sfn := func(_ *float64) bool { return true }
	dfn := func(v *float64) {}
	inputType := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	info := &whileInfo{fnValue: reflect.ValueOf(sfn), fnInputType: inputType}
	found := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), sfn, dfn, info)
	assert.True(t, found)
}
