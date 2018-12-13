package stream

import "github.com/wesovilabs/koazee/internal/filter"

type take struct {
	fn interface{}
}

func (m *take) run(s Stream) Stream {
	value, err := (&filter.Filter{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Filter discard the elements in the Stream that don't match with the provided filter
func (s Stream) Take(first, last int) Stream {
	s.operations = append(s.operations, &streamFilter{fn})
	return s
}
