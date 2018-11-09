package stream_test

import (
	"github.com/wesovilabs/koazee/stream"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	s := stream.New([]int{2, 3, 1, 2})
	value := s.First().Val()
	assert.Equal(t, 2, value)
}
