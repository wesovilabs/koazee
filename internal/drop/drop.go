package drop

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation drop
const OpCode = "drop"

// Option defines drop option
type Option int

// Options
const (
	None Option = iota
	Left
	Right
)

// Drop struct for operation
type Drop struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	Item       interface{}
	Option     Option
}

// Run performs the operations whenever is called
func (op *Drop) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, result := dispatch(op.ItemsValue, op.Item, info, op.Option); found {
		return reflect.ValueOf(result), nil
	}

	isFounded := false
	newItems := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, 0)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		val := op.ItemsValue.Index(index)
		switch op.Option {
		case None:
			if equalsValues(val, info.itemValue) == false {
				isFounded = true
				newItems = reflect.Append(newItems, val)
			}
		case Left:
			if isFounded || equalsValues(val, info.itemValue) {
				isFounded = true
				newItems = reflect.Append(newItems, val)
			}
		case Right:
			if equalsValues(val, info.itemValue) {
				isFounded = true
				newItems = reflect.Append(newItems, val)
			} else if isFounded == false {
				newItems = reflect.Append(newItems, val)
			}
		}
	}

	if isFounded {
		return newItems, nil
	}
	return op.ItemsValue, nil
}

func (op *Drop) validate() (*dropInfo, *errors.Error) {
	itemType := reflect.TypeOf(op.Item)
	if info := cache.get(op.ItemsType, itemType); info != nil {
		if equalsValues(info.itemValue, reflect.ValueOf(op.Item)) {
			return info, nil
		}
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
