package stream_test

import (
	"github.com/wesovilabs/koazee/internal/dropWhile"
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_DropWhile(t *testing.T) {

	s := stream.New([]int{3, 5, 6, 1, -1}).DropWhile(func(element int) bool {
		return element >= 3
	})
	array := s.Out().Val()
	assert.Equal(t, []int{1, -1}, array)
}

func TestStream_DropWhile_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(dropWhile.OpCode, "The dropWhile operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).DropWhile(10).Out().Err())
	/**
		assert.Equal(
			t,
			errors.EmptyStream(dropWhile.OpCode, "A nil Stream can not be droped"),
			koazee.Stream().DropWhile(func() {}).Out().Err())
	**/
	assert.Equal(
		t,
		errors.InvalidArgument(dropWhile.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{2, 3, 2}).DropWhile(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(dropWhile.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).DropWhile(func(val string) bool { return true }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(dropWhile.OpCode, "The provided function must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).DropWhile(func(val int) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(dropWhile.OpCode, "The type of the Output in the provided function must be bool"),
		koazee.StreamOf([]int{2, 3, 2}).DropWhile(func(val int) string { return "a" }).Out().Err())

}
