package stream_test

import (
	"github.com/wesovilabs/koazee/stream"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_Count(t *testing.T) {
	counter, _ := stream.New([]bool{true,false,false}).Count()
	assert.Equal(t, 3, counter)
}
