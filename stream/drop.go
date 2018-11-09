package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

const OperationDropIdentifier = ":drop"

type drop struct {
	input interface{}
}

func (op *drop) name() string {
	return OperationDropIdentifier
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
		if !eqyalsValues(items.Index(index), element) {
			newItems = reflect.Append(newItems, items.Index(index))
		}
	}
	s.items = newItems.Interface()
	return s
}

func (op *drop) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "You can not drop an element in a nil stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	elementType := reflect.TypeOf(op.input)
	if itemsType.Kind() != reflect.Ptr && op.input == nil {
		return errors.InvalidArgument(op.name(), "You can not drop a nil object in a stream of values")
	}
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"You can not drop an element whose type is %s in a stream of type %s", elementType, itemsType)
	}
	return nil
}

func (s *stream) Drop(input interface{}) S {
	s.operations = append(s.operations, &drop{input})
	return s
}
