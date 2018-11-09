package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee"

	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_At(t *testing.T) {
	count, _ := numberStream.Count()
	for index := 0; index < count; index++ {
		at:= numberStream.At(index)
		assert.Equal(t, numbersArray[index], at.Val())
	}
	count, _ = numberPointerStream.Count()
	for index := 0; index < count; index++ {
		at:= numberPointerStream.At(index)
		assert.Equal(t, numberPointersArray[index], at.Val())
	}
	count, _ = textStream.Count()
	for index := 0; index < count; index++ {
		at:= textStream.At(index)
		assert.Equal(t, textsArray[index], at.Val())
	}
	count, _ = textPointerStream.Count()
	for index := 0; index < count; index++ {
		at:= textPointerStream.At(index)
		assert.Equal(t, textPointersArray[index], at.Val())
	}
	count, _ = booleanStream.Count()
	for index := 0; index < count; index++ {
		at:= booleanStream.At(index)
		assert.Equal(t, booleansArray[index], at.Val())
	}
	count, _ = booleanPointerStream.Count()
	for index := 0; index < count; index++ {
		at:= booleanPointerStream.At(index)
		assert.Equal(t, booleanPointersArray[index], at.Val())
	}
	count, _ = structStream.Count()
	for index := 0; index < count; index++ {
		at:= structStream.At(index)
		assert.Equal(t, structsArray[index], at.Val())
	}
	count, _ = structPointerStream.Count()
	for index := 0; index < count; index++ {
		at:= structPointerStream.At(index)
		assert.Equal(t, structPointersArray[index], at.Val())
	}
}

func TestStream_At_Validate(t *testing.T) {
	count, _ := numberStream.Count()
	at:= numberStream.At(count)
	assert.Nil(t, at.Val())
	assert.NotNil(t, at.Err())
	assert.Equal(t, errors.InvalidStreamIndex("Invalid index 4 in stream of 4 elements"), at.Err())

	stream := koazee.StreamOf(6)
	at= stream.At(4)
	assert.Nil(t, at.Val())
	assert.NotNil(t, at.Err())
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), at.Err())

}
