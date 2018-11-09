package stream

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_newStream(t *testing.T) {
	stream := new([]int{2, 3})
	assert.Equal(t, reflect.TypeOf(1), stream.elementsType)
	count, _ := stream.Count()
	assert.Equal(t, 2, count)
	assert.Equal(t, []int{2, 3}, stream.items)
}

func TestStream_emptyStream(t *testing.T) {
	stream := emptyStream(reflect.TypeOf("hello"))
	assert.Equal(t, reflect.TypeOf("hello"), stream.elementsType)
	count, _ := stream.Count()
	assert.Equal(t, 0, count)
	assert.Equal(t, make([]string, 0), stream.items)
}

func TestStream_setItems(t *testing.T) {
	stream := emptyStream(nil)
	stream.setItems([]string{"Home"})
	assert.Equal(t, []string{"Home"}, stream.items)
}

func TestStream_hasItems(t *testing.T) {
	stream := emptyStream(nil)
	assert.Equal(t, false, stream.hasItems())
	stream.setItems([]int{2, 3})
	assert.Equal(t, true, stream.hasItems())
	count, _ := stream.Count()
	assert.Equal(t, 2, count)
}
