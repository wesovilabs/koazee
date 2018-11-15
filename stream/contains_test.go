package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Contains(t *testing.T) {
	contains, _ := stream.New([]int{1, 3}).Contains(20)
	assert.Equal(t, false, contains)
	contains, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).Contains(person{
		firstName: "Ivan",
		age:       50,
	})
	assert.Equal(t, true, contains)
	contains, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).Contains(person{
		firstName: "Ivan",
		age:       10,
	})
	assert.Equal(t, false, contains)
}

func TestStream_Contains_validate(t *testing.T) {
	contains, err := stream.New([]string{"Adopt", "a", "dog", ",", "don't", "buy"}).Contains(true)
	assert.Equal(
		t,
		false,
		contains,
	)
	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeContains,
			"The Stream contains elements of type string and the passed argument has type bool"),
		err,
	)

	contains, err = stream.New([]person{{"Iván", 34}}).Contains(nil)
	assert.Equal(
		t,
		false,
		contains,
	)
	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeContains,
			"It can not be checked if an array of non-pointers "+
				"contains a nil value"),
		err,
	)

	contains, err = stream.New(nil).Contains(5)
	assert.Equal(
		t,
		false,
		contains,
	)
	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeContains, "It can not be "+
			"checked if an element is in a nil Stream"),
		err,
	)
	_, err = stream.New([]int{}).Add("home").Contains("home")
	// To verify how errors are propagated
	assert.Equal(
		t,
		stream.OpCodeAdd, err.Operation())

}
