package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/indexesof"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_IndexesOf(t *testing.T) {
	index, _ := stream.New([]int{1, 3, 5, 7}).IndexesOf(20)
	assert.Equal(t, []int{}, index)

	index, _ = stream.New([]string{"koazee", "Koazee", "koazee"}).IndexesOf("koazee")
	assert.Equal(t, []int{0, 2}, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).IndexesOf(person{
		firstName: "Ivan",
		age:       50,
	})
	assert.Equal(t, []int{1}, index)

	index, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).IndexesOf(person{
		firstName: "Ivan",
		age:       10,
	})
	assert.Equal(t, []int{}, index)
}

func TestStream_IndexesOf_validate(t *testing.T) {
	index, err := stream.New([]string{"Adopt", "a", "dog", ",", "don't", "buy"}).IndexesOf(true)
	assert.Equal(t, []int{}, index)
	assert.Equal(
		t,
		errors.InvalidArgument(indexesof.OpCode,
			"The Stream indexesOf elements of type string and the passed argument has type bool"),
		err,
	)

	index, err = stream.New([]person{{"Iván", 34}}).IndexesOf(nil)
	assert.Equal(t, []int{}, index)
	assert.Equal(
		t,
		errors.InvalidArgument(indexesof.OpCode,
			"It can not be checked if an array of non-pointers "+
				"indexesOf a nil value"),
		err,
	)

	index, _ = stream.New(nil).IndexesOf(5)
	assert.Equal(t, []int{}, index)
	_, err = stream.New([]int{}).Add("home").IndexesOf("home")
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
