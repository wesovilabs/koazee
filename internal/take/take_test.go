package take

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestTake_Run(t *testing.T) {
	op := Take{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4), utils.IntPtr(22), utils.IntPtr(1), utils.IntPtr(9), utils.IntPtr(89), utils.IntPtr(40)}),
		Len:        6,
		FirstIndex: 1,
		LastIndex:  3,
	}
	newStream, err := op.Run()
	assert.Nil(t, err)
	out := newStream.Interface().([]*int)
	assert.Len(t, out, 2)
}
