package stream

import (
	"github.com/wesovilabs/koazee/internal/reduce"
	"reflect"
)

// Reduce operation to calculate a single value from a stream
func (s Stream) Reduce(fn interface{}) Output {
	current := s.run()
	if current.err != nil {
		return Output{reflect.ValueOf(nil), current.err}
	}
	value, err := (&reduce.Reduce{ItemsType: current.itemsType, ItemsValue: current.itemsValue, Func: fn}).Run()
	return Output{value, err}
}
