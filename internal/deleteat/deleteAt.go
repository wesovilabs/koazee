package deleteat

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation at
const OpCode = "deleteAt"

// DeleteAt structure
type DeleteAt struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Len        int
	Index      int
}

// Run perform operation
func (op *DeleteAt) Run() (reflect.Value, *errors.Error) {
	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, items := dispatch(op.ItemsValue, &op.ItemsType, op.Len, op.Index); found {
		return reflect.ValueOf(items), nil
	}
	return reflect.ValueOf(nil), nil
}

func (op *DeleteAt) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be delete an element "+
			"from an empty Stream")
	}
	if op.Index < 0 || op.Len-1 < op.Index {
		return errors.InvalidIndex(OpCode,
			"The length of this Stream is %d, so the deletion index must be "+
				"between 0 and %d", op.Len, op.Len-1)
	}
	return nil
}
