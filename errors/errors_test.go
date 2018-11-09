package errors_test

import (
	errors2 "errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestError_Error(t *testing.T) {
	err := errors.New("code", "message %s", "argument")
	assert.Equal(t, "code: message argument", err.Error())
}

func TestError_AddError(t *testing.T) {
	err := errors.New("code", "message %s", "parent")
	errChild1 := errors.New("code", "message child1")
	errChild2 := errors.New("code", "message child2")
	err.AddError(errChild1)
	err.AddError(errChild2)
	err.AddError(errors2.New("error common"))
	assert.Len(t, err.Children, 3)
}
