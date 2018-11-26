package errors_test

import (
	"github.com/wesovilabs/koazee/internal/add"
	"github.com/wesovilabs/koazee/internal/contains"
	"github.com/wesovilabs/koazee/internal/drop"
	"github.com/wesovilabs/koazee/internal/duplicates"
	"github.com/wesovilabs/koazee/internal/filter"
	"github.com/wesovilabs/koazee/internal/last"
	"testing"

	atOperation "github.com/wesovilabs/koazee/internal/at"
	mapInternal "github.com/wesovilabs/koazee/internal/maps"
	reduceInternal "github.com/wesovilabs/koazee/internal/reduce"
	"github.com/wesovilabs/koazee/stream"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
)

func TestError_Error(t *testing.T) {
	err := errors.New(add.OpCode, "unknown", "argument")
	assert.Equal(t, "[add:unknown] argument", err.Error())
	err = errors.New(atOperation.OpCode, "unknown", "argument")
	assert.Equal(t, "[at:unknown] argument", err.Error())
	err = errors.New(contains.OpCode, "unknown", "argument")
	assert.Equal(t, "[contains:unknown] argument", err.Error())
	err = errors.New(drop.OpCode, "unknown", "argument")
	assert.Equal(t, "[drop:unknown] argument", err.Error())
	err = errors.New(filter.OpCode, "unknown", "argument")
	assert.Equal(t, "[filter:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeForEach, "unknown", "argument")
	assert.Equal(t, "[forEach:unknown] argument", err.Error())
	err = errors.New(last.OpCode, "unknown", "argument")
	assert.Equal(t, "[last:unknown] argument", err.Error())
	err = errors.New(mapInternal.OpCode, "unknown", "argument")
	assert.Equal(t, "[map:unknown] argument", err.Error())
	err = errors.New(stream.OpCodeOut, "unknown", "argument")
	assert.Equal(t, "[out:unknown] argument", err.Error())
	err = errors.New(reduceInternal.OpCode, "unknown", "argument")
	assert.Equal(t, "[reduce:unknown] argument", err.Error())
	err = errors.New(duplicates.OpCode, "unknown", "argument")
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
