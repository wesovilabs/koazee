package stream

import "reflect"

type streamCount struct {

}

func (m *streamCount) run(s Stream) Output {
	current := s.run()
	if current.err != nil {
		return Output{
			value: reflect.ValueOf(0),
			error: current.err,
		}
	}
	return Output{
		value: reflect.ValueOf(current.itemsLen),
		error: nil,
	}
}

// Count returns the element in the Stream in the given position
func (s Stream) Count() Val {
	op := &streamCount{}
	return Val{nil, op, s}
}
