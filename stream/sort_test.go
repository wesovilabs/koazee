package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

func TestStream_Sort(t *testing.T) {
	stream := koazee.StreamOf(structsArray)
	stream = stream.Sort(func(person1, person2 person) int {
		if person1.age < person2.age {
			return -1
		} else if person1.age > person2.age {
			return 1
		}
		return 0
	})
	array := stream.Out().Val()
	assert.True(t, array.([]person)[0].age > array.([]person)[1].age)
	assert.True(t, array.([]person)[1].age > array.([]person)[2].age)

	stream = koazee.StreamOf(structsArray)
	stream = stream.Sort(func(person1, person2 person) int {
		if person1.age < person2.age {
			return 1
		} else if person1.age > person2.age {
			return -1
		}
		return 0
	})
	array = stream.Out().Val()
	assert.True(t, array.([]person)[0].age < array.([]person)[1].age)
	assert.True(t, array.([]person)[1].age < array.([]person)[2].age)
}
