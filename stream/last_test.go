package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_Last(t *testing.T) {
	last := numberStream.Last()
	assert.Equal(t, numbersArray[len(numbersArray)-1], last.Val())
	last = numberPointerStream.Last()
	assert.Equal(t, numberPointersArray[len(numberPointersArray)-1], last.Val())
	last = textStream.Last()
	assert.Equal(t, textsArray[len(textsArray)-1], last.Val())
	last = textPointerStream.Last()
	assert.Equal(t, textPointersArray[len(textPointersArray)-1], last.Val())
	last = booleanStream.Last()
	assert.Equal(t, booleansArray[len(booleansArray)-1], last.Val())
	last = booleanPointerStream.Last()
	assert.Equal(t, booleanPointersArray[len(booleanPointersArray)-1], last.Val())
	last = multiTypesStream.Last()
	assert.Equal(t, multiTypesArray[len(multiTypesArray)-1], last.Val())
	last = multiTypesPointerStream.Last()
	assert.Equal(t, multiTypePointersArray[len(multiTypePointersArray)-1], last.Val())
	last = structStream.Last()
	assert.Equal(t, structsArray[len(structsArray)-1], last.Val())
	last = structPointerStream.Last()
	assert.Equal(t, structPointersArray[len(structPointersArray)-1], last.Val())
}

func TestStream_Last_ValidationErrors(t *testing.T) {
	stream := koazee.StreamOf(6)
	last := stream.Last()
	assert.Nil(t, last.Val())
	assert.NotNil(t, last.Err())
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), last.Err())

	stream = koazee.StreamOf([]string{})
	last = stream.Last()
	assert.Nil(t, last.Val())
	assert.NotNil(t, last.Err())
	assert.Equal(t, errors.InvalidStreamIndex("Invalid index -1 in stream of 0 elements"), last.Err())
}
