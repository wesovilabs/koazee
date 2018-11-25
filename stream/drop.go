package stream

import (
	"github.com/wesovilabs/koazee/operation/drop"
	"reflect"
)

type streamDrop struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

func (a *streamDrop) run(s Stream) Stream {
	if s.itemsLen==0{
		return s
	}
	value, err := (&drop.Drop{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Item: a.Item}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

func (s Stream) Drop(input interface{}) Stream {
	s.operations = append(s.operations, &streamDrop{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Item: input})
	return s
}
