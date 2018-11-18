package first

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation first
const OpCode = "first"

type First struct {
	ItemsValue *reflect.Value
	Len        int
}

func (op *First) Run() (interface{}, *errors.Error) {
	if err := op.validate(); err != nil {
		return nil, err
	}
	out := op.ItemsValue.Index(0).Interface()
	return out, nil
}

func (op *First) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	return nil
}
