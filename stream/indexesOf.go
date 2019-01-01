package stream

import (
	"github.com/wesovilabs/koazee/errors"

	"github.com/wesovilabs/koazee/internal/indexesof"
)

// IndexesOf check if the passed element is found in the Stream
func (s Stream) IndexesOf(element interface{}) ([]int, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return []int{}, current.err
	}
	if current.itemsLen == 0 {
		return []int{}, nil
	}
	return (&indexesof.IndexesOf{Items: current.itemsValue, ItemsType: current.itemsType, Element: element}).Run()
}
