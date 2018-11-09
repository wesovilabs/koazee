package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

const OperationAddIdentifier = ":add"

// OperationAdd struct for defining add operation
type add struct {
	input interface{}
}

func (op *add) name() string {
	return OperationAddIdentifier
}

// Run performs the operations whenever is called
func (op *add) run(s *stream) *stream {
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

func (op *add) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "You can not add an element in a nil stream")
	}
	itemsType := reflect.TypeOf(s.items)
	elementType := reflect.TypeOf(op.input)
	if itemsType.Kind() != reflect.Ptr && op.input == nil {
		return errors.InvalidArgument(op.name(), "You can not add a nil object in a stream of values")
	}
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"You can not add an element whose type is %s in a stream of type %s", elementType, itemsType)
	}
	return nil
}

func (s *stream) Add(input interface{}) S {
	s.operations = append(s.operations, &add{input})
	return s
}
