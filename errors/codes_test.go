package errors_test

import (
	"github.com/wesovilabs/koazee/operation/add"
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
	err := errors.InvalidType(add.OpCode, "My error is %s, %d times", "fail", 3)
	assert.Equal(t, "err.invalid-type", err.Code())
	assert.Equal(t, "[add:err.invalid-type] My error is fail, 3 times", err.Error())
}

func TestInvalidStreamIndex(t *testing.T) {
	err := errors.InvalidIndex(add.OpCode, "Index not valid")
	assert.Equal(t, "err.invalid-index", err.Code())
	assert.Equal(t, "[add:err.invalid-index] Index not valid", err.Error())
}

func TestItemsNil(t *testing.T) {
	err := errors.EmptyStream(add.OpCode, "items is nil")
	assert.Equal(t, "err.items-nil", err.Code())
	assert.Equal(t, "[add:err.items-nil] items is nil", err.Error())
}

func TestInvalidArgument(t *testing.T) {
	err := errors.InvalidArgument("default", "Argument %s is not valid", "name")
	assert.Equal(t, "err.invalid-argument", err.Code())
	assert.Equal(t, "[default:err.invalid-argument] Argument name is not valid", err.Error())
}
