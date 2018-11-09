package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

func TestStream_Add(t *testing.T) {
	s := koazee.StreamOf(textsArray)
	counter, _ := s.Count()
	s.Add("animals")
	newCounter, _ := s.Count()
	assert.Equal(t, counter+1, newCounter)

	s = koazee.StreamOf(structPointersArray)
	counter, _ = s.Count()
	s.Add(&person{firstName: "Ivan", age: 34})
	newCounter, _ = s.Count()
	assert.Equal(t, counter+1, newCounter)

}
