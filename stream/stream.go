package stream

import (
	"github.com/wesovilabs/koazee/errors"
)

// S interface for dealing with streams
type S interface {
	Add(interface{}) S
	Drop(interface{}) S
	First() output
	Last() output
	At(int) output
	Reduce(function interface{}) output
	Contains(interface{}) (bool, *errors.Error)
	Count() (int, *errors.Error)
	Map(interface{}) S
	Filter(function interface{}) S
	ForEach(function interface{}) S
	Sort(function interface{}) S
	RemoveDuplicates() S
}

// Outoput strcture for operations. It contains an element and and error
type output struct {
	value interface{}
	error *errors.Error
}

func (o output) Val() interface{}   { return o.value }
func (o output) Err() *errors.Error { return o.error }
func (o output) Bool() bool         { return o.value.(bool) }
func (o output) String() string     { return o.value.(string) }
func (o output) Int() int           { return o.value.(int) }
func (o output) Int8() int8         { return o.value.(int8) }
func (o output) Int16() int16       { return o.value.(int16) }
func (o output) Int32() int32       { return o.value.(int32) }
func (o output) Int64() int64       { return o.value.(int64) }
func (o output) Uint() uint         { return o.value.(uint) }
func (o output) Uint8() uint8       { return o.value.(uint8) }
func (o output) Uint16() uint16     { return o.value.(uint16) }
func (o output) Uint32() uint32     { return o.value.(uint32) }
func (o output) Uint64() uint64     { return o.value.(uint64) }
func (o output) Float32() float32   { return o.value.(float32) }
func (o output) Float64() float64   { return o.value.(float64) }

type lazyOp interface {
	name() string
	run(*stream) *stream
}

type stream struct {
	items      interface{}
	err        *errors.Error
	operations []lazyOp
}

func (s *stream) run() *stream {
	if len(s.operations) <= 0 {
		return s
	}
	op := s.operations[0]
	s = op.run(s)
	if s.err != nil {
		return s
	}
	s.operations = s.operations[1:]
	return s.run()
}

// New creates error stream with given input
func New(items interface{}) S {
	return &stream{
		items: items,
	}
}

// Error initialize the stream with error
func Error(err *errors.Error) S {
	return &stream{
		err: err,
	}
}
