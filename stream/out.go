package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCodeOut identifier for operation out
const OpCodeOut = "out"

type out struct {
	items reflect.Value
}

func (op *out) name() string {
	return OpCodeOut
}

func (op *out) run() *Output {
	if err := op.validate(); err != nil {
		return &Output{reflect.ValueOf(nil), err}
	}
	return &Output{op.items, nil}
}

func (op *out) validate() *errors.Error {
	if op.items == reflect.ValueOf(nil) {
		return errors.EmptyStream(op.name(), "It can not be outputted a nil Stream")
	}
	return nil
}

// At returns the element in the Stream in the given position
func (s Stream) Out() *Output {
	current := s.run()
	if current.err != nil {
		return &Output{reflect.ValueOf(nil), current.err}
	}
	return (&out{current.itemsValue}).run()
}
