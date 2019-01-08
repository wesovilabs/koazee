package reduce

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCode code for operation reduce
const OpCode = "reduce"

// Reduce struct for this operation
type Reduce struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Func       interface{}
}

// Run performs the operation
func (r *Reduce) Run() (reflect.Value, *errors.Error) {
	info, err := r.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	if !info.hasError {
		if found, result := dispatch(r.ItemsValue, r.Func, info); found {
			v := reflect.ValueOf(result)
			return v, nil
		}
	}
	var argv = make([]reflect.Value, 2)
	acc := reflect.New(info.fnIn1Type).Elem()
	for i := 0; i < r.ItemsValue.Len(); i++ {
		argv[0] = acc
		argv[1] = r.ItemsValue.Index(i)
		output := info.fnValue.Call(argv)
		acc = output[0]
		if info.hasError && !output[1].IsNil() {
			return reflect.ValueOf(nil), errors.UserError(OpCode, output[1].Interface().(error))
		}
	}
	return acc, nil
}

func (r *Reduce) validate() (*reduceInfo, *errors.Error) {
	if r.ItemsValue == reflect.ValueOf(nil) {
		return nil, errors.EmptyStream(OpCode, "A nil Stream can not be reduced")
	}
	fnType := reflect.TypeOf(r.Func)
	if reduceInfo := cache.get(r.ItemsType, fnType); reduceInfo != nil {
		return reduceInfo, nil
	}
	info := &reduceInfo{}
	info.fnType = fnType
	info.fnValue = reflect.ValueOf(r.Func)

	if fnType.Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The reduce operation requires a function as argument")
	}
	if fnType.NumIn() != 2 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 2 arguments")
	}
	numOut := fnType.NumOut()
	errType := reflect.TypeOf((*error)(nil)).Elem()
	if !(numOut == 1 || (numOut == 2 && fnType.Out(1).Implements(errType))) {
		return nil, errors.InvalidArgument(OpCode, "The provided function must return 1 value or the second value must be an error")
	}
	fnIn2 := reflect.New(fnType.In(1)).Elem()
	if fnIn2.Type() != r.ItemsType {
		return nil, errors.InvalidArgument(OpCode, "The type of the "+
			"second argument in the provided function must be %s", r.ItemsType.String())
	}
	info.fnIn1Type = fnType.In(0)
	info.fnIn2Type = fnType.In(1)
	info.fnOutType = fnType.Out(0)
	info.hasError = numOut == 2
	if info.fnIn1Type != info.fnOutType {
		return nil, errors.InvalidArgument(OpCode, "The type of the first argument and "+
			"the Output in the provided function must be the same")
	}
	cache.add(r.ItemsType, fnType, info)
	return info, nil
}
