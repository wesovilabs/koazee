package stream

import (
	"github.com/wesovilabs/koazee/internal/reduce"
)

// Reduce
func (s *Stream) Reduce(fn interface{}) Output {
	current := s.run()
	if current.err != nil {
		return Output{nil, current.err}
	}
	value, err := (&reduce.Reduce{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: fn}).Run()
	return Output{value, err}
}
