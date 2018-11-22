package stream

import (
	"github.com/wesovilabs/koazee/operation/add"
	"reflect"
)

type streamAdd struct {
	ItemsValue *reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

func (a *streamAdd) run(s *Stream) *Stream {
	value, err := (&add.Add{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Item: a.Item}).Run()
	if err != nil {
		s.err = err
		return s
	}
	s.items = value
	s.itemsLen++
	return s
}

func (s *Stream) Add(input interface{}) *Stream {
	s.operations = append(s.operations, &streamAdd{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Item: input})
	return s
}
