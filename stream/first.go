package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

const OperationFirstIdentifier = ":first"

type first struct {
	items interface{}
}

func (op *first) name() string {
	return OperationFirstIdentifier
}

func (op *first) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	return output{itemsValue.Index(0).Interface(), nil}
}

func (op *first) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "You can not take an element for a nil stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.ItemsNil(op.name(), "You can not take an element for an empty stream")
	}
	return nil
}

// At returns the element in the stream in the given position
func (s *stream) First() output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return first{current.items,}.run()
}
