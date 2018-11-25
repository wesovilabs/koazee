package stream

import (
	at2 "github.com/wesovilabs/koazee/operation/at"
	"reflect"
)

// At returns the element in the Stream in the given position
func (s *Stream) At(index int) *Output {
	current := s.run()
	if current.err != nil {
		return &Output{reflect.ValueOf(nil), current.err}
	}
	val, err := (&at2.At{ItemsValue: current.itemsValue, Len: current.itemsLen, Index: index}).Run()
	return &Output{val, err}
}
