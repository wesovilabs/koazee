package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	atOperation "github.com/wesovilabs/koazee/internal/at"
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
		errors.InvalidIndex(atOperation.OpCode, "The length of this Stream is 1, so the index must be between 0 and 0"),
		stream.New([]string{"hi"}).At(-1).Err())

	assert.Equal(
		t,
		errors.InvalidIndex(atOperation.OpCode, "The length of this Stream is 1, so the index must be between 0 and 0"),
		stream.New([]string{"hi"}).At(100000).Err())

	assert.Equal(
		t,
		errors.EmptyStream(atOperation.OpCode, "It can not be taken an element from an empty Stream"),
		stream.New(nil).At(0).Err())

	assert.Equal(
		t,
		errors.EmptyStream(atOperation.OpCode, "It can not be taken an element from an empty Stream"),
		stream.New([]int{}).At(0).Err())

	// To verify how errors are propagated
	/**
		assert.Equal(
			t,
			add.OpCode,
			stream.New([]int{}).Add("home").At(0).Err().Operation())
	**/
}
