package stream_test

import (
	"fmt"
	"testing"

	"github.com/wesovilabs/koazee/internal/add"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	reduceInternal "github.com/wesovilabs/koazee/internal/reduce"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Reduce(t *testing.T) {
	reduce := stream.New([]int{4, 5, 6}).Reduce(func(accum, val int) int {
		return accum + val
	})
	assert.Equal(t, 15, reduce.Val())

	reduce = stream.New([]*person{
		{
			firstName: "John",
			age:       35,
		},
		{
			firstName: "Jane",
			age:       60,
		},
		{
			firstName: "Jean",
			age:       23,
		},
	}).Reduce(func(total int, val *person) int {
		return total + val.age
	})
	assert.Equal(t, 118, reduce.Val())

	reduce = stream.New([]int{1, 2, 3, 4, 5}).Reduce(func(accum, val int) (int, error) {
		return accum + val, nil
	})
	assert.Equal(t, 15, reduce.Val())
	assert.True(t, reduce.Err() == nil)

	reduce = stream.New([]int{1, 2, 3, 4, 5}).Reduce(func(accum, val int) (int, error) {
		return accum + val, fmt.Errorf("test error")
	})
	assert.Error(t, reduce.Err(), "test error")

}

func TestStream_Reduce_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The reduce operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Reduce(10).Err())

	assert.Equal(
		t,
		errors.EmptyStream(reduceInternal.OpCode, "A nil Stream can not be reduced"),
		koazee.Stream().Reduce(func() {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func() {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val bool) {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The provided function must return 1 value or the second value must be an error"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val, val2 bool) {}).Err())

	assert.True(
		t,
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val, val2 int) (int, error) { return 0, nil }).Err() == nil)

	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The type of the second argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val, val2 string) bool { return true }).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(reduceInternal.OpCode, "The type of the first argument "+
			"and the Output in the provided function must be the same"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(acc string, val int) bool { return false }).Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode,
		stream.New([]int{}).Add("home").Out().Err().Operation())
}
