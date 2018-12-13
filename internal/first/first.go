package first

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation first
const OpCode = "first"

// First struct for the operation
type First struct {
	ItemsValue reflect.Value
	Len        int
}

// Run performs the operation
func (op *First) Run() (reflect.Value, *errors.Error) {
	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), err
	}
	out := op.ItemsValue.Index(0)
	return out, nil
}

func (op *First) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	return nil
}
