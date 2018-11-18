package at

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation at
const OpCode = "at"

type At struct {
	ItemsValue *reflect.Value
	Len        int
	Index      int
}

func (op *At) Run() (interface{}, *errors.Error) {
	if err := op.validate(); err != nil {
		return nil, err
	}
	out := op.ItemsValue.Index(op.Index).Interface()
	return out, nil
}

func (op *At) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	if op.Index < 0 || op.Len-1 < op.Index {
		return errors.InvalidIndex(OpCode,
			"The length of this Stream is %d, so the index must be "+
				"between 0 and %d", op.Len, op.Len-1)
	}
	return nil
}
