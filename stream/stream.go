package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/logger"
)

// S is an interface that provides the operations over a stream
type S interface {
	// Add this operation is used to add a new element into the stream. The element will be added in the last position
	Add(interface{}) S
	// Drop this operation is used to drop an existing element in the stream
	Drop(interface{}) S
	// First this operation is used to obtain the first element in the stream
	First() output
	// Last this operation is used to obtain the last element in the stream
	Last() output
	// Att this operation is used to obtain the element in the stream that is in the given position
	At(int) output
	// Reduce this operation is used to obtain a result after applying the provided function over all the items in the stream
	Reduce(function interface{}) output
	// Reduce this operation is used to check if an element is found in the stream
	Contains(interface{}) (bool, *errors.Error)
	// Reduce this operation is used to obtain the number of elements in the stream
	Count() (int, *errors.Error)
	// Reduce this operation is used to convert the current elements in the stream into a different type
	Map(interface{}) S
	// Reduce this operation is used to discard those elements that don't match with the provided function
	Filter(function interface{}) S
	// Reduce this operation is used to do something over all the elements in the stream
	ForEach(function interface{}) S
	// Reduce this operation is used to sort the elements in the stream
	Sort(function interface{}) S
	// Reduce this operation is used to remove duplicates elements in the stream
	RemoveDuplicates() S
	// Reduce this operation is used to obtain the value of the stream
	Out() output
	// Reduce this operation is used to load  or replace the elements in the stream
	With(interface{}) S
	// Reduce this operation is used to joing 2 or more streams in a single one
	Compose(...S) S
}

type output struct {
	value interface{}
	error *errors.Error
}

// Val reurn the output of the stream
func (o output) Val() interface{}   { return o.value }
// Err reurn the error in the stream
func (o output) Err() *errors.Error { return o.error }
// Bool parses the output of the stream as a bool type
func (o output) Bool() bool {
	if reflect.TypeOf(o.value).Kind() == reflect.Bool {
		return o.value.(bool)
	}
	return false
}
// String parses the output of the stream as a string type
func (o output) String() string {
	if reflect.TypeOf(o.value).Kind() == reflect.String {
		return o.value.(string)
	}
	return ""
}
// Int parses the output of the stream as a int type
func (o output) Int() int {
	if reflect.TypeOf(o.value).Kind() == reflect.Int {
		return o.value.(int)
	}
	return 0
}
// Int8 parses the output of the stream as a int8 type
func (o output) Int8() int8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int8 {
		return o.value.(int8)
	}
	return 0
}
// Int16 parses the output of the stream as a int16 type
func (o output) Int16() int16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int16 {
		return o.value.(int16)
	}
	return 0
}
// Int32 parses the output of the stream as a int32 type
func (o output) Int32() int32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int32 {
		return o.value.(int32)
	}
	return 0
}
// Int64 parses the output of the stream as a int64 type
func (o output) Int64() int64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int64 {
		return o.value.(int64)
	}
	return 0
}
// Uint parses the output of the stream as a Uint type
func (o output) Uint() uint {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint {
		return o.value.(uint)
	}
	return 0
}
// Uint8 parses the output of the stream as a Uint type8
func (o output) Uint8() uint8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint8 {
		return o.value.(uint8)
	}
	return 0
}
// Uint16 parses the output of the stream as a Uint type16
func (o output) Uint16() uint16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint16 {
		return o.value.(uint16)
	}
	return 0
}
// Uint32 parses the output of the stream as a Uint type32
func (o output) Uint32() uint32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint32 {
		return o.value.(uint32)
	}
	return 0
}
// Uint63 parses the output of the stream as a Uint type64
func (o output) Uint64() uint64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint64 {
		return o.value.(uint64)
	}
	return 0
}
// Float32 parses the output of the stream as a Uint float32
func (o output) Float32() float32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Float32 {
		return o.value.(float32)
	}
	return 0.00
}
// Float64 parses the output of the stream as a Uint float64
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

// New creates a stream with the provided array of elements
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
