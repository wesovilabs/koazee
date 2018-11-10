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
func (op *drop) run(s *stream) *stream {
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

func (op *drop) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "An element can not be dropped in a nil stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	elementType := reflect.TypeOf(op.input)
	if itemsType.Kind() != reflect.Ptr && op.input == nil {
		return errors.InvalidArgument(op.name(), "A nil value can not be dropped from a stream of non-pointers values")
	}
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"An element whose type is %s can not be dropped from a stream of type %s", elementType, itemsType)
	}
	return nil
}

func (s stream) Drop(input interface{}) S {
	s.operations = append(s.operations, &drop{input})
	return s
}
