package stream

import (
	"github.com/wesovilabs/koazee/operation/duplicates"
)

type streamDuplicates struct {
}

func (m *streamDuplicates) run(s Stream) Stream {
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&duplicates.RemoveDuplicates{ItemsType: s.itemsType, ItemsValue: s.itemsValue}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// RemoveDuplicates remove all thoese elements are duplicated in the Stream, leaving only
// an element with the same value
func (s Stream) RemoveDuplicates() Stream {
	s.operations = append(s.operations, &streamDuplicates{})
	return s
}
