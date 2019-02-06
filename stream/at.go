package stream

import (
	at2 "github.com/wesovilabs/koazee/internal/at"
	"reflect"
)

type streamAt struct {
	index int
}

func (m *streamAt) run(s Stream) Output {
	current := s.run()
	if current.err != nil {
		return Output{reflect.ValueOf(nil), current.err}
	}
	val, err := (&at2.At{ItemsValue: current.itemsValue, Len: current.itemsLen, Index: m.index}).Run()
	return Output{val, err}
}

// At returns the element in the Stream in the given position
func (s Stream) At(index int) Val {
	op := &streamAt{index}
	return Val{nil, op, s}
}
