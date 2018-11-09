package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

const OperationContainsIdentifier = ":contains"

type contains struct {
	items   interface{}
	element interface{}
}

func (op *contains) name() string {
	return OperationContainsIdentifier
}

func (op *contains) run() (bool, *errors.Error) {
	if err := op.validate(); err != nil {
		return false, err
	}
	items := reflect.ValueOf(op.items)
	elementValue := reflect.ValueOf(op.element)
	for index := 0; index < items.Len(); index++ {
		if eqyalsValues(items.Index(index), elementValue) {
			return true, nil
		}
	}
	return false, nil
}

func (op *contains) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "You can not take an element for a nil stream")
	}
	itemsType := reflect.TypeOf(op.items).Elem()
	if itemsType.Kind() != reflect.Ptr && op.element == nil {
		return errors.InvalidArgument(op.name(), "You can not check if an array of values contains a nil object")
	}
	elementType := reflect.TypeOf(op.element)
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"The stream contains elements of type %s and your argument has type %s", itemsType, elementType)
	}
	return nil
}

// Contains check if the passed element is found in the stream
func (s *stream) Contains(element interface{}) (bool, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return false, current.err
	}
	return contains{current.items, element}.run()
}
