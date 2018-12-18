package foreach

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
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
	if !info.hasError {
		if found := dispatch(op.ItemsValue, op.ItemsType, info); found {
			return op.ItemsValue, nil
		}
	}
	function := reflect.ValueOf(op.Func)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		item := op.ItemsValue.Index(index)
		argv := make([]reflect.Value, 1)
		argv[0] = item
		result := function.Call(argv)
		if info.hasError && !result[0].IsNil() {
			return reflect.ValueOf(nil), errors.UserError(OpCode, result[0].Interface().(error))
		}
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
	numOut := function.Type().NumOut()
	errType := reflect.TypeOf((*error)(nil)).Elem()
	if !(numOut == 0 || (numOut == 1 && fnType.Out(0).Implements(errType))) {
		return nil, errors.InvalidArgument(OpCode, "The provided function can not return any value or must return only an error")
	}
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided function "+
				"must be %s", op.ItemsType.String())
	}
	item.fnValue = reflect.ValueOf(op.Func)
	item.hasError = numOut == 1
	cache.add(op.ItemsType, fnType, item)
	return item, nil

}
