package stream

import (
	"github.com/wesovilabs/koazee/logger"
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

const OperationLastIdentifier = ":last"

type last struct {
	items   interface{}
	traceID string
}

func (op *last) name() string {
	return OperationLastIdentifier
}

func (op *last) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(itemsValue.Len() - 1).Interface()
	logger.DebugInfo(op.traceID, "%s %v -> %v", op.name(), op.items, out)
	return output{out, nil}
}

func (op *last) validate() *errors.Error {
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
func (s stream) Last() output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&last{current.items, s.traceID}).run()
}
