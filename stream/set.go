package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/set"
	"reflect"
)

type streamSet struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
	Index      int
}

func (a *streamSet) run(s Stream) Stream {
	if s.itemsLen == 0 {
		elementType := reflect.TypeOf(a.Item)
		if s.itemsType != nil && (s.itemsType != elementType) {
			s.err = errors.InvalidType(set.OpCode, "invalid type")
			return s
		}
		slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, 0)
		s = s.withItemsValue(slice)
	}
	value, err := (&set.Set{ItemsType: s.itemsType, ItemsValue: s.itemsValue, ItemValue: a.Item,
		Len: s.itemsValue.Len(), Index: a.Index}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Set operation to set a new element in the stream
func (s Stream) Set(index int, input interface{}) Stream {
	s.operations = append(s.operations, &streamSet{ItemsValue: s.itemsValue,
		ItemsType: s.itemsType, Item: input, Index: index})
	return s
}
