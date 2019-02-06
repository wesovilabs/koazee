package koazee

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	stream := StreamOf(100)
	assert.Equal(t, stream.Out().Err(),
		errors.InvalidType(":load", "Unsupported type! Only arrays are permitted"))
	stream = StreamOf([]string{"welcome", "to", "my", "place"})
	count:= stream.Count().Do().Int()
	assert.Equal(t, count, 4)
}
