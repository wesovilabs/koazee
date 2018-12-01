package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeSort identifier for operation sort
const OpCodeSort = "sort"

type sort struct {
	fn interface{}
}

func (op *sort) name() string {
	return OpCodeSort
}

func (op *sort) run(s Stream) Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	items := reflect.ValueOf(s.items)
	out := quickSort(items, itemsType, op.fn)
	s = s.withItemsValue(out)
	return s
}

func (op *sort) validate(s Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "A nil Stream can not be sorted")
	}
	itemsType := reflect.TypeOf(s.items).Elem()

	function := reflect.ValueOf(op.fn)
	if function.Type().Kind() != reflect.Func {
		return errors.InvalidArgument(op.name(), "The filter operation requires a function as argument")
	}
	if function.Type().NumIn() != 2 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 2 arguments")
	}
	if function.Type().NumOut() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must return 1 value")
	}
	fnIn1 := reflect.New(function.Type().In(0)).Elem()
	fnIn2 := reflect.New(function.Type().In(1)).Elem()
	fnOut := reflect.New(function.Type().Out(0)).Elem()

	if fnIn1.Type() != fnIn2.Type() || fnIn1.Type() != itemsType {
		return errors.InvalidArgument(op.name(), "The type of the both arguments must be  %s", itemsType.String())
	}
	if fnOut.Type().Kind() != reflect.Int {
		return errors.InvalidArgument(op.name(), "The type of the Output in the provided function must be int")
	}
	return nil
}

// Sort sorts the elements in the Stream by applying the provided function
func (s Stream) Sort(fn interface{}) Stream {
	s.operations = append(s.operations, &sort{fn})
	return s
}
