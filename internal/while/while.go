package while

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode identifier for operation Filter
const OpCode = "while"

// While struct for the operation
type While struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value

	SentinelFunc interface{}
	DoFunc       interface{}
}

// Run performs the operation
func (op *While) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	if found, result := dispatch(op.ItemsValue, op.SentinelFunc, op.DoFunc, info); found {
		return reflect.ValueOf(result), nil
	}

	sentinelFunc := reflect.ValueOf(op.SentinelFunc)
	doFunc := reflect.ValueOf(op.DoFunc)
	for index := 0; index < op.ItemsValue.Len(); index++ {
		item := op.ItemsValue.Index(index)
		argv := make([]reflect.Value, 1)
		argv[0] = item
		if sentinelFunc.Call(argv)[0].Bool() {
			doFunc.Call(argv)
		}
	}
	return op.ItemsValue, nil
}

func (op *While) validate() (*whileInfo, *errors.Error) {
	item := &whileInfo{}

	// SentinelFunc
	sfnType := reflect.TypeOf(op.SentinelFunc)
	dfnType := reflect.TypeOf(op.DoFunc)
	if val := cache.get(op.ItemsType, sfnType, dfnType); val != nil {
		return val, nil
	}

	sfunction := reflect.ValueOf(op.SentinelFunc)
	item.fnValue = sfunction
	if sfunction.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The while operation requires a function as first argument")
	}
	if sfunction.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 1 argument")
	}
	if sfunction.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must return 1 value")
	}
	sfnOut := reflect.New(sfunction.Type().Out(0)).Elem()
	sfnIn := reflect.New(sfunction.Type().In(0)).Elem()

	if sfnIn.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided function must be %s",
			op.ItemsType.String())
	}
	if sfnOut.Kind() != reflect.Bool {
		return nil, errors.InvalidArgument(OpCode, "The type of the Output in the provided function must be bool")
	}
	item.fnInputType = sfnIn.Type()

	// DoFunc
	doFn := reflect.ValueOf(op.DoFunc)
	item.doFnValue = doFn
	if doFn.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The while operation requires a function as second argument")
	}
	if doFn.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided DoFunc must retrieve 1 argument")
	}
	if doFn.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided DoFunc must return 1 value")
	}
	doFnOut := reflect.New(doFn.Type().Out(0)).Elem()
	doFnIn := reflect.New(doFn.Type().In(0)).Elem()
	if doFnIn.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided DoFunc "+
				"must be %s", op.ItemsType.String())
	}
	if doFnOut.Type() != doFnIn.Type() {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the Output in the provided DoFunc "+
				"must be %s", op.ItemsType.String())
	}

	item.doFnInputType = doFnIn.Type()
	item.doFnOutputType = doFnOut.Type()

	cache.add(op.ItemsType, sfnType, dfnType, item)

	return item, nil
}
