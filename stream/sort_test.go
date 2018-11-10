package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

func TestStream_Sort(t *testing.T) {
	stream := koazee.StreamOf([]person{{"John", 23}, {"David", 30}, {"Michael", 27}})
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

	stream = koazee.StreamOf([]person{{"John", 23}, {"David", 30}, {"Michael", 27}})
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

func TestStream_Sort_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The filter operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Sort(10).Out().Err())

	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeSort, "A nil stream can not be sorted"),
		koazee.Stream().Sort(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The provided function must retrieve 2 arguments"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func(val int) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The provided function must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func(val int, val2 string) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The provided function must return 1 value"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func(val int, val2 string) (string, string) { return "", "" }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The type of the both arguments must be  int"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func(val int, val2 string) string { return "hi" }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(stream.OpCodeSort, "The type of the output in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).Sort(func(val int, val2 int) string { return "hi" }).Out().Err())

}
