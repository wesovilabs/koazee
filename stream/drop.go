package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeDrop identifier for operation drop
const OpCodeDrop = "drop"

type drop struct {
	input interface{}
}

func (op *drop) name() string {
	return OpCodeDrop
}

// Run performs the operations whenever is called
func (op *drop) run(s *Stream) *Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	element := reflect.ValueOf(op.input)
	elementType := reflect.TypeOf(op.input)
	items := reflect.ValueOf(s.items)
	newItems := reflect.MakeSlice(reflect.SliceOf(elementType), 0, 0)
	for index := 0; index < items.Len(); index++ {
		if !equalsValues(items.Index(index), element) {
			newItems = reflect.Append(newItems, items.Index(index))
		}
	}
	s.items = newItems.Interface()
	return s
}

func (op *drop) validate(s *Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "An element can not be dropped in a nil Stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	elementType := reflect.TypeOf(op.input)
	if itemsType.Kind() != reflect.Ptr && op.input == nil {
		return errors.InvalidArgument(op.name(), "A nil value can not be dropped from a Stream of non-pointers values")
	}
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"An element whose type is %s can not be dropped from a Stream of type %s", elementType, itemsType)
	}
	return nil
}

func (s *Stream) Drop(input interface{}) *Stream {
	s.operations = append(s.operations, &drop{input})
	return s
}
