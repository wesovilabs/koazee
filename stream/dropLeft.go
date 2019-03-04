package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/internal/drop"
)

type streamDropLeft struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
	Option     drop.Option
}

func (a *streamDropLeft) run(s Stream) Stream {
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&drop.Drop{
		ItemsType:  s.itemsType,
		ItemsValue: s.itemsValue,
		Item:       a.Item,
		Option:     drop.Left,
	}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// DropLeft remove left of element that matches input
func (s Stream) DropLeft(input interface{}) Stream {
	s.operations = append(s.operations, &streamDropLeft{
		ItemsValue: s.itemsValue,
		ItemsType:  s.itemsType,
		Item:       input,
		Option:     drop.Left,
	})
	return s
}
