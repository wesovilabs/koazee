package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_Contains(t *testing.T) {
	contains, _ := numberStream.Contains(20)
	assert.Equal(t, false, contains)
	contains, _ = numberStream.Contains(numbersArray[1])
	assert.Equal(t, true, contains)
	contains, _ = numberPointerStream.Contains(intPtr(200))
	assert.Equal(t, false, contains)
	contains, _ = numberPointerStream.Contains(numberPointersArray[1])
	assert.Equal(t, true, contains)

	contains, _ = textStream.Contains("byebye")
	assert.Equal(t, false, contains)
	contains, _ = textStream.Contains(textsArray[1])
	assert.Equal(t, true, contains)
	contains, _ = textPointerStream.Contains(stringPtr("yuhuuuu"))
	assert.Equal(t, false, contains)
	contains, _ = textPointerStream.Contains(textPointersArray[1])
	assert.Equal(t, true, contains)

	contains, _ = booleanStream.Contains(false)
	assert.Equal(t, true, contains)
	contains, _ = booleanStream.Contains(booleansArray[1])
	assert.Equal(t, true, contains)
	contains, _ = booleanPointerStream.Contains(booleanPtr(true))
	assert.Equal(t, true, contains)
	contains, _ = booleanPointerStream.Contains(booleanPointersArray[1])
	assert.Equal(t, true, contains)

	contains, _ = structStream.Contains(person{
		firstName: "Juan",
		age:       50,
	})
	assert.Equal(t, false, contains)
	contains, _ = structStream.Contains(structsArray[1])
	assert.Equal(t, true, contains)
	contains, _ = structPointerStream.Contains(&person{
		firstName: "Juan",
		age:       50,
	})
	assert.Equal(t, false, contains)
	contains, _ = structPointerStream.Contains(structPointersArray[1])
	assert.Equal(t, true, contains)
}

func TestStream_Contains_ValidateErrors(t *testing.T) {
	stream := koazee.StreamOf([]int{})
	val, err := stream.Contains("home")
	assert.False(t, val)
	assert.Equal(t, errors.InvalidType("stream contains elements of type int but the element is string"), err)
	stream = koazee.StreamOf([]int{})
	val, err = stream.Contains(nil)
	assert.False(t, val)
	assert.Equal(t, errors.InvalidArgument("Value whose value is nil is not permitted"), err)
}
