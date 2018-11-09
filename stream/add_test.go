package stream_test

import (
	"testing"

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
