package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Last(t *testing.T) {
	s := stream.New([]string{"hello", "my", "friend"})
	last := s.Last().Val()
	assert.Equal(t, "friend", last)
}
