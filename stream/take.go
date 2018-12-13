package stream

import "github.com/wesovilabs/koazee/internal/take"

type take struct {
	fn interface{}
}

func (m *take) run(s Stream) Stream {
	value, err := (&take.Take{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Take takes the elements from the stream
func (s Stream) Take(first, last int) Stream {
	s.operations = append(s.operations, &take{first, last})
	return s
}
