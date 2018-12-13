package stream

import (
	"github.com/wesovilabs/koazee/internal/first"
	"reflect"
)

// First returns the first element in the stream
func (s Stream) First() *Output {
	current := s.run()
	if current.err != nil {
		return &Output{reflect.ValueOf(nil), current.err}
	}
	value, err := (&first.First{ItemsValue: current.itemsValue, Len: current.itemsLen}).Run()
	return &Output{value: value, error: err}
}
