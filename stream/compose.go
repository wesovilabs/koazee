package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeCompose identifier for operation compose
const OpCodeCompose = "compose"

// compose struct for defining add operation
type compose struct {
	streams []S
}

func (op *compose) name() string {
	return OpCodeCompose
}

// Run performs the operations whenever is called
func (op *compose) run(s stream) stream {
	if err := op.validate(op.streams); err != nil {
		s.err = err
		return s
	}
	s.streams = make([]stream, 0)
	for _, str := range op.streams {
		s.streams = append(s.streams, str.(stream))
	}
	return s
}

func (op *compose) validate(streams []S) *errors.Error {
	if len(streams) < 2 {
		return errors.InvalidArgument(op.name(), "To compose a new stream, 2 or more streams mut be passed as argument")
	}
	var lastType reflect.Type

	for _, s := range streams {

		if s.(stream).items != nil {
			itemsType := reflect.TypeOf(s.(stream).items)
			if itemsType != lastType && lastType != nil {
				return errors.InvalidType(op.name(), "An stream can not be composed with stream of different type")
			}
			lastType = itemsType
		}
	}
	return nil
}

func (s stream) Compose(streams ...S) S {
	s = (&compose{streams}).run(s)
	return s
}
