package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

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
	run(*Stream) *Stream
}

type Stream struct {
	items      interface{}
	itemsValue *reflect.Value
	itemsType  reflect.Type
	itemsLen   int
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
		itemsLen:   itemsValue.Len(),
	}
}

// Error initialize the Stream with error
func Error(err *errors.Error) *Stream {
	return &Stream{
		err: err,
	}
}
