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

func (op *last) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(itemsValue.Len() - 1).Interface()
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, out)
	return output{out, nil}
}

func (op *last) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "It can not be taken an element from a nil stream")
	}
	itemsValue := reflect.ValueOf(op.items)
	len := itemsValue.Len()
	if len == 0 {
		return errors.ItemsNil(op.name(), "It can not be taken an element from an empty stream")
	}
	return nil
}

// At returns the element in the stream in the given position
func (s stream) Last() output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&last{current.items}).run()
}
