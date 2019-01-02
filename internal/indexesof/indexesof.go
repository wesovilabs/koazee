package indexesof

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation IndexesOf
const OpCode = "indexesOf"

// IndexesOf struct for handling the operation
type IndexesOf struct {
	ItemsType reflect.Type
	Items     reflect.Value
	Element   interface{}
}

// Run performs the operation
func (op *IndexesOf) Run() ([]int, *errors.Error) {
	if op.Items.Len() == 0 {
		return []int{}, errors.EmptyStream(OpCode, "stream is nil")
	}
	if err := op.validate(); err != nil {
		return []int{}, err
	}
	ret := make([]int, 0)
	element := reflect.ValueOf(op.Element)
	if found, result := dispatch(op.Items, op.Element, op.ItemsType); found {
		return result, nil
	}
	for index := 0; index < op.Items.Len(); index++ {
		if equalsValues(op.Items.Index(index), element) {
			ret = append(ret, index)
		}
	}
	if len(ret) > 0 {
		return ret, nil
	} else {
		return []int{}, errors.InvalidArgument(OpCode, "not found")
	}
}

func (op *IndexesOf) validate() *errors.Error {
	info := &indexesOfInfo{}
	elementType := reflect.TypeOf(op.Element)
	if val := cache.get(op.ItemsType, elementType); val != nil {
		return nil
	}
	if op.ItemsType.Kind() != reflect.Ptr && op.Element == nil {
		return errors.InvalidArgument(OpCode, "It can not be checked if an array of non-pointers indexesOf a nil value")
	}
	if elementType != op.ItemsType {
		return errors.InvalidArgument(OpCode,
			"The Stream indexesOf elements of type %s and the passed argument has type %s", op.ItemsType, elementType)
	}
	cache.add(op.ItemsType, elementType, info)
	return nil
}
