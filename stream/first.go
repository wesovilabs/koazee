package stream

import "github.com/wesovilabs/koazee/operation/first"

// At returns the element in the Stream in the given position
func (s *Stream) First() *Output {
	current := s.run()
	if current.err != nil {
		return &Output{nil, current.err}
	}
	value, err := (&first.First{ItemsValue: current.itemsValue, Len: current.itemsLen}).Run()
	return &Output{value: value, error: err}
}
