package stream

import (
	"github.com/wesovilabs/koazee/errors"
	reflectUtils "github.com/wesovilabs/koazee/reflection"
	"reflect"
)

// OpCodeReduce identifier for operation reduce
const OpCodeReduce = "reduce"

type key struct {
	fn       reflect.Type
	elements reflect.Type
}

var cacheReduce = make(map[key]*CachedReduce)

var argv = make([]reflect.Value, 2)

func runReduce(itemsValue *reflect.Value, itemsType reflect.Type, fn interface{}) Output {
	function, fnIn1, err := validateReduce(itemsValue, itemsType, fn)
	if err != nil {
		return Output{nil, err}
	}
	if itemsType.Kind() == reflect.Int && fnIn1.Kind() == reflect.Int {
		return Output{reflectUtils.IntIntToInt(itemsValue, fn), nil}
	}
	acc := *fnIn1
	for i := 0; i < itemsValue.Len(); i++ {
		argv[0] = acc
		argv[1] = itemsValue.Index(i)
		acc = function.Call(argv)[0]
	}
	return Output{acc.Interface(), nil}
}

func validateReduce(itemsValue *reflect.Value, itemsType reflect.Type, fn interface{}) (
	*reflect.Value, *reflect.Value, *errors.Error) {
	if c, ok := cacheReduce[key{reflect.TypeOf(fn), itemsType}]; ok {
		return c.function, c.accType, nil
	}
	if itemsValue == nil {
		return nil, nil, errors.EmptyStream(OpCodeReduce, "A nil Stream can not be reduced")
	}
	function := reflect.ValueOf(fn)
	funcType := function.Type()
	if funcType.Kind() != reflect.Func {
		return nil, nil, errors.InvalidArgument(OpCodeReduce, "The reduce operation requires a function as argument")
	}

	if funcType.NumIn() != 2 {
		return nil, nil, errors.InvalidArgument(OpCodeReduce, "The provided function must retrieve 2 arguments")
	}
	if funcType.NumOut() != 1 {
		return nil, nil, errors.InvalidArgument(OpCodeReduce, "The provided function must return 1 value")
	}

	fnIn2 := reflect.New(funcType.In(1)).Elem()

	if fnIn2.Type() != itemsType {
		return nil, nil, errors.InvalidArgument(OpCodeReduce, "The type of the "+
			"second argument in the provided function must be %s", itemsType.String())
	}
	fnIn1 := reflect.New(funcType.In(0)).Elem()
	fnOut := reflect.New(funcType.Out(0)).Elem()
	if fnIn1.Type() != fnOut.Type() {
		return nil, nil, errors.InvalidArgument(OpCodeReduce, "The type of the first argument and "+
			"the Output in the provided function must be the same")
	}
	cacheReduce[key{reflect.TypeOf(fn), itemsType}] = &CachedReduce{&function, &fnIn1}

	return &function, &fnIn1, nil
}

// Reduce
func (s *Stream) Reduce(fn interface{}) Output {
	current := s.run()
	if current.err != nil {
		return Output{nil, current.err}
	}
	return runReduce(current.itemsValue, current.itemsType, fn)
}

type CachedReduce struct {
	function *reflect.Value
	accType  *reflect.Value
}
