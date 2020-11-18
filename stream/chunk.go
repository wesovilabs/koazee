package stream

import (
	"github.com/wesovilabs/koazee/internal/chunk"
)

type streamChunk struct {
	size uint
}

// Chunk the elements into the stream with split into groups the length of specified size.
func (s Stream) Chunk(size uint) Stream {
	current := s.run()
	if current.err != nil {
		s.err = current.err
		return s
	}
	if s.itemsLen == 0 {
		return s
	}
	value, err := (&chunk.Chunk{ItemsValue: current.itemsValue, ItemsType: current.itemsType, Size: size}).Run()
	if err != nil {
		s.err = err
		return s
	}
	return New(value.Interface())
}
