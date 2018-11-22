package contains

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation Contains
const OpCode = "contains"

type Contains struct {
	ItemsType reflect.Type
	Items     *reflect.Value
	Element   interface{}
}

func (op *Contains) Run() (bool, *errors.Error) {
	if op.Items == nil {
		return false, nil
	}
	if err := op.validate(); err != nil {
		return false, err
	}
	element := reflect.ValueOf(op.Element)
	if found, result := dispatch(op.Items, op.Element, op.ItemsType); found {
		return result, nil
	}
	for index := 0; index < op.Items.Len(); index++ {
		if equalsValues(op.Items.Index(index), element) {
			return true, nil
		}
	}
	return false, nil

}

func (op *Contains) validate() *errors.Error {
	info := &containsInfo{}
	elementType := reflect.TypeOf(op.Element)
	if val := cache.get(op.ItemsType, elementType); val != nil {
		return nil
	}
	if op.ItemsType.Kind() != reflect.Ptr && op.Element == nil {
		return errors.InvalidArgument(OpCode, "It can not be checked if an array of non-pointers contains a nil value")
	}
	if elementType != op.ItemsType {
		return errors.InvalidArgument(OpCode,
			"The Stream contains elements of type %s and the passed argument has type %s", op.ItemsType, elementType)
	}
	cache.add(op.ItemsType, elementType, info)
	return nil
}
