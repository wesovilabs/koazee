package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Reduce(t *testing.T) {
	reduce := stream.New([]int{4, 5, 6}).Reduce(func(accum, val int) int {
		return accum + val
	})
	assert.Equal(t, 15, reduce.Val())

	reduce = stream.New([]*person{
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
		},
	}).Reduce(func(total int, val *person) int {
		return total + val.age
	})
	assert.Equal(t, 118, reduce.Val())
}
