package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Count(t *testing.T) {
	counter, _ := stream.New([]bool{true, false, false}).Count()
	assert.Equal(t, 3, counter)
}
