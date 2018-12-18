package stream_test

import (
	"fmt"
	"testing"

	"github.com/wesovilabs/koazee/internal/foreach"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_ForEach(t *testing.T) {
	s := koazee.StreamOf([]int{2, 1, 3, 2})
	array := s.ForEach(func(value int) {

	}).Out().Val()
	assert.Equal(t, []int{2, 1, 3, 2}, array)

	s = koazee.StreamOf([]int{1, 2, 3, 4})
	err := s.ForEach(func(value int) error {
		return fmt.Errorf("test error")
	}).Out().Err().UserError()
	assert.EqualError(t, err, "test error")
}

func TestStream_ForEach_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The forEach operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).ForEach(10).Out().Err())
	/**
		assert.Equal(
			t,
			errors.EmptyStream(foreach.OpCode, "A nil Stream can not be used to perform ForEach operation"),
			koazee.Stream().ForEach(func() {}).Out().Err())
	**/
	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val string) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The provided function can not return any value or must return only an error"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val int) bool { return false }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(foreach.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val, val2 int) {}).Out().Err())

	err := koazee.StreamOf([]int{2, 3, 2}).ForEach(func(val int) error { return nil }).Out().Err()
	assert.True(t, err == nil)
}
