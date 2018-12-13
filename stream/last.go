package stream

import (
	"github.com/wesovilabs/koazee/internal/last"
	"reflect"
)

// Last returns the last element in the stream
func (s Stream) Last() *Output {
	current := s.run()
	if current.err != nil {
		return &Output{reflect.ValueOf(nil), current.err}
	}
	value, err := (&last.Last{ItemsValue: current.itemsValue, Len: current.itemsLen}).Run()
	return &Output{value: value, error: err}
}
