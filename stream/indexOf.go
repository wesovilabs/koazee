package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/indexof"
)

// IndexOf check if the passed element is found in the Stream
func (s Stream) IndexOf(element interface{}) (int, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return indexof.InvalidIndex, current.err
	}
	if current.itemsLen == 0 {
		return indexof.InvalidIndex, nil
	}
	return (&indexof.IndexOf{Items: current.itemsValue, ItemsType: current.itemsType, Element: element}).Run()
}
