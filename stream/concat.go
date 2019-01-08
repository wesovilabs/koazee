package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/internal/concat"
)

type streamConcat struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       reflect.Value
}

func (m *streamConcat) run(s Stream) Stream {
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&concat.Concat{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Item: m.Item}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Concat concats the orderof the elements in the stream
func (s Stream) Concat(input interface{}) Stream {
	s.operations = append(s.operations, &streamConcat{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Item: reflect.ValueOf(input)})
	return s
}
