package stream

import (
	"github.com/wesovilabs/koazee/internal/reduce"
	"reflect"
)

// OpCode identifier for operation sort

type streamReduce struct {
	fn interface{}
}

func (m *streamReduce) run(s Stream) Output {
	value, err := (&reduce.Reduce{ItemsType: s.itemsType, ItemsValue: s.itemsValue, Func: m.fn}).Run()
	if err != nil {
		return Output{reflect.ValueOf(nil), err}
	}
	return Output{value, err}
}

// Reduce operation to calculate a single value from a stream
func (s Stream) Reduce(fn interface{}) StreamVal {
	op := &streamReduce{fn}
	return StreamVal{nil, op, s}
}
