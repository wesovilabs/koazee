package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
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
