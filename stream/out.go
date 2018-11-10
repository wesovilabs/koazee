package stream

import (
	"github.com/wesovilabs/koazee/errors"
)

const OpCodeOut = "out"

type out struct {
	items interface{}
}

func (op *out) name() string {
	return OpCodeOut
}

func (op *out) run() output {
	return output{op.items, nil}
}

func (op *out) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "You can not take an element for a nil stream")
	}
	return nil
}

// At returns the element in the stream in the given position
func (s stream) Out() output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&out{current.items}).run()
}
