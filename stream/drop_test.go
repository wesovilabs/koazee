package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
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

}
func TestStream_Drop_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeDrop, "An element whose type is int "+
			"can not be dropped from a stream of type string"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Drop(10).Out().Err())

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeDrop, "An element can not be "+
			"dropped in a nil stream"),
		koazee.Stream().Drop(10).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeDrop,
			"A nil value can not be dropped from a stream of non-pointers values"),
		koazee.Stream().Drop(nil).With([]int{2, 3, 1}).Out().Err())
}
