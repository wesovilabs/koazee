package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
)

func TestStream_Drop(t *testing.T) {
	stream := koazee.StreamOf(textsArray)
	counter, _ := stream.Count()
	stream.Drop(textsArray[1])
	newCounter, _ := stream.Count()
	assert.Equal(t, counter-1, newCounter)

	stream = koazee.StreamOf(structPointersArray)
	counter, _ = stream.Count()
	stream.Drop(&person{
		firstName: "Jane",
		age:       50,
	})
	newCounter, _ = stream.Count()
	assert.Equal(t, counter-1, newCounter)

}

func TestStream_Drop_ValidationErrors(t *testing.T) {
	stream := koazee.StreamOf([]int{})
	stream.Drop("home")
	assert.Equal(t, errors.InvalidType("stream contains elements of type int but the element is string"), stream.Error())

	stream = koazee.StreamOf([]int{})
	stream.Drop(nil)
	assert.Equal(t, errors.InvalidArgument("Value whose value is nil is not permitted"), stream.Error())

}
