package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_At(t *testing.T) {
	count, _ := numberStream.Count()
	for index := 0; index < count; index++ {
		at := numberStream.At(index)
		assert.Equal(t, numbersArray[index], at.Val())
	}
	count, _ = numberPointerStream.Count()
	for index := 0; index < count; index++ {
		at := numberPointerStream.At(index)
		assert.Equal(t, numberPointersArray[index], at.Val())
	}
	count, _ = textStream.Count()
	for index := 0; index < count; index++ {
		at := textStream.At(index)
		assert.Equal(t, textsArray[index], at.Val())
	}
	count, _ = textPointerStream.Count()
	for index := 0; index < count; index++ {
		at := textPointerStream.At(index)
		assert.Equal(t, textPointersArray[index], at.Val())
	}
	count, _ = booleanStream.Count()
	for index := 0; index < count; index++ {
		at := booleanStream.At(index)
		assert.Equal(t, booleansArray[index], at.Val())
	}
	count, _ = booleanPointerStream.Count()
	for index := 0; index < count; index++ {
		at := booleanPointerStream.At(index)
		assert.Equal(t, booleanPointersArray[index], at.Val())
	}
	count, _ = structStream.Count()
	for index := 0; index < count; index++ {
		at := structStream.At(index)
		assert.Equal(t, structsArray[index], at.Val())
	}
	count, _ = structPointerStream.Count()
	for index := 0; index < count; index++ {
		at := structPointerStream.At(index)
		assert.Equal(t, structPointersArray[index], at.Val())
	}
}
