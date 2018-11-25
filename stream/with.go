package stream

import (
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
func (op *with) run(s Stream) Stream {
	if reflect.TypeOf(op.data).Kind() == reflect.Slice {
		itemsType := reflect.TypeOf(op.data).Elem()
		s.items = op.data
		s.itemsType = itemsType
		s.itemsValue = reflect.ValueOf(op.data)
		s.itemsLen = s.itemsValue.Len()
		return s
	}

	return Error(errors.InvalidType(op.name(),
		"Unsupported type! Only arrays are permitted"))
}

func (s Stream) With(data interface{}) Stream {
	return (&with{data}).run(s)
}
