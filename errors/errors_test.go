package errors_test

import (
	"testing"

	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestError_Error(t *testing.T) {
	err := errors.New(stream.OpCodeCount, "unknown", "argument")
	assert.Equal(t, "[count:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeAdd, "unknown", "argument")
	assert.Equal(t, "[add:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeAt, "unknown", "argument")
	assert.Equal(t, "[at:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeContains, "unknown", "argument")
	assert.Equal(t, "[contains:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeDrop, "unknown", "argument")
	assert.Equal(t, "[drop:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeFilter, "unknown", "argument")
	assert.Equal(t, "[filter:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeForEach, "unknown", "argument")
	assert.Equal(t, "[forEach:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeLast, "unknown", "argument")
	assert.Equal(t, "[last:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeMap, "unknown", "argument")
	assert.Equal(t, "[map:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeOut, "unknown", "argument")
	assert.Equal(t, "[out:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeReduce, "unknown", "argument")
	assert.Equal(t, "[reduce:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeRemoveDuplicates, "unknown", "argument")
	assert.Equal(t, "[removeDuplicates:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeSort, "unknown", "argument")
	assert.Equal(t, "[sort:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeWith, "unknown", "argument")
	assert.Equal(t, "[with:unknown] argument", err.Error())
}

func TestError_Code(t *testing.T) {
	err := errors.InvalidType(stream.OpCodeForEach, "This is so weird").
		With("when", "today")
	assert.Equal(t, errors.ErrInvalidType.String(), err.Code())
	assert.Equal(t, stream.OpCodeForEach, err.Operation())
	assert.Equal(t, "[forEach:err.invalid-type] This is so weird\n  "+
		"- when: today", err.Error())
}
