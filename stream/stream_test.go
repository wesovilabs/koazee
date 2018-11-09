package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_New(t *testing.T) {
	stream := New([]int{2, 3})
	count, _ := stream.Count()
	assert.Equal(t, 2, count)
	assert.Equal(t, []int{2, 3}, stream.Out().Val())
}
