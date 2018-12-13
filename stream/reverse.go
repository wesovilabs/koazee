package stream

import (
	"github.com/wesovilabs/koazee/internal/reverse"
)

type streamReverse struct {
}

func (m *streamReverse) run(s Stream) Stream {
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&reverse.Reverse{ItemsType: s.itemsType, ItemsValue: s.itemsValue}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Reverse reverses the orderof the elements in the stream
func (s Stream) Reverse() Stream {
	s.operations = append(s.operations, &streamReverse{})
	return s
}
