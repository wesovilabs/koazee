package stream_test

import (
	"fmt"
	"testing"

	"github.com/wesovilabs/koazee"

	"github.com/wesovilabs/koazee/errors"

	"github.com/stretchr/testify/assert"
)

func TestStream_Reduce(t *testing.T) {
	reduce := numberStream.Reduce(func(accum, val int) int {
		return accum + val
	})
	assert.Equal(t, 17, reduce.Val())
	reduce = numberPointerStream.Reduce(func(total, val *int) *int {
		if total == nil {
			return val
		}
		out := *total + *val
		return &out
	})
	assert.Equal(t, intPtr(17), reduce.Val())
	reduce = textStream.Reduce(func(total, val string) string {
		return fmt.Sprintf("%s.%s", total, val)
	})
	assert.Equal(t, ".home.welcome.software.tech.geek", reduce.Val())
	reduce = textPointerStream.Reduce(func(total, val *string) *string {
		if total == nil {
			out := fmt.Sprintf("%s", *val)
			return &out
		}
		out := fmt.Sprintf("%s.%s", *total, *val)
		return &out
	})
	assert.Equal(t, "home.welcome.software.tech.geek",
		*(reduce.Val().(*string)))
	reduce = booleanStream.Reduce(func(total, val bool) bool {
		return total || val
	})
	assert.Equal(t, true, reduce.Val())
	reduce = booleanPointerStream.Reduce(func(total, val *bool) *bool {
		if total == nil {
			return val
		}
		out := *total || *val
		return &out
	})
	assert.Equal(t, booleanPtr(true), reduce.Val())
	reduce = structStream.Reduce(func(total int, val person) int {
		return total + val.age
	})
	assert.Equal(t, 88, reduce.Val())
	reduce = structPointerStream.Reduce(func(total int, val *person) int {
		return total + val.age
	})
	assert.Equal(t, 88, reduce.Val())
}

func TestStream_Reduce_ValidateErrors(t *testing.T) {
	reduce := textStream.Reduce(func(total, val string) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("The type of the first argument and the output should be the same"))

	reduce= textStream.Reduce(func(bool, val int) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("The type of the first argument and the output should be the same"))

	reduce = textStream.Reduce(func(bool, bool, string) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("Reduce function should retrieve 2 arguments"))

	reduce= textStream.Reduce(func(bool, string) (bool, string) {
		return false, ""
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("Reduce function should return 1 value"))

	reduce= textStream.Reduce(func(total, val string) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("The type of the first argument and the output should be the same"))

	reduce = textStream.Reduce(func(total, val string) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidArgument("The type of the first argument and the output should be the same"))

	nilStream := koazee.StreamOf(nil)
	reduce = nilStream.Reduce(func(bool, val int) bool {
		return false
	})
	assert.Nil(t, reduce.Val())
	assert.NotNil(t, reduce.Err())
	assert.Equal(t, reduce.Err(), errors.InvalidType("Unsupported type! Only arrays are permitted"))
}
