package add

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Add
const OpCode = "add"

// Add struct for defining Add operation
type Add struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

// Run performs the operations whenever is called
func (op *Add) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, result := dispatch(op.ItemsValue, op.Item, info); found {
		return reflect.ValueOf(result), nil
	}
	newItems := reflect.Append(op.ItemsValue, info.itemValue)
	return newItems, nil
}

func (op *Add) validate() (*addInfo, *errors.Error) {
	itemType := reflect.TypeOf(op.Item)
	if info := cache.get(op.ItemsType, itemType); info != nil {
		return info, nil
	}
	info := &addInfo{itemType: &itemType}
	if op.ItemsValue.Len() > 0 {
		if op.ItemsValue.Kind() != reflect.Ptr && op.Item == nil {
			return nil, errors.InvalidArgument(OpCode, "A nil value can not be added in a Stream of non-pointers values")
		}
		if op.ItemsType != itemType {
			return nil, errors.InvalidArgument(OpCode,
				"An element whose type is %s can not be added in a Stream of type %s", itemType, op.ItemsType)
		}
	}
	value := reflect.ValueOf(op.Item)
	info.itemValue = value
	cache.add(op.ItemsType, itemType, info)
	return info, nil
}
