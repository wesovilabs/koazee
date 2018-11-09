package stream_test

import (
	"github.com/wesovilabs/koazee/stream"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_RemoveDuplicates(t *testing.T) {
	array := stream.New([]int{10, 3, 3, 2, 10}).RemoveDuplicates().Out().Val()
	assert.Equal(t, []int{10, 3, 2}, array)
}
