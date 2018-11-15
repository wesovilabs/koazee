package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeContains identifier for operation contains
const OpCodeContains = "contains"

type contains struct {
	items   interface{}
	element interface{}
}

func (op *contains) name() string {
	return OpCodeContains
}

func (op *contains) run() (bool, *errors.Error) {
	if err := op.validate(); err != nil {
		return false, err
	}
	items := reflect.ValueOf(op.items)
	elementValue := reflect.ValueOf(op.element)
	for index := 0; index < items.Len(); index++ {
		if equalsValues(items.Index(index), elementValue) {
			logger.DebugInfo("%s %v in %v", op.name(), op.element, op.items)
			return true, nil
		}
	}
	logger.DebugInfo("%s %v not in %v", op.name(), op.element, op.items)
	return false, nil
}

func (op *contains) validate() *errors.Error {
	if op.items == nil {
		return errors.EmptyStream(op.name(), "It can not be checked if an element is in a nil Stream")
	}
	itemsType := reflect.TypeOf(op.items).Elem()
	if itemsType.Kind() != reflect.Ptr && op.element == nil {
		return errors.InvalidArgument(op.name(), "It can not be checked if an array of non-pointers contains a nil value")
	}
	elementType := reflect.TypeOf(op.element)
	if elementType != itemsType {
		return errors.InvalidArgument(op.name(),
			"The Stream contains elements of type %s and the passed argument has type %s", itemsType, elementType)
	}
	return nil
}

// Contains check if the passed element is found in the Stream
func (s *Stream) Contains(element interface{}) (bool, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return false, current.err
	}
	return (&contains{current.items, element}).run()
}
