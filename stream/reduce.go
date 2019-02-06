package stream

import (
	"github.com/wesovilabs/koazee/internal/reduce"
	"reflect"
)

type streamReduce struct {
	fn interface{}
}

func (m *streamReduce) run(s Stream) Output {
	currentStrema := s.run()
	value, err := (&reduce.Reduce{ItemsType: currentStrema.itemsType, ItemsValue: currentStrema.itemsValue, Func: m.fn}).Run()
	if err != nil {
		return Output{reflect.ValueOf(nil), err}
	}
	return Output{value, err}
}

// Reduce operation to calculate a single value from a stream
func (s Stream) Reduce(fn interface{}) Val {
	op := &streamReduce{fn}
	return Val{nil, op, s}
}
