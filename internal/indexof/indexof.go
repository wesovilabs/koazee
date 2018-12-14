package indexof

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation IndexOf
const OpCode = "indexOf"
const InvalidIndex = -1

type IndexOf struct {
	ItemsType reflect.Type
	Items     reflect.Value
	Element   interface{}
}

func (op *IndexOf) Run() (int, *errors.Error) {
	if op.Items.Len() == 0 {
		return InvalidIndex, errors.EmptyStream(OpCode, "stream is nil")
	}
	if err := op.validate(); err != nil {
		return InvalidIndex, err
	}

	element := reflect.ValueOf(op.Element)
	if found, result := dispatch(op.Items, op.Element, op.ItemsType); found {
		return result, nil
	}

	for index := 0; index < op.Items.Len(); index++ {
		if equalsValues(op.Items.Index(index), element) {
			return index, nil
		}
	}

	return InvalidIndex, errors.InvalidArgument(OpCode, "not found")
}

func (op *IndexOf) validate() *errors.Error {
	info := &indexOfInfo{}
	elementType := reflect.TypeOf(op.Element)
	if val := cache.get(op.ItemsType, elementType); val != nil {
		return nil
	}
	if op.ItemsType.Kind() != reflect.Ptr && op.Element == nil {
		return errors.InvalidArgument(OpCode, "It can not be checked if an array of non-pointers indexOf a nil value")
	}
	if elementType != op.ItemsType {
		return errors.InvalidArgument(OpCode,
			"The Stream indexOf elements of type %s and the passed argument has type %s", op.ItemsType, elementType)
	}
	cache.add(op.ItemsType, elementType, info)
	return nil
}
