package duplicates

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation RemoveDuplicates
const OpCode = "removeDuplicates"

type RemoveDuplicates struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type

}

func (op *RemoveDuplicates) Run() (reflect.Value, *errors.Error) {

	if err := op.validate(); err != nil {
		return reflect.ValueOf(nil), err
	}
	keys := op.prepareMapWithKeys()
	if op.ItemsValue.Index(0).Kind() == reflect.Ptr {
		newItems := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, 0)
		for index := 0; index < op.ItemsValue.Len(); index++ {
			val := op.ItemsValue.Index(index).Elem().Interface()
			if keys[val] == 0 {
				continue
			}
			keys[val] = 0
			newItems = reflect.Append(newItems, op.ItemsValue.Index(index))
		}
		return newItems, nil
	}
	newItems := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, 0)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		val := op.ItemsValue.Index(index).Interface()
		if keys[val] == 0 {
			continue
		}
		keys[val] = 0
		newItems = reflect.Append(newItems, op.ItemsValue.Index(index))
	}
	return newItems, nil
}

func (op *RemoveDuplicates) validate() *errors.Error {
	return nil
}

func (op *RemoveDuplicates) prepareMapWithKeys() map[interface{}]int {
	keys := make(map[interface{}]int)

	if op.ItemsValue.Index(0).Kind() == reflect.Ptr {
		for index := 0; index < op.ItemsValue.Len(); index++ {
			v := op.ItemsValue.Index(index).Elem().Interface()
			keys[v]++
		}
	}else{
		for index := 0; index < op.ItemsValue.Len(); index++ {
			v := op.ItemsValue.Index(index).Interface()
			keys[v]++
		}
	}
	return keys
}