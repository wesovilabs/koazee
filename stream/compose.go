package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

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
	if s.streams == nil {
		s.streams = make([]stream, 0)
	}
	for _, str := range op.streams {
		s.streams = append(s.streams, str.(stream))
	}
	return s
}

func (op *compose) validate(streams []S) *errors.Error {
	if streams == nil || len(streams) == 0 {
		return errors.InvalidArgument(op.name(), "You need to provide streams to compose one")
	}
	var lastType reflect.Type = nil
	for _, s := range streams {
		if s.(stream).items != nil {
			itemsType := reflect.TypeOf(s.(stream).items)
			if itemsType != lastType && lastType != reflect.TypeOf(nil) {
				return errors.InvalidType(op.name(), "You can not compose an stream if all the streams don not have the same type")
			}
		}
	}
	return nil
}

func (s stream) Compose(streams ...S) S {
	s = (&compose{streams}).run(s)
	return s
}
