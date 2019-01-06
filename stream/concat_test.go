package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/concat"
	"github.com/wesovilabs/koazee/stream"
)

func TestStream_Concat(t *testing.T) {
	s := stream.New([]int{3, 5, 6, 1, -1}).Concat([]int{3, 5})
	array := s.Out().Val()
	assert.Equal(t, []int{3, 5, 6, 1, -1, 3, 5}, array)

	s = stream.New([]int{2, 3, 2}).Concat([]int{})
	array = s.Out().Val()
	assert.Equal(t, []int{2, 3, 2}, array)
}

func TestStream_Concat_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(concat.OpCode, "The concat operation requires same type of array"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Concat([]int{10}).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(concat.OpCode, "The concat operation requires same type of array"),
		koazee.StreamOf([]int{2, 3, 2}).Concat(10).Out().Err())
}
