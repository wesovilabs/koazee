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

func (op *sort) run(s *stream) *stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	items := reflect.ValueOf(s.items)
	s.items = quickSort(items, itemsType, op.fn).Interface()
	return s
}

func (op *sort) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "You can not filter a nil stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()

	function := reflect.ValueOf(op.fn)
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
		return errors.InvalidArgument(op.name(), "The type of the output in the provided function must be int")
	}
	return nil
}

// Sort sorts the elements in the stream by applying the provided function
func (s stream) Sort(fn interface{}) S {
	s.operations = append(s.operations, &sort{fn})
	return s
}
