package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// Output structure for returning single values
type Output struct {
	value reflect.Value
	error *errors.Error
}

// Val reurn the Output of the Stream
func (o Output) Val() interface{} {
	v := (o.value)
	if !o.value.IsValid() {
		return nil
	}
	return v.Interface()
}

// Err reurn the error in the Stream
func (o Output) Err() *errors.Error {
	if o.error == nil || o.error.Code() == "" {
		return nil
	}
	return o.error
}

// Bool parses the Output of the Stream as a bool type
func (o Output) Bool() bool {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Bool {
		return o.value.Interface().(bool)
	}
	return false
}

// String parses the Output of the Stream as a string type
func (o Output) String() string {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.String {
		return o.value.Interface().(string)
	}
	return ""
}

// Int parses the Output of the Stream as a int type
func (o Output) Int() int {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Int {
		return o.value.Interface().(int)
	}
	return 0
}

// Int8 parses the Output of the Stream as a int8 type
func (o Output) Int8() int8 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Int8 {
		return o.value.Interface().(int8)
	}
	return 0
}

// Int16 parses the Output of the Stream as a int16 type
func (o Output) Int16() int16 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Int16 {
		return o.value.Interface().(int16)
	}
	return 0
}

// Int32 parses the Output of the Stream as a int32 type
func (o Output) Int32() int32 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Int32 {
		return o.value.Interface().(int32)
	}
	return 0
}

// Int64 parses the Output of the Stream as a int64 type
func (o Output) Int64() int64 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Int64 {
		return o.value.Interface().(int64)
	}
	return 0
}

// Uint parses the Output of the Stream as a Uint type
func (o Output) Uint() uint {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Uint {
		return o.value.Interface().(uint)
	}
	return 0
}

// Uint8 parses the Output of the Stream as a Uint type8
func (o Output) Uint8() uint8 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Uint8 {
		return o.value.Interface().(uint8)
	}
	return 0
}

// Uint16 parses the Output of the Stream as a Uint type16
func (o Output) Uint16() uint16 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Uint16 {
		return o.value.Interface().(uint16)
	}
	return 0
}

// Uint32 parses the Output of the Stream as a Uint type32
func (o Output) Uint32() uint32 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Uint32 {
		return o.value.Interface().(uint32)
	}
	return 0
}

// Uint64 parses the Output of the Stream as a Uint type64
func (o Output) Uint64() uint64 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Uint64 {
		return o.value.Interface().(uint64)
	}
	return 0
}

// Float32 parses the Output of the Stream as a Uint float32
func (o Output) Float32() float32 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Float32 {
		return o.value.Interface().(float32)
	}
	return 0.00
}

// Float64 parses the Output of the Stream as a Uint float64
func (o Output) Float64() float64 {
	if reflect.TypeOf(o.value.Interface()).Kind() == reflect.Float64 {
		return o.value.Interface().(float64)
	}
	return 0.00
}

type lazyOp interface {
	run(Stream) Stream
}

// Stream stream structure
type Stream struct {
	items      interface{}
	itemsValue reflect.Value
	itemsType  reflect.Type
	itemsLen   int
	err        *errors.Error
	operations []lazyOp
}

func (s Stream) run() Stream {
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
func New(items interface{}) Stream {
	s := Stream{}
	if items != nil {
		itemsValue := reflect.ValueOf(items)
		itemsType := reflect.TypeOf(items).Elem()
		s.items = items
		s.itemsType = itemsType
		s.itemsValue = itemsValue
		s.itemsLen = itemsValue.Len()
	}
	return s
}

func (s Stream) withItemsValue(items reflect.Value) Stream {
	s.itemsValue = items
	s.items = items.Interface()
	s.itemsLen = items.Len()
	return s
}

// Error initialize the Stream with error
func Error(err *errors.Error) Stream {
	return Stream{
		err: err,
	}
}
