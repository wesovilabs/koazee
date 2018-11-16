package reduce

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

const opCode = "reduce"

type Reduce struct {
	itemsType  reflect.Type
	itemsValue *reflect.Value
	fn         interface{}
}

func (r *Reduce) Run() (interface{}, *errors.Error) {
	info, err := r.validate()
	if err != nil {
		return nil, err
	}
	if result := dispatch(r.itemsValue, r.fn, info); result != nil {
		return result, nil
	}
	var argv = make([]reflect.Value, 2)
	acc := reflect.New(info.fnIn1Type).Elem()
	for i := 0; i < r.itemsValue.Len(); i++ {
		argv[0] = acc
		argv[1] = r.itemsValue.Index(i)
		output := info.fnValue.Call(argv)
		acc = output[0]
	}
	return acc.Interface(), nil
}

func (r *Reduce) validate() (*reduceInfo, *errors.Error) {
	fnType := reflect.TypeOf(r.fn)
	if reduceInfo, ok := cache[fnType]; ok {
		return reduceInfo, nil
	}
	info := &reduceInfo{}
	info.fnValue = reflect.ValueOf(r.fn)
	if r.itemsValue == nil {
		return nil, errors.EmptyStream(opCode, "A nil Stream can not be reduced")
	}
	if fnType.Kind() != reflect.Func {
		return nil, errors.InvalidArgument(opCode, "The reduce operation requires a function as argument")
	}
	if fnType.NumIn() != 2 {
		return nil, errors.InvalidArgument(opCode, "The provided function must retrieve 2 arguments")
	}
	if fnType.NumOut() != 1 {
		return nil, errors.InvalidArgument(opCode, "The provided function must return 1 value")
	}
	fnIn2 := reflect.New(fnType.In(1)).Elem()
	if fnIn2.Type() != r.itemsType {
		return nil, errors.InvalidArgument(opCode, "The type of the "+
			"second argument in the provided function must be %s", r.itemsType.String())
	}
	info.fnIn1Type = fnType.In(0)
	info.fnIn2Type = fnType.In(1)
	info.fnOutType = fnType.Out(0)
	if info.fnIn1Type != info.fnOutType {
		return nil, errors.InvalidArgument(opCode, "The type of the first argument and "+
			"the Output in the provided function must be the same")
	}
	cache[fnType] = info
	return info, nil
}


