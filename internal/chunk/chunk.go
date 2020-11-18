package chunk

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode code
const OpCode = "chunk"

// Chunk operation struct
type Chunk struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Size       uint
}

// Run performs the operation
func (op *Chunk) Run() (reflect.Value, *errors.Error) {
	if found, result := dispatch(op.ItemsValue, op.ItemsType, op.Size); found {
		return result, nil
	}

	len := op.ItemsValue.Len()
	size := int(op.Size)
	output := reflect.MakeSlice(reflect.SliceOf(reflect.SliceOf(op.ItemsType)), 0, outerCapacity(len, size))
	for index := 0; index < len; index += size {
		rangeLast := chunkRangeLast(index, size, len)
		output = reflect.Append(output, op.ItemsValue.Slice3(index, rangeLast, rangeLast))
	}
	return output, nil
}
