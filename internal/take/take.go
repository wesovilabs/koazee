package take

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Filter
const OpCode = "take"

type Take struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	FirstIndex int
	LastIndex  int
	Len        int
}

func (op *Take) Run() (reflect.Value, *errors.Error) {
	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), err
	}

	v := op.ItemsValue.Slice(op.FirstIndex, op.LastIndex)
	return v, nil
}

func (op *Take) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	if op.FirstIndex < 0 || op.Len-1 < op.LastIndex || op.FirstIndex > op.LastIndex {
		return errors.InvalidIndex(OpCode,
			"The length of this Stream is %d, so the indexes must be "+
				"between 0 and %d, being firstIndex lower thant lastIndex", op.Len, op.Len-1)
	}
	return nil
}
