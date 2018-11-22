package add

import (
	"fmt"
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Add
const OpCode = "add"

// Add struct for defining Add operation
type Add struct {
	ItemsValue *reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

// Run performs the operations whenever is called
func (op *Add) Run() (interface{}, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return nil, err
	}
	/**
	if found, items := dispatch(op.ItemsValue, info.itemValue, info); found {
		return items, nil
	}
	**/
	if op.ItemsValue == nil {
		return []interface{}{info.itemValue}, nil
	}
	return reflect.Append(*op.ItemsValue, info.itemValue), nil
}

func (op *Add) validate() (*addInfo, *errors.Error) {
	itemType := reflect.TypeOf(op.Item)
	if info := cache.get(op.ItemsType, itemType); info != nil {
		return info, nil
	}
	fmt.Println("cache...")
	info := &addInfo{itemType: &itemType}
	if op.ItemsValue != nil {
		if op.ItemsValue.Kind() != reflect.Ptr && op.Item == nil {
			return nil, errors.InvalidArgument(OpCode, "A nil value can not be added in a Stream of non-pointers values")
		}
		if op.ItemsType != itemType {
			return nil, errors.InvalidArgument(OpCode,
				"An element whose type is %s can not be added in a Stream of type %s", itemType, op.ItemsType)
		}
	}
	info.itemValue = reflect.ValueOf(op.Item)
	cache.add(op.ItemsType, itemType, info)
	return info, nil
}
