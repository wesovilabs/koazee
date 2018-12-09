package stream

import "github.com/wesovilabs/koazee/internal/deleteat"

type deleteAt struct {
	index int
}

func (m *deleteAt) run(s Stream) Stream {
	value, err := (&deleteat.DeleteAt{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Len: s.itemsLen, Index: m.index}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Filter discard the elements in the Stream that don't match with the provided filter
func (s Stream) DeleteAt(index int) Stream {
	s.operations = append(s.operations, &deleteAt{index})
	return s
}
