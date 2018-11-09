package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_RemoveDuplicates(t *testing.T) {
	array, _ := numberStream.RemoveDuplicates().ToArray()
	assert.Equal(t, []int{2, 10, 3}, array)
	array, _ = booleanStream.RemoveDuplicates().ToArray()
	assert.Equal(t, []bool{false, true}, array)
	array, _ = numberPointerStream.RemoveDuplicates().ToArray()
	assert.Equal(t, []*int{intPtr(2), intPtr(10), intPtr(3)}, array)
}
