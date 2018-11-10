package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/logger"
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
	Out() output
	With(interface{}) S
	Compose(...S) S
}

// Outoput strcture for operations. It contains an element and and error
type output struct {
	value interface{}
	error *errors.Error
}

func (o output) Val() interface{}   { return o.value }
func (o output) Err() *errors.Error { return o.error }
func (o output) Bool() bool {
	if reflect.TypeOf(o.value).Kind() == reflect.Bool {
		return o.value.(bool)
	}
	return false
}
func (o output) String() string {
	if reflect.TypeOf(o.value).Kind() == reflect.String {
		return o.value.(string)
	}
	return ""
}
func (o output) Int() int {
	if reflect.TypeOf(o.value).Kind() == reflect.Int {
		return o.value.(int)
	}
	return 0
}
func (o output) Int8() int8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int8 {
		return o.value.(int8)
	}
	return 0
}
func (o output) Int16() int16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int16 {
		return o.value.(int16)
	}
	return 0
}
func (o output) Int32() int32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int32 {
		return o.value.(int32)
	}
	return 0
}
func (o output) Int64() int64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int64 {
		return o.value.(int64)
	}
	return 0
}
func (o output) Uint() uint {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint {
		return o.value.(uint)
	}
	return 0
}
func (o output) Uint8() uint8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint8 {
		return o.value.(uint8)
	}
	return 0
}
func (o output) Uint16() uint16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint16 {
		return o.value.(uint16)
	}
	return 0
}
func (o output) Uint32() uint32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint32 {
		return o.value.(uint32)
	}
	return 0
}
func (o output) Uint64() uint64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint64 {
		return o.value.(uint64)
	}
	return 0
}
func (o output) Float32() float32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Float32 {
		return o.value.(float32)
	}
	return 0.00
}
func (o output) Float64() float64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Float64 {
		return o.value.(float64)
	}
	return 0.00
}

type lazyOp interface {
	name() string
	run(*stream) *stream
}

type stream struct {
	items      interface{}
	err        *errors.Error
	streams    []stream
	operations []lazyOp
}

func (s *stream) run() *stream {
	if s.streams != nil && len(s.streams) > 0 {
		for _, child := range s.streams {
			out := child.run().Out()
			if out.error != nil {
				s.err = out.error
				return s
			}
			if s.items == nil {
				s.items = reflect.ValueOf(out.Val()).Interface()
			} else {
				s.items = reflect.AppendSlice(
					reflect.ValueOf(s.items),
					reflect.ValueOf(out.value),
				).Interface()
			}
		}
	}
	if len(s.operations) == 0 {
		return s
	}
	op := s.operations[0]
	previousItems := s.items
	s = op.run(s)
	if s.err != nil {
		return s
	}
	logger.DebugInfo("%s %v -> %v", op.name(), previousItems, s.items)
	s.operations = s.operations[1:]
	return s.run()
}

// New creates error stream with given input
func New(items interface{}) S {
	return stream{
		items: items,
	}
}

// Error initialize the stream with error
func Error(err *errors.Error) S {
	return &stream{
		err: err,
	}
}
