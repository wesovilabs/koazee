package stream

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

type outOp interface {
	run(Stream) Output
}

type Val struct {
	items  interface{}
	op     outOp
	stream Stream
}

// With load the provided elements in the stream
func (s Val) With(items interface{}) Output {
	itemsValue := reflect.ValueOf(items)
	itemsType := reflect.TypeOf(items).Elem()
	s.stream.items = items
	s.stream.itemsType = itemsType
	s.stream.itemsValue = itemsValue
	s.stream.itemsLen = itemsValue.Len()
	return s.Do()
}

// Do perform actions for streamVal
func (s Val) Do() Output {
	if s.stream.items == nil {
		return Output{reflect.ValueOf(nil),
			errors.New("", errors.ErrItemsNil, "Stream cannot be run, there's no data")}
	}
	sOut := s.stream.run()
	if sOut.err != nil {
		return Output{reflect.ValueOf(nil), sOut.err}
	}
	return s.op.run(s.stream)
}
