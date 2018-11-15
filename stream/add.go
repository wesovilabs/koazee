package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeAdd identifier for operation add
const OpCodeAdd = "add"

// OperationAdd struct for defining add operation
type add struct {
	input interface{}
}

func (op *add) name() string {
	return OpCodeAdd
}

// Run performs the operations whenever is called
func (op *add) run(s *Stream) *Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	element := reflect.ValueOf(op.input)
	v := reflect.ValueOf(s.items)
	newItems := reflect.Append(v, element)
	s.items = newItems.Interface()
	return s
}

func (op *add) validate(s *Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "An element can not be added in a nil Stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	elementType := reflect.TypeOf(op.input)
	if itemsType.Kind() != reflect.Ptr && op.input == nil {
		return errors.InvalidArgument(op.name(), "A nil value can not be added in a Stream of non-pointers values")
	}
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"An element whose type is %s can not be added in a Stream of type %s", elementType, itemsType)
	}
	return nil
}

func (s *Stream) Add(input interface{}) *Stream {
	s.operations = append(s.operations, &add{input})
	return s
}
