package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/operation/contains"
)

// Contains check if the passed element is found in the Stream
func (s *Stream) Contains(element interface{}) (bool, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return false, current.err
	}
	return (&contains.Contains{Items: current.itemsValue, ItemsType: current.itemsType, Element: element}).Run()
}
