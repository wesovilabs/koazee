package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
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
	array, _ := stream.ToArray()
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
	array, _ = stream.ToArray()
	assert.True(t, array.([]person)[0].age < array.([]person)[1].age)
	assert.True(t, array.([]person)[1].age < array.([]person)[2].age)
}

func TestStream_Sort_ValidationErrors(t *testing.T) {
	numbers := []int{2, 4, 3, -1, 5, 9, 6, 1, 7}
	stream := koazee.StreamOf(5)
	stream.Sort(func(val int) bool {
		return true
	})
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), stream.Error())

	stream = koazee.StreamOf(numbers)
	stream.Sort(func(val int) bool {
		return true
	})
	assert.Equal(t, errors.InvalidArgument("Sort function should retrieve 2 arguments"), stream.Error())

	stream = koazee.StreamOf(numbers)
	stream.Sort(func(val int, val2 int) bool {
		return true
	})
	assert.Equal(t, errors.InvalidArgument("The output for function should be int"), stream.Error())

	stream = koazee.StreamOf(numbers)
	stream.Sort(func(val int, val2 string) bool {
		return true
	})
	assert.Equal(t, errors.InvalidArgument("Both arguments should have type int"), stream.Error())

	stream = koazee.StreamOf(numbers)
	stream.Sort(func(val int, val2 int) {

	})
	assert.Equal(t, errors.InvalidArgument("Sort function should return 1 value"), stream.Error())

}
