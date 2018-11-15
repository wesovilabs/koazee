package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeAt identifier for operation at
const OpCodeAt = "at"

type at struct {
	items interface{}
	index int
}

func (op *at) name() string {
	return OpCodeAt
}

func (op *at) run() Output {
	if err := op.validate(); err != nil {
		return Output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(op.index).Interface()
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, out)
	return Output{out, nil}
}

func (op *at) validate() *errors.Error {
	if op.items == nil {
		return errors.EmptyStream(op.name(), "It can not be taken an element "+
			"from a nil Stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.EmptyStream(op.name(), "It can not be taken an element "+
			"from an empty Stream")
	}
	if op.index < 0 || len-1 < op.index {
		return errors.InvalidIndex(op.name(),
			"The length of this Stream is %d, so the index must be "+
				"between 0 and %d", len, len-1)
	}
	return nil
}

// At returns the element in the Stream in the given position
func (s *Stream) At(index int) Output {
	current := s.run()
	if current.err != nil {
		return Output{nil, current.err}
	}
	return (&at{current.items, index}).run()
}
