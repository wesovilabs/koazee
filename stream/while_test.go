package stream_test

import (
	"strings"
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/while"

	"github.com/stretchr/testify/assert"
)

func TestStream_While(t *testing.T) {
	s := koazee.StreamOf([]int{2, 1, 3, 2})
	array := s.While(func(value int) bool {
		return value >= 2
	}, func(value int) int {
		return value + 1
	}).Out().Val()
	assert.Equal(t, []int{3, 1, 4, 3}, array)

	s2 := koazee.StreamOf([]string{"a", "abc", "b", "bc"})
	array2 := s2.While(func(value string) bool {
		return strings.HasPrefix(value, "a")
	}, func(value string) string {
		return strings.ToUpper(value)
	}).Out().Val()
	assert.Equal(t, []string{"A", "ABC", "b", "bc"}, array2)
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
		errors.InvalidArgument(while.OpCode, "The provided DoFunc must return 1 value"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func(v int) {}).Out().Err())
	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The type of the argument in the provided DoFunc must be int"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func(v string) int { return 0 }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(while.OpCode, "The type of the Output in the provided DoFunc must be int"),
		koazee.StreamOf([]int{1, 2, 3, 4}).While(func(v int) bool { return false }, func(v int) string { return "" }).Out().Err())

}
