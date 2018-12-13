package stream

import (
	"github.com/wesovilabs/koazee/internal/take"
)

type streamTake struct {
	firstIndex int
	lastIndex  int
}

func (m *streamTake) run(s Stream) Stream {
	value, err := (&take.Take{ItemsType: s.itemsType, ItemsValue: s.itemsValue,
		FirstIndex: m.firstIndex, LastIndex: m.lastIndex, Len: s.itemsValue.Len()}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Take take the elements between the given indexes
func (s Stream) Take(firstIndex, lastIndex int) Stream {
	s.operations = append(s.operations, &streamTake{firstIndex, lastIndex})
	return s
}
