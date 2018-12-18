package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/internal/foreach"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_ForEach(t *testing.T) {
	s := koazee.StreamOf([]int{2, 1, 3, 2})
	array := s.ForEach(func(value int) int {
		return value + 1
	}).Out().Val()
	assert.Equal(t, []int{3, 2, 4, 3}, array)
}

func TestStream_ForEach_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The forEach operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).ForEach(10).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The provided DoFunc must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val string) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val string) int { return 0 }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The type of the Output in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val int) string { return "" }).Out().Err())

}
