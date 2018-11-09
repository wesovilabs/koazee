package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_At(t *testing.T) {
	element := stream.New([]int{2, 3, 1}).At(1).Int()
	assert.Equal(t, 3, element)
}
