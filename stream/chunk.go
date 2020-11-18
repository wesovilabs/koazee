package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/internal/chunk"
)

type streamChunk struct {
	size uint
}

func (m *streamChunk) run(s Stream) Stream {
	if s.itemsLen == 0 {
		result := s
		result.itemsType = reflect.SliceOf(reflect.SliceOf(s.itemsType)).Elem()
		return result
	}
	value, err := (&chunk.Chunk{ItemsValue: s.itemsValue, ItemsType: s.itemsType, Size: m.size}).Run()
	if err != nil {
		result := s
		result.itemsType = reflect.SliceOf(reflect.SliceOf(s.itemsType)).Elem()
		result.err = err
		return result
	}
	result := s.withItemsValue(value)
	result.itemsType = reflect.TypeOf(value.Interface()).Elem()
	return result
}

// Chunk the elements into the stream with split into groups the length of specified size.
func (s Stream) Chunk(size uint) Stream {
	s.operations = append(s.operations, &streamChunk{size: size})
	return s
}
