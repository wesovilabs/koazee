package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

/**
// Stream is an interface that provides the operations over a Stream
type Stream interface {
	// Add this operation is used to add a new element into the Stream. The element will be added in
	// the last position
	Add(interface{}) Stream
	// Drop this operation is used to drop an existing element in the Stream
	Drop(interface{}) Stream
	// First this operation is used to obtain the first element in the Stream
	First() Output
	// Last this operation is used to obtain the last element in the Stream
	Last() Output
	// Att this operation is used to obtain the element in the Stream that is in the given position
	At(int) Output
	// Reduce this operation is used to obtain a result after applying the provided function
	// over all the items in the Stream
	Reduce(function interface{}) Output
	// Reduce this operation is used to check if an element is found in the Stream
	Contains(interface{}) (bool, *errors.Error)
	// Reduce this operation is used to obtain the number of elements in the Stream
	Count() (int, *errors.Error)
	// Reduce this operation is used to convert the current elements in the Stream into a different type
	Map(interface{}) Stream
	// Reduce this operation is used to discard those elements that don't match with the provided function
	Filter(function interface{}) Stream
	// Reduce this operation is used to do something over all the elements in the Stream
	ForEach(function interface{}) Stream
	// Reduce this operation is used to sort the elements in the Stream
	Sort(function interface{}) Stream
	// Reduce this operation is used to remove duplicates elements in the Stream
	RemoveDuplicates() Stream
	// Reduce this operation is used to obtain the value of the Stream
	Out() Output
	// Reduce this operation is used to load  or replace the elements in the Stream
	With(interface{}) Stream
	// Reduce this operation is used to joing 2 or more streams in a single one
	Compose(...Stream) Stream
}
**/
// Output structure for returning single values
type Output struct {
	value interface{}
	error *errors.Error
}

// Val reurn the Output of the Stream
func (o Output) Val() interface{} { return o.value }

// Err reurn the error in the Stream
func (o Output) Err() *errors.Error { return o.error }

// Bool parses the Output of the Stream as a bool type
func (o Output) Bool() bool {
	if reflect.TypeOf(o.value).Kind() == reflect.Bool {
		return o.value.(bool)
	}
	return false
}

// String parses the Output of the Stream as a string type
func (o Output) String() string {
	if reflect.TypeOf(o.value).Kind() == reflect.String {
		return o.value.(string)
	}
	return ""
}

// Int parses the Output of the Stream as a int type
func (o Output) Int() int {
	if reflect.TypeOf(o.value).Kind() == reflect.Int {
		return o.value.(int)
	}
	return 0
}

// Int8 parses the Output of the Stream as a int8 type
func (o Output) Int8() int8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int8 {
		return o.value.(int8)
	}
	return 0
}

// Int16 parses the Output of the Stream as a int16 type
func (o Output) Int16() int16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int16 {
		return o.value.(int16)
	}
	return 0
}

// Int32 parses the Output of the Stream as a int32 type
func (o Output) Int32() int32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int32 {
		return o.value.(int32)
	}
	return 0
}

// Int64 parses the Output of the Stream as a int64 type
func (o Output) Int64() int64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Int64 {
		return o.value.(int64)
	}
	return 0
}

// Uint parses the Output of the Stream as a Uint type
func (o Output) Uint() uint {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint {
		return o.value.(uint)
	}
	return 0
}

// Uint8 parses the Output of the Stream as a Uint type8
func (o Output) Uint8() uint8 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint8 {
		return o.value.(uint8)
	}
	return 0
}

// Uint16 parses the Output of the Stream as a Uint type16
func (o Output) Uint16() uint16 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint16 {
		return o.value.(uint16)
	}
	return 0
}

// Uint32 parses the Output of the Stream as a Uint type32
func (o Output) Uint32() uint32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint32 {
		return o.value.(uint32)
	}
	return 0
}

// Uint63 parses the Output of the Stream as a Uint type64
func (o Output) Uint64() uint64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Uint64 {
		return o.value.(uint64)
	}
	return 0
}

// Float32 parses the Output of the Stream as a Uint float32
func (o Output) Float32() float32 {
	if reflect.TypeOf(o.value).Kind() == reflect.Float32 {
		return o.value.(float32)
	}
	return 0.00
}

// Float64 parses the Output of the Stream as a Uint float64
func (o Output) Float64() float64 {
	if reflect.TypeOf(o.value).Kind() == reflect.Float64 {
		return o.value.(float64)
	}
	return 0.00
}

type lazyOp interface {
	name() string
	run(*Stream) *Stream
}

type Stream struct {
	items      interface{}
	itemsValue *reflect.Value
	itemsType  reflect.Type
	err        *errors.Error
	operations []lazyOp
}

func (s *Stream) run() *Stream {
	if len(s.operations) == 0 {
		return s
	}
	s = s.operations[0].run(s)
	if s.err != nil {
		return s
	}
	s.operations = s.operations[1:]
	return s.run()
}

// New creates a Stream with the provided array of elements
func New(items interface{}) *Stream {
	if items == nil {
		return &Stream{}
	}
	itemsValue := reflect.ValueOf(items)
	itemsType := reflect.TypeOf(items).Elem()
	return &Stream{
		items:      items,
		itemsType:  itemsType,
		itemsValue: &(itemsValue),
	}

}

// Error initialize the Stream with error
func Error(err *errors.Error) *Stream {
	return &Stream{
		err: err,
	}
}
