package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"

	"github.com/stretchr/testify/assert"
)

func TestStream_ForEach(t *testing.T) {
	s := koazee.StreamOf([]int{2, 1, 3, 2})
	array := s.ForEach(func(value int) {

	}).Out().Val()
	assert.Equal(t, []int{2, 1, 3, 2}, array)
}
