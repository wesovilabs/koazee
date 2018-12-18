package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/while"

	"github.com/stretchr/testify/assert"
)

func TestStream_While(t *testing.T) {
	s := koazee.StreamOf([]int{2, 1, 3, 2})
	var out []int
	array := s.While(func(value int) bool {
		return value >= 2
	}, func(value int) {
		out = append(out, value)
	}).Out().Val()
	assert.Equal(t, []int{2, 1, 3, 2}, array)
	assert.Equal(t, []int{2, 3, 2}, out)
}

func TestStream_While_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The while operation requires a function as first argument"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(1, 2).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func() {}, func() {}).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The provided function must return 1 value"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) {}, func() {}).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v string) bool { return false }, func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The type of the Output in the provided function must be bool"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) int { return 0 }, func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The while operation requires a function as second argument"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, 2).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The provided DoFunc must retrieve 1 argument"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func() {}).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The provided function can not return any value"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func(v int) int { return 0 }).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The type of the argument in the provided DoFunc must be int"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func(v string) {}).Out().Err())
}
