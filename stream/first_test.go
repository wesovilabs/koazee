package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	s := stream.New([]int{2, 3, 1, 2})
	value := s.First().Val()
	assert.Equal(t, 2, value)
}
