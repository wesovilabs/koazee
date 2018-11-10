package stream_test

import (
	"testing"

	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
)

func TestStream_With(t *testing.T) {
	elements := []string{"Stop", "fishing"}
	assert.Equal(
		t,
		elements,
		stream.New(nil).With(elements).Out().Val(),
	)
}

func TestStream_With_validate(t *testing.T) {

	elements := 5
	assert.Equal(
		t,
		errors.InvalidType(stream.OpCodeWith, "Unsupported type! Only arrays are permitted"),
		stream.New(nil).With(elements).Out().Err(),
	)

}
