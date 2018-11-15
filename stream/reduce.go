package stream

import (
	"github.com/wesovilabs/koazee/errors"
	. "github.com/wesovilabs/koazee/reflection"
	"reflect"
)

// OpCodeReduce identifier for operation reduce
const OpCodeReduce = "reduce"

type key struct {
	fn       reflect.Type
	elements reflect.Type
}

var reduceCache = make(map[key]*reduceCacheItem)

var argv = make([]reflect.Value, 2)

func runReduce(itemsValue *reflect.Value, itemsType reflect.Type, fn interface{}) Output {
	function, fnIn1, err := validateReduce(itemsValue, itemsType, fn)
	if err != nil {
		return Output{nil, err}
	}
	if o := runPrimitiveReduce(itemsValue, itemsType, fn, fnIn1); o != nil {
		return *o
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
	if c, ok := reduceCache[key{reflect.TypeOf(fn), itemsType}]; ok {
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
	go func() { reduceCache[key{reflect.TypeOf(fn), itemsType}] = &reduceCacheItem{&function, &fnIn1} }()
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

type reduceCacheItem struct {
	function *reflect.Value
	accType  *reflect.Value
}

func runPrimitiveReduce(itemsValue *reflect.Value, itemsType reflect.Type, fn interface{}, fnIn1 *reflect.Value) *Output {
	if itemsType.Kind() == reflect.Int {
		switch fnIn1.Kind() {
		case reflect.Int:
			return &Output{IntIntToInt(itemsValue, fn), nil}
		case reflect.String:
			return &Output{StringIntToString(itemsValue, fn), nil}
		case reflect.Bool:
			return &Output{BoolIntToBool(itemsValue, fn), nil}
		case reflect.Float32:
			return &Output{Float32IntToFloat32(itemsValue, fn), nil}
		case reflect.Float64:
			return &Output{Float64IntToFloat64(itemsValue, fn), nil}
		}
	} else if itemsType.Kind() == reflect.String {
		switch fnIn1.Kind() {
		case reflect.Int:
			return &Output{IntStringToInt(itemsValue, fn), nil}
		case reflect.String:
			return &Output{StringStringToString(itemsValue, fn), nil}
		case reflect.Bool:
			return &Output{BoolStringToBool(itemsValue, fn), nil}
		case reflect.Float32:
			return &Output{Float32StringToFloat32(itemsValue, fn), nil}
		case reflect.Float64:
			return &Output{Float64StringToFloat64(itemsValue, fn), nil}
		}
	} else if itemsType.Kind() == reflect.Bool {
		switch fnIn1.Kind() {
		case reflect.Int:
			return &Output{IntBoolToInt(itemsValue, fn), nil}
		case reflect.String:
			return &Output{StringBoolToString(itemsValue, fn), nil}
		case reflect.Bool:
			return &Output{BoolBoolToBool(itemsValue, fn), nil}
		case reflect.Float32:
			return &Output{Float32BoolToFloat32(itemsValue, fn), nil}
		case reflect.Float64:
			return &Output{Float64BoolToFloat64(itemsValue, fn), nil}
		}
	} else if itemsType.Kind() == reflect.Float32 {
		switch fnIn1.Kind() {
		case reflect.Int:
			return &Output{IntFloat32ToInt(itemsValue, fn), nil}
		case reflect.String:
			return &Output{StringFloat32ToString(itemsValue, fn), nil}
		case reflect.Bool:
			return &Output{BoolFloat32ToBool(itemsValue, fn), nil}
		case reflect.Float32:
			return &Output{Float32Float32ToFloat32(itemsValue, fn), nil}
		case reflect.Float64:
			return &Output{Float64Float32ToFloat64(itemsValue, fn), nil}
		}
	}
	return nil
}
