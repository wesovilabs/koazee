package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

const OpCodeFirst = ":first"

type first struct {
	items interface{}
}

func (op *first) name() string {
	return OpCodeFirst
}

func (op *first) run() output {
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	itemsValue := reflect.ValueOf(op.items)
	out := itemsValue.Index(0).Interface()
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, out)
	return output{out, nil}
}

func (op *first) validate() *errors.Error {
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
func (s stream) First() output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&first{current.items}).run()
}
