package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/indexof"
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_IndexOf(t *testing.T) {
	index, _ := stream.New([]int{1, 3,5,7}).IndexOf(20)
	assert.Equal(t, indexof.InvalidIndex, index)

	index, _ = stream.New([]string{"koazee", "Koazee","koazee"}).IndexOf("koazee")
	assert.Equal(t, 0, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).IndexOf(person{
		firstName: "Ivan",
		age:       50,
	})
	assert.Equal(t, 1, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).IndexOf(person{
		firstName: "Ivan",
		age:       10,
	})
	assert.Equal(t, indexof.InvalidIndex, index)
}

func TestStream_IndexOf_validate(t *testing.T) {
	index, err := stream.New([]string{"Adopt", "a", "dog", ",", "don't", "buy"}).IndexOf(true)
	assert.Equal(t, indexof.InvalidIndex, index)
	assert.Equal(
		t,
		errors.InvalidArgument(indexof.OpCode,
			"The Stream indexOf elements of type string and the passed argument has type bool"),
		err,
	)

	index, err = stream.New([]person{{"Iván", 34}}).IndexOf(nil)
	assert.Equal(t, indexof.InvalidIndex, index)
	assert.Equal(
		t,
		errors.InvalidArgument(indexof.OpCode,
			"It can not be checked if an array of non-pointers "+
				"indexOf a nil value"),
		err,
	)

	index, _ = stream.New(nil).IndexOf(5)
	assert.Equal(t, indexof.InvalidIndex, index)
	_, err = stream.New([]int{}).Add("home").IndexOf("home")
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
