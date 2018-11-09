package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

const OperationAtIdentifier = ":at"

type at struct {
	items   interface{}
	index   int
	traceID string
}

func (op *at) name() string {
	return OperationAtIdentifier
}

func (op *at) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(op.index).Interface()
	logger.DebugInfo(op.traceID, "%s %v -> %v", op.name(), op.items, out)
	return output{out, nil}
}

func (op *at) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "You can not take an element for a nil stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.ItemsNil(op.name(), "You can not take an element for an empty stream")
	}
	if op.index < 0 || len-1 < op.index {
		return errors.InvalidStreamIndex(op.name(), "Invalid index %d; A valid index must be between 0 and %d", op.index, len-1)
	}
	return nil
}

// At returns the element in the stream in the given position
func (s stream) At(index int) output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&at{current.items, index, s.traceID}).run()
}
