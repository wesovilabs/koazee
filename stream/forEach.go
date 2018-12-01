package stream

import (
	"github.com/wesovilabs/koazee/internal/foreach"
)

type streamForEach struct {
	fn interface{}
}

func (m *streamForEach) run(s Stream) Stream {
	value, err := (&foreach.ForEach{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// ForEach executes the provided function for all the elements in the Stream
func (s Stream) ForEach(fn interface{}) Stream {
	s.operations = append(s.operations, &streamForEach{fn})
	return s
}
