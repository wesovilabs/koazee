package stream

import (
	"github.com/wesovilabs/koazee/internal/while"
)

type streamWhile struct {
	sentinel interface{}
	do       interface{}
}

func (m *streamWhile) run(s Stream) Stream {
	value, err := (&while.While{ItemsType: s.itemsType, ItemsValue: s.itemsValue, SentinelFunc: m.sentinel, DoFunc: m.do}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return s.withItemsValue(value)
}

// While executes the provided function for all the elements in the Stream
func (s Stream) While(sentinel interface{}, do interface{}) Stream {
	s.operations = append(s.operations, &streamWhile{sentinel, do})
	return s
}
