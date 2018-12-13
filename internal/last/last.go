package last

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation last
const OpCode = "last"

// Last struct for operation
type Last struct {
	ItemsValue reflect.Value
	Len        int
}

// Run performs the operation
func (op *Last) Run() (reflect.Value, *errors.Error) {
	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), err
	}
	out := op.ItemsValue.Index(op.Len - 1)
	return out, nil
}

func (op *Last) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	return nil
}
