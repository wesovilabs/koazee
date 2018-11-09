package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestError_Error(t *testing.T) {
	err := errors.New(":default", "unknown", "argument")
	assert.Equal(t, "ERROR koazee:default - unknown: argument", err.Error())
}
