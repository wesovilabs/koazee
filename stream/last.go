package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeLast identifier for operation last
const OpCodeLast = "last"

type last struct {
	items interface{}
}

func (op *last) name() string {
	return OpCodeLast
}

func (op *last) run() Output {
	if err := op.validate(); err != nil {
		return Output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(itemsValue.Len() - 1).Interface()
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, out)
	return Output{out, nil}
}

func (op *last) validate() *errors.Error {
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
func (s *Stream) Last() Output {
	current := s.run()
	if current.err != nil {
		return Output{nil, current.err}
	}
	return (&last{current.items}).run()
}
