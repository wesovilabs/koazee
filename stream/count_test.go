package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Count(t *testing.T) {
	counter := stream.New([]bool{true, false, false}).Count().Do().Int()
	assert.Equal(t, 3, counter)
}

func TestStream_Count_validate(t *testing.T) {

	err := stream.New([]int{}).Add("home").Count().Do().Err()
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
