package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/contains"
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
	c, err := stream.New([]string{"Adopt", "a", "dog", ",", "don't", "buy"}).Contains(true)
	assert.Equal(
		t,
		false,
		c,
	)
	assert.Equal(
		t,
		errors.InvalidArgument(contains.OpCode,
			"The Stream contains elements of type string and the passed argument has type bool"),
		err,
	)

	c, err = stream.New([]person{{"Iván", 34}}).Contains(nil)
	assert.Equal(
		t,
		false,
		c,
	)
	assert.Equal(
		t,
		errors.InvalidArgument(contains.OpCode,
			"It can not be checked if an array of non-pointers "+
				"contains a nil value"),
		err,
	)

	c, _ = stream.New(nil).Contains(5)
	assert.Equal(
		t,
		false,
		c,
	)

	_, err = stream.New([]int{}).Add("home").Contains("home")
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
