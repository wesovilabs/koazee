package stream

import "reflect"

type outOp interface {
	run(Stream) Output
}

type StreamVal struct {
	items  interface{}
	op     outOp
	stream Stream
}

// With load the provided elements in the stream
func (s StreamVal) With(items interface{}) Output {
	itemsValue := reflect.ValueOf(items)
	itemsType := reflect.TypeOf(items).Elem()
	s.stream.items = items
	s.stream.itemsType = itemsType
	s.stream.itemsValue = itemsValue
	s.stream.itemsLen = itemsValue.Len()
	return s.Out()
}

// Do perform actions for streamVal
func (s StreamVal) Out() Output {
	sOut := s.stream.run()
	if sOut.err != nil {
		return Output{reflect.ValueOf(nil), sOut.err}
	}
	return s.op.run(s.stream)
}
