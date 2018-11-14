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
func (op *with) run(s stream) S {
	if reflect.TypeOf(op.data).Kind() == reflect.Slice {
		itemsType := reflect.TypeOf(op.data).Elem()
		itemsValue := reflect.ValueOf(op.data)
		s.items = op.data
		s.itemsType = itemsType
		s.itemsValue = &(itemsValue)
		return s
	}

	return Error(errors.InvalidType(op.name(),
		"Unsupported type! Only arrays are permitted"))
}

func (s stream) With(data interface{}) S {
	return (&with{data}).run(s)
}

func items(data interface{}) interface{} {
	elemType := reflect.TypeOf(data).Elem()
	v := reflect.ValueOf(data)
	slice := reflect.MakeSlice(reflect.SliceOf(elemType), v.Len(), v.Len())
	for index := 0; index < v.Len(); index++ {
		slice.Index(index).Set(v.Index(index))
	}
	return slice.Interface()
}
