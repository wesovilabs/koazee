package errors_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestErrCode_String(t *testing.T) {
	var code = "myError"
	errCode := errors.ErrCode(code)
	assert.Equal(t, code, errCode.String())
}

func TestInvalidType(t *testing.T) {
	err := errors.InvalidType(stream.OpCodeAdd, "My error is %s, %d times", "fail", 3)
	assert.Equal(t, "koazee:invalid-type", err.Code())
	assert.Equal(t, "My error is fail, 3 times", err.Error())
}

func TestInvalidStreamIndex(t *testing.T) {
	err := errors.InvalidIndex(stream.OpCodeAdd, "Index not valid")
	assert.Equal(t, "koazee:index", err.Code())
	assert.Equal(t, "Index not valid", err.Error())
}

func TestItemsNil(t *testing.T) {
	err := errors.ItemsNil(stream.OpCodeAdd, "items is nil")
	assert.Equal(t, "koazee:items-nil", err.Code())
	assert.Equal(t, "items is nil", err.Error())
}

func TestInvalidArgument(t *testing.T) {
	err := errors.InvalidArgument("default", "Argument %s is not valid", "name")
	assert.Equal(t, "koazee:argument", err.Code())
	assert.Equal(t, "Argument name is not valid", err.Error())
}
