package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/lastindexof"
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_LastIndexOf(t *testing.T) {
	index, _ := stream.New([]int{1, 3, 5, 7}).LastIndexOf(20)
	assert.Equal(t, lastindexof.InvalidIndex, index)

	index, _ = stream.New([]string{"koazee", "Koazee", "koazee"}).LastIndexOf("koazee")
	assert.Equal(t, 2, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).LastIndexOf(person{
		firstName: "Ivan",
		age:       50,
	})
	assert.Equal(t, 1, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).LastIndexOf(person{
		firstName: "Ivan",
		age:       10,
	})
	assert.Equal(t, lastindexof.InvalidIndex, index)
}

func TestStream_LastIndexOf_validate(t *testing.T) {
	index, err := stream.New([]string{"Adopt", "a", "dog", ",", "don't", "buy"}).LastIndexOf(true)
	assert.Equal(t, lastindexof.InvalidIndex, index)
	assert.Equal(
		t,
		errors.InvalidArgument(lastindexof.OpCode,
			"The Stream indexOf elements of type string and the passed argument has type bool"),
		err,
	)

	index, err = stream.New([]person{{"Iván", 34}}).LastIndexOf(nil)
	assert.Equal(t, lastindexof.InvalidIndex, index)
	assert.Equal(
		t,
		errors.InvalidArgument(lastindexof.OpCode,
			"It can not be checked if an array of non-pointers "+
				"indexOf a nil value"),
		err,
	)

	index, _ = stream.New(nil).LastIndexOf(5)
	assert.Equal(t, lastindexof.InvalidIndex, index)
	_, err = stream.New([]int{}).Add("home").LastIndexOf("home")
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
