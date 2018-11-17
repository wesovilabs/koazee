package stream

import (
	"fmt"
	Map_ "github.com/wesovilabs/koazee/internal/map"
)

type streamMap struct {
	fn interface{}
}

func (m *streamMap) run(s *Stream) *Stream {
	value, err := (&Map_.Map{s.itemsType, s.itemsValue, m.fn}).Run()
	if err != nil {
		s.err = err
		return s
	}
	fmt.Println(value)
	s.items = value
	return s
}

// Map performs a mutation over all the elements in the Stream and return a new Stream
func (s *Stream) Map(fn interface{}) *Stream {
	s.operations = append(s.operations, &streamMap{fn})
	return s
}
