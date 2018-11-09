package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_First(t *testing.T) {
	value := numberStream.First()
	assert.Equal(t, numbersArray[0], value.Val())
	value = numberPointerStream.First()
	assert.Equal(t, numberPointersArray[0], value.Val())
	value = textStream.First()
	assert.Equal(t, textsArray[0], value.Val())
	value = textPointerStream.First()
	assert.Equal(t, textPointersArray[0], value.Val())
	value = booleanStream.First()
	assert.Equal(t, booleansArray[0], value.Val())
	value = booleanPointerStream.First()
	assert.Equal(t, booleanPointersArray[0], value.Val())
	value = multiTypesStream.First()
	assert.Equal(t, multiTypesArray[0], value.Val())
	value = multiTypesPointerStream.First()
	assert.Equal(t, multiTypePointersArray[0], value.Val())
	value = structStream.First()
	assert.Equal(t, structsArray[0], value.Val())
	value = structPointerStream.First()
	assert.Equal(t, structPointersArray[0], value.Val())
}

func TestStream_First_ValidationErrors(t *testing.T) {
	stream := koazee.StreamOf(6)
	value := stream.First()
	assert.Nil(t, value.Val())
	assert.NotNil(t, value.Err())
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), value.Err())

	stream = koazee.StreamOf([]string{})
	value = stream.First()
	assert.Nil(t, value.Val())
	assert.NotNil(t, value.Err())
	assert.Equal(t, errors.InvalidStreamIndex("Invalid index 0 in stream of 0 elements"), value.Err())
}
