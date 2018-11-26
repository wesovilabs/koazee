package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Out(t *testing.T) {
	s := stream.New([]int{2, 3, 1, 2})
	value := s.Out().Val()
	assert.Equal(t, []int{2, 3, 1, 2}, value)
}

func TestStream_Out_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeOut, "It can not be outputted a nil Stream"),
		stream.New(nil).Out().Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode,
		stream.New([]int{}).Add("home").Out().Err().Operation())

}
