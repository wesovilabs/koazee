package stream_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"

	mapInternal "github.com/wesovilabs/koazee/internal/maps"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_Map(t *testing.T) {
	s := stream.New([]string{"HomE", "welcome", "SoftWare", "tech", "GEEK"})
	out := s.Map(strings.ToUpper).Out().Val()

	assert.Equal(t, []string{"HOME", "WELCOME", "SOFTWARE", "TECH", "GEEK"}, out)
	s = stream.New([]*person{
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
	})
	out = s.Map(func(p *person) string {
		return p.firstName
	}).Out().Val()
	assert.Equal(t, []string{"John", "Jane", "Jean"}, out)

	s = stream.New([]string{"1", "2", "3"})
	o := s.Map(strconv.Atoi).Out()
	assert.Equal(t, []int{1, 2, 3}, o.Val())
	assert.Equal(t, o.Err().UserError(), nil)

	s = stream.New([]string{"1", "2", "three"})
	o = s.Map(strconv.Atoi).Out()
	assert.Equal(t, nil, o.Val())
	assert.IsType(t, &strconv.NumError{}, o.Err().UserError())
}

func TestStream_Map_validation(t *testing.T) {
	assert.Equal(
		t,
		errors.InvalidArgument(mapInternal.OpCode, "The map operation requires a function as argument"),
		koazee.StreamOf([]string{"Freedom", "for", "the", "animals"}).Map(10).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(mapInternal.OpCode, "The provided function must retrieve 1 argument"),
		koazee.StreamOf([]int{2, 3, 2}).Map(func() {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(mapInternal.OpCode, "The type of the argument in the provided function must be int"),
		koazee.StreamOf([]int{2, 3, 2}).Map(func(val string) bool { return true }).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(mapInternal.OpCode, "The provided function must return 1 value or the second value must be an error"),
		koazee.StreamOf([]int{2, 3, 2}).Map(func(val int) {}).Out().Err())

	assert.Equal(
		t,
		errors.InvalidArgument(mapInternal.OpCode, "The provided function must return 1 value or the second value must be an error"),
		koazee.StreamOf([]int{2, 3, 2}).Map(func(val int) (string, int) { return "", 0 }).Out().Err())

	assert.Equal(
		t,
		nil,
		koazee.StreamOf([]int{2, 3, 2}).Map(func(val int) (string, error) { return "", nil }).Out().Err().UserError())
}
