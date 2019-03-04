package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/internal/drop"
)

type streamDropRight struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
	Option     drop.Option
}

func (a *streamDropRight) run(s Stream) Stream {
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&drop.Drop{
		ItemsType:  s.itemsType,
		ItemsValue: s.itemsValue,
		Item:       a.Item,
		Option:     drop.Right,
	}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// DropRight remove right of element that matches input
func (s Stream) DropRight(input interface{}) Stream {
	s.operations = append(s.operations, &streamDropRight{
		ItemsValue: s.itemsValue,
		ItemsType:  s.itemsType,
		Item:       input,
		Option:     drop.Right,
	})
	return s
}
