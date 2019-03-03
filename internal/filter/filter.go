package filter

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Filter
const OpCode = "filter"

// Filter struct for operation
type Filter struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Func       interface{}
}

// Run performs the operation
func (op *Filter) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found, result := dispatch(op.ItemsValue, op.Func, info); found {
		return result, nil
	}

	if op.ItemsValue.Type().Kind() == reflect.Map {
		newItems := reflect.MakeMap(op.ItemsValue.Type())
		fn := reflect.ValueOf(op.Func)
		for _, key := range op.ItemsValue.MapKeys() {
			item := op.ItemsValue.MapIndex(key)
			argv := make([]reflect.Value, 2)
			argv[0] = item
			argv[1] = key
			if fn.Call(argv)[0].Bool() {
				newItems.SetMapIndex(key, item)
			}
		}
		return newItems, nil
	}

	newItems := reflect.MakeSlice(reflect.SliceOf(op.ItemsType), 0, 0)
	fn := reflect.ValueOf(op.Func)
	fnNum := fn.Type().NumIn()
	for index := 0; index < op.ItemsValue.Len(); index++ {
		item := op.ItemsValue.Index(index)
		argv := make([]reflect.Value, fnNum)
		argv[0] = item
		if fnNum == 2 {
			argv[1] = reflect.ValueOf(index)
		}
		if fn.Call(argv)[0].Bool() {
			newItems = reflect.Append(newItems, item)
		}
	}
	return newItems, nil
}

func (op *Filter) validate() (*filterInfo, *errors.Error) {
	item := &filterInfo{}
	fnType := reflect.TypeOf(op.Func)
	if val := cache.get(op.ItemsType, fnType); val != nil {
		return val, nil
	}
	function := reflect.ValueOf(op.Func)
	item.fnValue = function
	if function.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The filter operation requires a function as argument")
	}
	fnNumIn := function.Type().NumIn()
	if fnNumIn != 1 && fnNumIn != 2 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 1 or 2 argument")
	}
	if function.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must return 1 value")
	}
	fnOut := reflect.New(function.Type().Out(0)).Elem()
	fnIn1 := reflect.New(function.Type().In(0)).Elem()

	if fnIn1.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided function must be %s",
			op.ItemsType.String())
	}
	if fnOut.Kind() != reflect.Bool {
		return nil, errors.InvalidArgument(OpCode, "The type of the Output in the provided function must be bool")
	}
	item.fnIn1Type = fnIn1.Type()

	if fnNumIn == 2 {
		fnIn2 := reflect.New(function.Type().In(1)).Elem()

		var keyType reflect.Type
		if op.ItemsValue.Type().Kind() == reflect.Map {
			keyType = op.ItemsValue.MapKeys()[0].Type()
		} else {
			keyType = reflect.TypeOf(0)
		}

		if fnIn2.Type() != keyType {
			return nil, errors.InvalidArgument(OpCode,
				"The type of the argument 2 in the provided function must be %s",
				keyType.String())
		}
		item.fnIn2Type = fnIn2.Type()
	}

	cache.add(op.ItemsType, fnType, item)
	return item, nil
}
