package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/internal/groupby"
	"reflect"
)

// IndexesOf check if the passed element is found in the Stream
func (s Stream) GroupBy(fn interface{}) (reflect.Value, *errors.Error) {
	current := s.run()
	if current.err != nil {
		return nil, current.err
	}
	if current.itemsLen == 0 {
		return nil, nil
	}
	return (&groupby.GroupBy{ItemsValue: current.itemsValue, ItemsType: current.itemsType, Func: fn}).Run()
}
