package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeFirst identifier for operation first
const OpCodeFirst = "first"

type first struct {
	items interface{}
}

func (op *first) name() string {
	return OpCodeFirst
}

func (op *first) run() Output {
	if err := op.validate(); err != nil {
		return Output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(0).Interface()
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, out)
	return Output{out, nil}
}

func (op *first) validate() *errors.Error {
	if op.items == nil {
		return errors.EmptyStream(op.name(), "It can not be taken an element from a nil Stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.EmptyStream(op.name(), "It can not be taken an element from an empty Stream")
	}
	return nil
}

// At returns the element in the Stream in the given position
func (s *Stream) First() Output {
	current := s.run()
	if current.err != nil {
		return Output{nil, current.err}
	}
	return (&first{current.items}).run()
}
