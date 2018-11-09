package stream_test

import (
	"testing"

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
