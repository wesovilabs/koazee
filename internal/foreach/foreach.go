package foreach

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation Filter
const OpCode = "forEach"

// ForEach struct for the operation
type ForEach struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Func       interface{}
}

// Run performs the operation
func (op *ForEach) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found := dispatch(op.ItemsValue, op.ItemsType, info); found {
		return op.ItemsValue, nil
	}
	function := reflect.ValueOf(op.Func)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		item := op.ItemsValue.Index(index)
		argv := make([]reflect.Value, 1)
		argv[0] = item
		function.Call(argv)
	}
	return op.ItemsValue, nil
}

func (op *ForEach) validate() (*forEachInfo, *errors.Error) {
	item := &forEachInfo{}
	fnType := reflect.TypeOf(op.Func)
	if val := cache.get(op.ItemsType, fnType); val != nil {
		return val, nil
	}

	function := reflect.ValueOf(op.Func)
	if function.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The forEach operation requires a function as argument")
	}
	if function.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 1 argument")
	}
	if function.Type().NumOut() != 0 {
		return nil, errors.InvalidArgument(OpCode, "The provided function can not return any value")
	}
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided function "+
				"must be %s", op.ItemsType.String())
	}
	item.fnValue = reflect.ValueOf(op.Func)
	cache.add(op.ItemsType, fnType, item)
	return item, nil

}
