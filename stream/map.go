package stream

import (
	Map_ "github.com/wesovilabs/koazee/operation/maps"
)

type streamMap struct {
	fn interface{}
}

func (m *streamMap) run(s Stream) Stream {
	value, err := (&Map_.Map{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Map performs a mutation over all the elements in the Stream and return a new Stream
func (s Stream) Map(fn interface{}) Stream {
	s.operations = append(s.operations, &streamMap{fn})
	return s
}
