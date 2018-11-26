package stream

import (
	"github.com/wesovilabs/koazee/errors"
)

// Count function that returns the number of elements in the Stream
func (s Stream) Count() (int, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return 0, current.err
	}
	return current.itemsLen, nil
}
