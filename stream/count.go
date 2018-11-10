package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeCount identifier for operation count
const OpCodeCount = "count"

type count struct {
	items interface{}
}

func (op *count) name() string {
	return OpCodeCount
}

func (op *count) run() (int, *errors.Error) {
	if err := op.validate(); err != nil {
		return 0, err
	}
	itemsValue := reflect.ValueOf(op.items)
	logger.DebugInfo("%s %v len %v", op.name(), op.items, itemsValue.Len())
	return itemsValue.Len(), nil
}

func (op *count) validate() *errors.Error {
	if op.items == nil {
		return errors.EmptyStream(op.name(), "Count of a nil stream is not permitted")
	}
	return nil
}

// Count function that returns the number of elements in the stream
func (s stream) Count() (int, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return 0, current.err
	}
	return (&count{current.items}).run()
}
