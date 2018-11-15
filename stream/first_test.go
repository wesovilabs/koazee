package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	s := stream.New([]int{2, 3, 1, 2})
	value := s.First().Val()
	assert.Equal(t, 2, value)
}

func TestStream_First_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeFirst, "It can not be taken an element from a nil Stream"),
		stream.New(nil).First().Err())

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeFirst, "It can not be taken an element from an empty Stream"),
		stream.New([]int{}).First().Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		stream.OpCodeAdd,
		stream.New([]int{}).Add("home").First().Err().Operation())

}
