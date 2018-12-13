package reverse

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
	"sort"
)

// OpCode identifier for operation Reverse
const OpCode = "reverse"

// Reverse operation struct
type Reverse struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
}

// Run performs the operation
func (op *Reverse) Run() (reflect.Value, *errors.Error) {
	if found, result := dispatch(op.ItemsValue, op.ItemsType); found {
		v := reflect.ValueOf(result)
		return v, nil
	}
	input := op.ItemsValue.Interface()
	sort.SliceStable(input, func(i, j int) bool { return true })
	return reflect.ValueOf(input), nil
}
