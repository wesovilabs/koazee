package stream_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
)

func TestStream_Add(t *testing.T) {
	s := koazee.StreamOf([]string{"home", "welcome", "beer"})
	counter := s.Count().Do().Int()
	newCounter := s.Add("animals").Count().Do().Int()
	assert.Equal(t, counter+1, newCounter)

	s = koazee.StreamOf([]*person{{age: 20, firstName: "John"}})
	counter = s.Count().Do().Int()
	newCounter = s.Add(&person{firstName: "Ivan", age: 34}).Count().Do().Int()
	assert.Equal(t, counter+1, newCounter)

	s = koazee.Stream().With([]int{1})
	newCounter = s.Add(10).Count().Do().Int()
	assert.Equal(t, 2, newCounter)

}

func TestStream_Add_validation(t *testing.T) {

	s := koazee.
		StreamOf([]string{"Freedom", "for", "the", "animals"}).
		Add(10)
	assert.Equal(t, errors.InvalidArgument(add.OpCode,
		"An element whose type is int can not be added in a Stream "+
			"of type string"), s.Out().Err())

	s = koazee.
		StreamOf([]string{"Freedom", "for", "the", "animals"}).
		Add(func() int { return 1 })
	assert.Equal(t, errors.InvalidArgument(add.OpCode,
		"An element whose type is func() int can not be added in a Stream of type string"),
		s.Out().Err())

	s = koazee.Stream().Add(nil).With([]int{2, 3, 1})
	assert.Equal(t, errors.InvalidArgument(add.OpCode, ""+
		"A nil value can not be added in a Stream of non-pointers values"),
		s.Out().Err())
}
