package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
)

func TestStream_Add(t *testing.T) {
	stream := koazee.StreamOf(textsArray)
	counter, _ := stream.Count()
	stream.Add("animals")
	newCounter, _ := stream.Count()
	assert.Equal(t, counter+1, newCounter)

	stream = koazee.StreamOf(structPointersArray)
	counter, _ = stream.Count()
	stream.Add(&person{firstName: "Ivan", age: 34})
	newCounter, _ = stream.Count()
	assert.Equal(t, counter+1, newCounter)

}

func TestStream_Add_Validate(t *testing.T) {
	stream := koazee.StreamOf([]int{})
	stream.Add("home")
	assert.Equal(t, errors.InvalidType("stream contains elements of type int but the element is string"), stream.Error())

	stream = koazee.StreamOf([]int{})
	stream.Add(nil)
	assert.Equal(t, errors.InvalidArgument("Value whose value is nil is not permitted"), stream.Error())

}
