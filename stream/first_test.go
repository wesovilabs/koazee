package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/first"
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	s := stream.New([]int{2, 3, 1, 2})
	value := s.First().Do().Val()
	assert.Equal(t, 2, value)
}

func TestStream_First_validation(t *testing.T) {

	assert.Equal(
		t,
		errors.EmptyStream("", "Stream cannot be run, there's no data"),
		stream.New(nil).First().Do().Err())

	assert.Equal(
		t,
		errors.EmptyStream(first.OpCode, "It can not be taken an element from an empty Stream"),
		stream.New([]int{}).First().Do().Err())

	// To verify how errors are propagated
	assert.Equal(
		t,
		add.OpCode,
		stream.New([]int{}).Add("home").First().Do().Err().Operation())

}
