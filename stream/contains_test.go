package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Contains(t *testing.T) {
	contains, _ := stream.New([]int{1, 3}).Contains(20)
	assert.Equal(t, false, contains)
	contains, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).Contains(person{
		firstName: "Ivan",
		age:       50,
	})
	assert.Equal(t, true, contains)
	contains, _ = stream.New([]person{{firstName: "Ivan", age: 20}, {firstName: "Ivan", age: 50}}).Contains(person{
		firstName: "Ivan",
		age:       10,
	})
	assert.Equal(t, false, contains)
}
