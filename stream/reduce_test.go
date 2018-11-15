package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

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
}

func TestStream_Reduce_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The reduce operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Reduce(10).Err())

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeReduce, "A nil Stream can not be reduced"),
		koazee.Stream().Reduce(func() {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func() {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val bool) {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The provided function must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val, val2 bool) {}).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The type of the second argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(val, val2 string) bool { return true }).Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeReduce, "The type of the first argument "+
			"and the Output in the provided function must be the same"),
		koazee.StreamOf([]int{2, 3, 2}).Reduce(func(acc string, val int) bool { return false }).Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		stream.OpCodeAdd,
		stream.New([]int{}).Add("home").Out().Err().Operation())
}
