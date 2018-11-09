package stream_test

import (
	"strings"
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Map(t *testing.T) {
	s := numberStream.Map(func(element int) int {
		return element + 1
	})
	array, _ := s.ToArray()
	assert.Equal(t, []int{3, 11, 4, 3}, array)
	s = numberPointerStream.Map(func(element *int) *int {
		value := *element + 1
		return &value
	})
	array, _ = s.ToArray()
	assert.Equal(t, []*int{intPtr(3), intPtr(11),
		intPtr(4), intPtr(3)}, array)

	s = textStream.Map(func(element string) string {
		return strings.ToUpper(element)
	})
	array, _ = s.ToArray()
	assert.Equal(t, []string{"HOME", "WELCOME", "SOFTWARE", "TECH", "GEEK"},
		array)
	s = textPointerStream.Map(func(element *string) *string {
		value := strings.ToUpper(*element)
		return &value
	})
	array, _ = s.ToArray()
	assert.Equal(t, []*string{stringPtr("HOME"),
		stringPtr("WELCOME"), stringPtr("SOFTWARE"), stringPtr("TECH"),
		stringPtr("GEEK")}, array)

	s = booleanStream.Map(func(element bool) bool {
		return !element
	})
	array, _ = s.ToArray()
	assert.Equal(t, []bool{true, true, false, false}, array)
	s = booleanPointerStream.Map(func(element *bool) *bool {
		value := !*element
		return &value
	})
	array, _ = s.ToArray()
	assert.Equal(t, []*bool{booleanPtr(true), booleanPtr(true),
		booleanPtr(false), booleanPtr(false)}, array)

	s = structStream.Map(func(element person) person {
		element.age += 10
		return element
	})
	array, _ = s.ToArray()
	assert.Equal(t, []person{
		{
			firstName: "John",
			age:       35,
		},
		{
			firstName: "Jane",
			age:       60,
		},
		{
			firstName: "Jean",
			age:       23,
		}}, array)
	/**
	structPointerStreamFilter := stream.new(structPointersArray)
	s = structPointerStreamFilter.Map(func(element *person) *person {
		element.age += 10
		return element
	})
	assert.Equal(t, []*person{
		{
			firstName: "John",
			age:       35,
		},
		{
			firstName: "Jane",
			age:       60,
		},
		{
			firstName: "Jean",
			age:       23,
		}}, s.ToArray())
	**/

}
func TestStream_Map_ValidationErrors(t *testing.T) {
	s := stream.new([]string{"Hello", "my", "friend"}).Map(func(element int) bool {
		return false
	})
	array, err := s.ToArray()
	assert.Nil(t, array)
	assert.Equal(t, errors.InvalidArgument("The type of the first argument must be string"), err)

	s = stream.new([]string{"Hello", "my", "friend"}).Map(func(element string) (bool, int) {
		return false, 0
	})
	array, err = s.ToArray()
	assert.Nil(t, array)
	assert.Equal(t, errors.InvalidArgument("Map function should return 1 value"), err)

	s = stream.new([]string{"Hello", "my", "friend"}).Map(func(element string, element2 bool) bool {
		return false
	})
	array, err = s.ToArray()
	assert.Nil(t, array)
	assert.Equal(t, errors.InvalidArgument("Map function should retrieve 1 argument"), err)

	s = stream.new([]string{"Hello", "my", "friend"}).Map(func(element string) int {
		return len(element)
	})
	array, err = s.ToArray()
	assert.NotNil(t, array)
	arrayInt:=array.([]int)
	assert.Equal(t, 5, arrayInt[0])
	assert.Equal(t, 2, arrayInt[1])
	assert.Equal(t, 6, arrayInt[2])

	stream := koazee.StreamOf(6)
	stream = stream.Map(strings.ToUpper)
	assert.NotNil(t, stream.Error())
	assert.Equal(t, errors.InvalidType("Unsupported type! Only arrays are permitted"), stream.Error())

}
