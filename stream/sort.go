package stream

import (
	"github.com/wesovilabs/koazee/internal/sort"
	"reflect"
)

// OpCode identifier for operation sort

type streamSort struct {
	fn interface{}
}

func (m *streamSort) run(s Stream) Stream {
	value, err := (&sort.Sort{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	s = s.withItemsValue(value)
	s.itemsType = reflect.TypeOf(value.Interface()).Elem()
	return s
}

// Sort sorts the elements in the Stream by applying the provided function
func (s Stream) Sort(fn interface{}) Stream {
	s.operations = append(s.operations, &streamSort{fn})
	return s
}
