package concat

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Concat
const OpCode = "concat"

// Concat struct for operation
type Concat struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Item       reflect.Value
}

// Run performs the operation
func (op *Concat) Run() (reflect.Value, *errors.Error) {
	err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, result := dispatch(op.ItemsValue, op.Item, op.ItemsType); found {
		return result, nil
	}
	return reflect.ValueOf(nil), err
}

func (op *Concat) validate() *errors.Error {
	if op.ItemsValue.Type() != op.Item.Type() {
		return errors.InvalidArgument(OpCode, "The concat operation requires same type of array")
	}
	return nil
}
