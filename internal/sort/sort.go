package sort

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
	"sort"
)

const OpCode = "sort"

type Sort struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Func       interface{}
}

func (op *Sort) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if found := dispatch(op.ItemsValue, op.Func, info); found {
		return op.ItemsValue, nil
	}
	sort.Slice(op.ItemsValue.Interface(), func(i, j int) bool {
		left := op.ItemsValue.Index(i)
		right := op.ItemsValue.Index(j)
		argv := make([]reflect.Value, 2)
		argv[0] = left
		argv[1] = right
		return info.fnValue.Call(argv)[0].Int() < 0
	})
	return op.ItemsValue, nil
}

func (op *Sort) validate() (*sortInfo, *errors.Error) {
	fnType := reflect.TypeOf(op.Func)
	if val := cache.get(op.ItemsType, fnType); val != nil {
		return val, nil
	}
	item := &sortInfo{}
	function := reflect.ValueOf(op.Func)
	item.fnValue = function
	if function.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The filter operation requires a function as argument")
	}
	if function.Type().NumIn() != 2 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 2 arguments")
	}
	if function.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must return 1 value")
	}
	fnIn1 := reflect.New(function.Type().In(0)).Elem()
	fnIn2 := reflect.New(function.Type().In(1)).Elem()
	fnOut := reflect.New(function.Type().Out(0)).Elem()
	item.fnInput1Type = function.Type().In(0)
	item.fnInput2Type = function.Type().In(1)
	item.fnOutputType = function.Type().Out(0)
	if fnIn1.Type() != fnIn2.Type() || fnIn1.Type() != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode, "The type of the both arguments must be  %s", op.ItemsType.String())
	}
	if fnOut.Type().Kind() != reflect.Int {
		return nil, errors.InvalidArgument(OpCode, "The type of the Output in the provided function must be int")
	}
	cache.add(op.ItemsType, fnType, item)
	return item, nil
}
