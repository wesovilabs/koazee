package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

const OperationCountIdentifier = ".count"

type count struct {
	items interface{}
}

func (o *count) name() string {
	return OperationCountIdentifier
}

func (o *count) run() (int, *errors.Error) {
	if err := o.validate(); err != nil {
		return 0, err
	}
	itemsValue := reflect.ValueOf(o.items)
	return itemsValue.Len(), nil
}

func (o *count) validate() *errors.Error {
	if o.items == nil {
		return errors.ItemsNil(o.name(), "Count of a nil stream is not permitted")
	}
	return nil
}

// Count function that returns the number of elements in the stream
func (s *stream) Count() (int, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return 0, current.err
	}
	return (&count{current.items}).run()
}
