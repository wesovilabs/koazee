package stream_test

import (
	"github.com/wesovilabs/koazee/stream"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	s := stream.New(numbersArray)
	value := s.First().Val()
	assert.Equal(t, numbersArray[0], value)
	s = stream.New([]*int{intPtr(2), intPtr(3)})
	value = s.First().Val()
	assert.Equal(t, intPtr(2), value)
	s = stream.New([]string{"home", "lizard", "score"})
	value = s.First().Val()
	assert.Equal(t, "home", value)
	s = stream.New([]*string{stringPtr("a"), stringPtr("b")})
	value = s.First().Val()
	assert.Equal(t, stringPtr("a"), value)
	s = stream.New([]person{{age: 2, firstName: "John"}})
	value = s.First().Val()
	assert.Equal(t, person{age: 2, firstName: "John"}, value)
	s = stream.New([]*person{{age: 2, firstName: "John"}})
	value = s.First().Val()
	assert.Equal(t, &person{age: 2, firstName: "John"}, value)
}
