package filter

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func Test_filterString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayString[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element string) bool { return element != utils.ArrayString[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayString), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrString(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayStringPtr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *string) bool { return *element != *utils.ArrayStringPtr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayStringPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBool[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element bool) bool { return element != utils.ArrayBool[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayBool), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 2)
}

func Test_filterPtrBool(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayBoolPtr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *bool) bool { return *element != *utils.ArrayBoolPtr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayBoolPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 2)
}

func Test_filterInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element int) bool { return element != utils.ArrayInt[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrInt(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayIntPtr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *int) bool { return *element != *utils.ArrayIntPtr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayIntPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element int8) bool { return element != utils.ArrayInt8[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrInt8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt8Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *int8) bool { return *element != *utils.ArrayInt8Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element int16) bool { return element != utils.ArrayInt16[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrInt16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt16Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *int16) bool { return *element != *utils.ArrayInt16Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element int32) bool { return element != utils.ArrayInt32[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrInt32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt32Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *int32) bool { return *element != *utils.ArrayInt32Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element int64) bool { return element != utils.ArrayInt64[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrInt64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayInt64Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *int64) bool { return *element != *utils.ArrayInt64Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayInt64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element uint) bool { return element != utils.ArrayUint[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrUint(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUintPtr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *uint) bool { return *element != *utils.ArrayUintPtr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUintPtr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element uint8) bool { return element != utils.ArrayUint8[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrUint8(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint8Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *uint8) bool { return *element != *utils.ArrayUint8Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint8Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element uint16) bool { return element != utils.ArrayUint16[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrUint16(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint16Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *uint16) bool { return *element != *utils.ArrayUint16Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint16Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element uint32) bool { return element != utils.ArrayUint32[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrUint32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint32Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *uint32) bool { return *element != *utils.ArrayUint32Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element uint64) bool { return element != utils.ArrayUint64[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrUint64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayUint64Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *uint64) bool { return *element != *utils.ArrayUint64Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayUint64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element float32) bool { return element != utils.ArrayFloat32[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrFloat32(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat32Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *float32) bool { return *element != *utils.ArrayFloat32Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat32Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element float64) bool { return element != utils.ArrayFloat64[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}

func Test_filterPtrFloat64(t *testing.T) {
	typeElement := reflect.TypeOf(utils.ArrayFloat64Ptr[1])
	info := &filterInfo{fnIn1Type: typeElement}
	fn := func(element *float64) bool { return *element != *utils.ArrayFloat64Ptr[1] }
	found, output := dispatch(reflect.ValueOf(utils.ArrayFloat64Ptr), fn, info)
	assert.True(t, found)
	assert.Equal(t, output.Len(), 4)
}
