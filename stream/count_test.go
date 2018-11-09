package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_Count(t *testing.T) {
	counter, _ := numberStream.Count()
	assert.Equal(t, len(numbersArray), counter)
	counter, _ = numberPointerStream.Count()
	assert.Equal(t, len(numberPointersArray), counter)
	counter, _ = textStream.Count()
	assert.Equal(t, len(textsArray), counter)
	counter, _ = textPointerStream.Count()
	assert.Equal(t, len(textPointersArray), counter)
	counter, _ = booleanStream.Count()
	assert.Equal(t, len(booleansArray), counter)
	counter, _ = booleanPointerStream.Count()
	assert.Equal(t, len(booleanPointersArray), counter)
	counter, _ = multiTypesStream.Count()
	assert.Equal(t, len(multiTypesArray), counter)
	counter, _ = multiTypesPointerStream.Count()
	assert.Equal(t, len(multiTypePointersArray), counter)
	counter, _ = structStream.Count()
	assert.Equal(t, len(structsArray), counter)
	counter, _ = structPointerStream.Count()
	assert.Equal(t, len(structPointersArray), counter)
}
