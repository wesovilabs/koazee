package pop

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation pop
const OpCode = "pop"

type Pop struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Len        int
}

func (op *Pop) Run() (reflect.Value, reflect.Value, *errors.Error) {
	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), reflect.ValueOf(nil), err
	}
	if found, first, tail := dispatch(op.ItemsValue, op.ItemsType); found {
		return reflect.ValueOf(first), reflect.ValueOf(tail), nil
	}
	first, tail := op.popAny()
	return first, tail, nil
}

func (op *Pop) popAny() (reflect.Value, reflect.Value) {
	slice := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, op.Len-1)
	for index := 1; index < op.Len; index++ {
		slice = reflect.Append(slice, op.ItemsValue.Index(index))
	}
	return op.ItemsValue.Index(0), slice

}

func (op *Pop) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	return nil
}
