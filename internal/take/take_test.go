package take

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestTake_Run(t *testing.T) {
	sliceOfIntPrt(t)
	emptySlice(t)
	invalidFirstIndex(t)
	invalidLastIndex(t)
	invalidIndexex(t)
}

func emptySlice(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]string{}),
		Len:        0,
		FirstIndex: 1,
		LastIndex:  3,
	}
	newStream, err := op.Run()
	assert.Equal(t, reflect.ValueOf(nil), newStream)
	assert.NotNil(t, err)
	assert.Equal(t, "[take:err.items-nil] It can not be taken an element "+
		"from an empty Stream", err.Error())
}
func invalidFirstIndex(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]string{"a", "b", "c", "d"}),
		Len:        4,
		FirstIndex: -1,
		LastIndex:  3,
	}
	newStream, err := op.Run()
	assert.Equal(t, reflect.ValueOf(nil), newStream)
	assert.NotNil(t, err)
	assert.Equal(t, "[take:err.invalid-index] The length of this Stream is 4, so the indexes must "+
		"be between 0 and 3, being firstIndex lower thant lastIndex", err.Error())
}

func invalidLastIndex(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]string{"a", "b", "c", "d"}),
		Len:        4,
		FirstIndex: 1,
		LastIndex:  5,
	}
	newStream, err := op.Run()
	assert.Equal(t, reflect.ValueOf(nil), newStream)
	assert.NotNil(t, err)
	assert.Equal(t, "[take:err.invalid-index] The length of this Stream is 4, so the indexes must "+
		"be between 0 and 3, being firstIndex lower thant lastIndex", err.Error())
}

func invalidIndexex(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]string{"a", "b", "c", "d"}),
		Len:        4,
		FirstIndex: 1,
		LastIndex:  0,
	}
	newStream, err := op.Run()
	assert.Equal(t, reflect.ValueOf(nil), newStream)
	assert.NotNil(t, err)
	assert.Equal(t, "[take:err.invalid-index] The length of this Stream is 4, so the indexes must "+
		"be between 0 and 3, being firstIndex lower thant lastIndex", err.Error())
}

func sliceOfIntPrt(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4), utils.IntPtr(22), utils.IntPtr(1), utils.IntPtr(9), utils.IntPtr(89), utils.IntPtr(40)}),
		Len:        6,
		FirstIndex: 1,
		LastIndex:  3,
	}
	newStream, err := op.Run()
	assert.Nil(t, err)
	out := newStream.Interface().([]*int)
	assert.Len(t, out, 3)
	assert.Equal(t, []*int{utils.IntPtr(22), utils.IntPtr(1), utils.IntPtr(9)}, out)
}
