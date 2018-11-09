package stream_test

import (
	"testing"

	stream2 "github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Filter(t *testing.T) {

	stream := stream2.New([]int{3, 5, 6, 1, -1}).Filter(func(element int) bool {
		return element >= 3
	})
	array := stream.Out().Val()
	assert.Equal(t, []int{3, 5, 6}, array)
}
