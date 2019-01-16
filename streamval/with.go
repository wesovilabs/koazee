package streamval

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCodeWith identifier for operation with
const OpCodeWith = "with"

// with struct for defining add operation
type with struct {
	data interface{}
}

func (op *with) name() string {
	return OpCodeWith
}

// Run performs the operations whenever is called
func (op *with) run(s StreamVal) StreamVal {
	if reflect.TypeOf(op.data).Kind() == reflect.Slice {
		itemsType := reflect.TypeOf(op.data).Elem()
		stream:=koazee.StreamOf(op.data)
		stream.WithOperations(s)
		return s
	}

	return Error(errors.InvalidType(op.name(),
		"Unsupported type! Only arrays are permitted"))
}

// With load the provided elements in the stream
func (s StreamVal) With(data interface{}) StreamVal {
	return (&with{data}).run(s)
}
