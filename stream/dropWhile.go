package stream

import "github.com/wesovilabs/koazee/internal/dropWhile"

type streamDropWhile struct {
	fn interface{}
}

func (m *streamDropWhile) run(s Stream) Stream {
	value, err := (&dropWhile.DropWhile{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// DropWhile discard the elements in the Stream that don't match with the provided dropWhile
func (s Stream) DropWhile(fn interface{}) Stream {
	s.operations = append(s.operations, &streamDropWhile{fn})
	return s
}
