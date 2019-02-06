package stream

import (
	"github.com/wesovilabs/koazee/internal/first"
	"reflect"
)

type streamFirst struct {
}

func (m *streamFirst) run(s Stream) Output {
	current := s.run()
	if current.err != nil {
		return Output{reflect.ValueOf(nil), current.err}
	}
	value, err := (&first.First{ItemsValue: current.itemsValue, Len: current.itemsLen}).Run()
	return Output{value: value, error: err}
}

// Reduce operation to calculate a single value from a stream
func (s Stream) First() Val {
	op := &streamFirst{}
	return Val{nil, op, s}
}
