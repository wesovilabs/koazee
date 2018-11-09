package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
)

func TestStream_ForEach(t *testing.T) {
	array, _ := numberStream.ForEach(func(value int) {

	}).ToArray()
	assert.Equal(t, numbersArray, array)
}

func TestStream_ForEach_ValidationErrors(t *testing.T) {
	stream := koazee.StreamOf(6)
	_, err := stream.ForEach(func(item int) {

	}).ToArray()
	assert.NotNil(t, err)
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), err)

	stream = koazee.StreamOf(numbersArray)
	s := stream.ForEach(func(item string) {

	})
	assert.NotNil(t, s.Error())
	assert.Equal(t, errors.InvalidArgument("The type of the first argument must be int"), s.Error())

	stream = koazee.StreamOf(numbersArray)
	s = stream.ForEach(func(item int, item2 int) {

	})
	assert.NotNil(t, s.Error())
	assert.Equal(t, errors.InvalidArgument("ForEach function should retrieve 1 argument"), s.Error())

	stream = koazee.StreamOf(numbersArray)
	s = stream.ForEach(func(item int) bool {
		return false
	})
	assert.NotNil(t, s.Error())
	assert.Equal(t, errors.InvalidArgument("ForEach function should not return any output"), s.Error())
}
