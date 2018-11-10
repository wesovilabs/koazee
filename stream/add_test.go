package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

func TestStream_Add(t *testing.T) {
	s := koazee.StreamOf([]string{"home", "welcome", "beer"})
	counter, _ := s.Count()
	newCounter, _ := s.Add("animals").Count()
	assert.Equal(t, counter+1, newCounter)

	s = koazee.StreamOf([]*person{{age: 20, firstName: "John"}})
	counter, _ = s.Count()
	newCounter, _ = s.Add(&person{firstName: "Ivan", age: 34}).Count()
	assert.Equal(t, counter+1, newCounter)
}

func TestStream_Add_validation(t *testing.T) {

	s := koazee.
		StreamOf([]string{"Freedom", "for", "the", "animals"}).
		Add(10)
	assert.Equal(t, errors.InvalidArgument(stream.OpCodeAdd,
		"An element whose type is int can not be added in a stream "+
			"of type string"), s.Out().Err())

	s = koazee.
		StreamOf([]string{"Freedom", "for", "the", "animals"}).
		Add(func() int { return 1 })
	assert.Equal(t, errors.InvalidArgument(stream.OpCodeAdd,
		"An element whose type is func() int can not be added in a stream of type string"),
		s.Out().Err())

	s = koazee.Stream().Add(10)
	assert.Equal(t, errors.EmptyStream(stream.OpCodeAdd,
		"An element can not be added in a nil stream"), s.Out().Err())

	s = koazee.Stream().Add(nil).With([]int{2, 3, 1})
	assert.Equal(t, errors.InvalidArgument(stream.OpCodeAdd, ""+
		"A nil value can not be added in a stream of non-pointers values"),
		s.Out().Err())
}
