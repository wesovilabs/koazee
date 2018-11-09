package koazee

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	stream := StreamOf(100)
	count, _ := stream.Count()
	assert.Equal(t, count, 0)
	assert.Equal(t, stream.Out().Err(),
		errors.InvalidType(":load", "Unsupported type! Only arrays are permitted"))
	stream = StreamOf([]string{"welcome", "to", "my", "place"})
	count, _ = stream.Count()
	assert.Equal(t, count, 4)
	assert.Nil(t, stream.Out().Err())
}
