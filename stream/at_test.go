package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_At(t *testing.T) {
	element := stream.New([]int{2, 3, 1}).At(1).Int()
	assert.Equal(t, 3, element)
}

func TestStream_At_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidIndex(stream.OpCodeAt, "The length of this stream is 1, so the index must be between 0 and 0"),
		stream.New([]string{"hi"}).At(-1).Err())

	assert.Equal(
		t,
		errors.InvalidIndex(stream.OpCodeAt, "The length of this stream is 1, so the index must be between 0 and 0"),
		stream.New([]string{"hi"}).At(100000).Err())

	assert.Equal(
		t,
		errors.ItemsNil(stream.OpCodeAt, "It can not be taken an element from a nil stream"),
		stream.New(nil).At(0).Err())

	assert.Equal(
		t,
		errors.ItemsNil(stream.OpCodeAt, "It can not be taken an element from an empty stream"),
		stream.New([]int{}).At(0).Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		stream.OpCodeAdd,
		stream.New([]int{}).Add("home").At(0).Err().Operation())

}
