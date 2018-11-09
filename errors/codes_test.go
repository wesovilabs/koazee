package errors_test

import (
	"github.com/wesovilabs/koazee/stream"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestErrCode_String(t *testing.T) {
	var code = "myError"
	errCode := errors.ErrCode(code)
	assert.Equal(t, code, errCode.String())
}

func TestInvalidType(t *testing.T) {
	err := errors.InvalidType(stream.OperationAddIdentifier,"My error is %s, %d times", "fail", 3)
	assert.Equal(t, "koazee:invalid-type", err.Code.String())
	assert.Equal(t, "My error is fail, 3 times", err.Msg)
}

func TestInvalidStreamIndex(t *testing.T) {
	err := errors.InvalidStreamIndex(stream.OperationAddIdentifier,"Index not valid")
	assert.Equal(t, "koazee:index", err.Code.String())
	assert.Equal(t, "Index not valid", err.Msg)
}

func TestItemsNil(t *testing.T) {
	err := errors.ItemsNil(stream.OperationAddIdentifier,"items is nil")
	assert.Equal(t, "koazee:items-nil", err.Code.String())
	assert.Equal(t, "items is nil", err.Msg)
}

func TestInvalidArgument(t *testing.T) {
	err := errors.InvalidArgument("default","Argument %s is not valid", "name")
	assert.Equal(t, "koazee:argument", err.Code.String())
	assert.Equal(t, "Argument name is not valid", err.Msg)
}
