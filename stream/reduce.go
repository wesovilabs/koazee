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
	value, err := reduce.Reduce{s.itemsType, s.itemsValue, fn}.Run()
	return Output{value, err}
}
