package stream_test

import (
	"github.com/wesovilabs/koazee/operation/add"
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Count(t *testing.T) {
	counter, _ := stream.New([]bool{true, false, false}).Count()
	assert.Equal(t, 3, counter)
}

func TestStream_Count_validate(t *testing.T) {

	count, err := stream.New(nil).Count()
	assert.Nil(t, err)
	assert.Equal(
		t,
		0,
		count,
	)

	_, err = stream.New([]int{}).Add("home").Count()
	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode, err.Operation())

}
