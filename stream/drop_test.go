package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/internal/drop"
	"github.com/wesovilabs/koazee/utils"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
)

func TestStream_Drop(t *testing.T) {
	stream := koazee.StreamOf([]string{"home", "phone"})
	counter, _ := stream.Count()
	newCounter, _ := stream.Drop("phone").Count()
	assert.Equal(t, counter-1, newCounter)

	stream = koazee.StreamOf([]string{"home", "phone"})
	counter, _ = stream.Count()
	newCounter, _ = stream.Drop("missing").Count()
	assert.Equal(t, counter, newCounter)

	counter, _ = koazee.Stream().Drop(10).Count()
	assert.Equal(
		t,
		0,
		counter,
	)

}

func TestStream_Drop_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.InvalidArgument(drop.OpCode, "An element whose type is int "+
			"can not be dropped from a Stream of type string"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Drop(10).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(drop.OpCode,
			"A nil value can not be dropped from a Stream of non-pointers values"),
		koazee.Stream().Drop(nil).With([]int{2, 3, 1}).Out().Err())
}

func TestStream_Drop_With_Struct(t *testing.T) {
	expect := []utils.Person{
		{"John", "Doe", 23, true},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		Drop(utils.Person{"Jane", "Doe", 22, false}).
		Out().Val()
	assert.Equal(t, expect, actual)
}
