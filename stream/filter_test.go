package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_Filter(t *testing.T) {

	stream := numberStream.Filter(func(element int) bool {
		return element >= 3
	})
	array, _ := stream.ToArray()
	assert.Equal(t, []int{10, 3}, array)
	stream = numberPointerStream.Filter(func(element *int) bool {
		return *element >= 3
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []*int{intPtr(10),
		intPtr(3)}, array)

	stream = textStream.Filter(func(element string) bool {
		return len(element) > 4
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []string{"welcome", "software"}, array)
	stream = textPointerStream.Filter(func(element *string) bool {
		return len(*element) > 4
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []*string{stringPtr("welcome"),
		stringPtr("software")}, array)

	stream = booleanStream.Filter(func(element bool) bool {
		return element
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []bool{true, true}, array)
	stream = booleanPointerStream.Filter(func(element *bool) bool {
		return !(*element)
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []*bool{booleanPtr(false),
		booleanPtr(false)}, array)

	stream = structStream.Filter(func(element person) bool {
		return element.age > 18
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []person{
		{
			firstName: "John",
			age:       25,
		},
		{
			firstName: "Jane",
			age:       50,
		},
	}, array)

	stream = structPointerStream.Filter(func(element *person) bool {
		return element.age > 18
	})
	array, _ = stream.ToArray()
	assert.Equal(t, []*person{
		{
			firstName: "John",
			age:       25,
		},
		{
			firstName: "Jane",
			age:       50,
		},
	}, array)
}
