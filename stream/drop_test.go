package stream_test

import (
	"github.com/wesovilabs/koazee/operation/drop"
	"testing"

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
