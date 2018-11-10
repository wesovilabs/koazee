package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"
)

func TestStream_Compose(t *testing.T) {
	s1 := koazee.StreamOf([]string{"Bullfighting", "is"})
	s2 := koazee.StreamOf([]string{"cruel", "and", "torture"})
	assert.Equal(
		t,
		[]string{"Bullfighting", "is", "cruel", "and", "torture"},
		koazee.Stream().Compose(s1, s2).Out().Val(),
	)
	s1 = koazee.StreamOf([]string{})
	s2 = koazee.StreamOf([]string{"cruel", "and", "torture"})
	assert.Equal(
		t,
		[]string{"cruel", "and", "torture"},
		koazee.Stream().Compose(s1, s2).Out().Val(),
	)
	s1 = koazee.Stream().With([]string{"Bullfighting", "is"})
	s2 = koazee.StreamOf([]string{"cruel", "and", "torture"})
	assert.Equal(
		t,
		[]string{"Bullfighting", "is", "cruel", "and", "torture"},
		koazee.Stream().Compose(s1, s2).Out().Val(),
	)
}

func TestStream_Compose_validation(t *testing.T) {
	s1 := koazee.StreamOf([]int{2, 3})
	s2 := koazee.StreamOf([]string{"cruel", "and", "torture"})
	assert.Equal(
		t,
		errors.InvalidType(stream.OpCodeCompose, "An stream can not be composed with stream of different type"),
		koazee.Stream().Compose(s1, s2).Out().Err(),
	)
	s1 = koazee.Stream()
	s2 = koazee.Stream()
	assert.Equal(
		t,
		errors.EmptyStream(stream.OpCodeOut, "It can not be outputted a nil stream"),
		koazee.Stream().Compose(s1, s2).Out().Err(),
	)
}
