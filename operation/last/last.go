package last

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation last
const OpCode = "last"

type Last struct {
	ItemsValue *reflect.Value
	Len        int
}

func (op *Last) Run() (interface{}, *errors.Error) {
	if err := op.validate(); err != nil {
		return nil, err
	}
	out := op.ItemsValue.Index(op.Len - 1).Interface()
	return out, nil
}

func (op *Last) validate() *errors.Error {
	if op.Len == 0 {
		return errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	return nil
}
