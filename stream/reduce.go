package stream

import (
	"github.com/wesovilabs/koazee/operation/reduce"
	"reflect"
)

// Reduce
func (s Stream) Reduce(fn interface{}) Output {
	current := s.run()
	if current.err != nil {
		return Output{reflect.ValueOf(nil), current.err}
	}
	value, err := (&reduce.Reduce{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: fn}).Run()
	return Output{value, err}
}
