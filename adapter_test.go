package koazee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamAdapter_toStream(t *testing.T) {
	streamAdapter := &streamAdapter{}
	stream := streamAdapter.toStream([]int{100})
	count, _ := stream.Count()
	assert.Equal(t, count, 1)
	assert.Nil(t, stream.Error())
	stream = streamAdapter.toStream([]string{"welcome", "to", "my", "place"})
	count, _ = stream.Count()
	assert.Equal(t, 4, count)
	assert.Nil(t, stream.Error())
}
