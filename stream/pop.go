package stream

import (
	"github.com/wesovilabs/koazee/internal/pop"
	"reflect"
)

// At returns the element in the Stream in the given position
func (s Stream) Pop() (*Output, Stream) {
	current := s.run()
	if current.err != nil {
		return &Output{reflect.ValueOf(nil), current.err}, current
	}
	first, items, err := (&pop.Pop{ItemsValue: current.itemsValue, ItemsType: current.itemsType, Len: current.itemsLen}).Run()
	if err != nil {
		current.err = err
		return nil, current
	}
	return &Output{value: first, error: err}, current.withItemsValue(items)
}
