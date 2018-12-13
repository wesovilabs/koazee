package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/add"
	"reflect"
)

type streamAdd struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

func (a *streamAdd) run(s Stream) Stream {
	if s.itemsLen == 0 {
		elementType := reflect.TypeOf(a.Item)
		if s.itemsType != nil && (s.itemsType != elementType) {
			s.err = errors.InvalidType(add.OpCode, "invalid type")
			return s
		}
		slice := reflect.MakeSlice(reflect.SliceOf(elementType), 0, 0)
		s = s.withItemsValue(slice)
	}
	value, err := (&add.Add{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Item: a.Item}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// Add operation to add a new element in the stream
func (s Stream) Add(input interface{}) Stream {
	s.operations = append(s.operations, &streamAdd{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Item: input})
	return s
}
