package drop

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation drop
const OpCode = "drop"

type Drop struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
}

// Run performs the operations whenever is called
func (op *Drop) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, result := dispatch(op.ItemsValue, op.Item, info); found {
		return reflect.ValueOf(result), nil
	}
	newItems := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, 0)
	// items := op.ItemsValue
	// indexes := make([]int, 0)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		val := op.ItemsValue.Index(index)
		if equalsValues(val, info.itemValue) {
			// items = reflect.Append(items[:index], op.ItemsValue[index:])
			// indexes = append(indexes, index)
			newItems = reflect.Append(newItems, val)
		}
	}
	return newItems, nil
}

func (op *Drop) validate() (*dropInfo, *errors.Error) {
	itemType := reflect.TypeOf(op.Item)
	if info := cache.get(op.ItemsType, itemType); info != nil {
		return info, nil
	}
	info := &dropInfo{itemType: &itemType}
	if op.ItemsValue.Len() > 0 {
		if op.ItemsValue.Kind() != reflect.Ptr && op.Item == nil {
			return nil, errors.InvalidArgument(OpCode, "A nil value can not be dropped from a Stream of non-pointers values")
		}
		if op.ItemsType != itemType {
			return nil, errors.InvalidArgument(OpCode,
				"An element whose type is %s can not be dropped from a Stream of type %s", itemType, op.ItemsType)
		}
	}
	value := reflect.ValueOf(op.Item)
	info.itemValue = value
	cache.add(op.ItemsType, itemType, info)
	return info, nil
}
