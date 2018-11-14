package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
	reflectUtils "github.com/wesovilabs/koazee/reflection"
)

// OpCodeReduce identifier for operation reduce
const OpCodeReduce = "reduce"

type key struct {
	fn       reflect.Type
	elements reflect.Type
}

var cacheReduce = make(map[key]*CachedReduce)

type reduce struct {
	itemsValue *reflect.Value
	itemsType  reflect.Type
	fn         interface{}
}

func newReduce(itemsValue *reflect.Value, itemsType reflect.Type, fn interface{}) *reduce {
	r := &reduce{}
	r.fn = fn
	r.itemsValue = itemsValue
	r.itemsType = itemsType
	return r
}

func (op *reduce) name() string {
	return OpCodeReduce
}

var argv = make([]reflect.Value, 2)

func (op *reduce) run() output {
	if fn := reflectUtils.GetFunctionArgArgOut(op.fn); fn != nil {
		value := fn(op.itemsValue, op.fn)
		return output{value, nil}
	}
	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	function := reflect.ValueOf(op.fn)
	funcType := function.Type()
	acc := reflect.New(funcType.In(0)).Elem()
	for i := 0; i < op.itemsValue.Len(); i++ {
		argv[0] = acc
		argv[1] = op.itemsValue.Index(i)
		acc = function.Call(argv)[0]
	}
	return output{acc.Interface(), nil}
}

func (op *reduce) validate() *errors.Error {
	if _, ok := cacheReduce[key{reflect.TypeOf(op.fn), op.itemsType}]; ok {
		return nil
	}
	if op.itemsValue == nil {
		return errors.EmptyStream(op.name(), "A nil stream can not be reduced")
	}
	function := reflect.ValueOf(op.fn)
	funcType := function.Type()
	if funcType.Kind() != reflect.Func {
		return errors.InvalidArgument(op.name(), "The reduce operation requires a function as argument")
	}

	if funcType.NumIn() != 2 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 2 arguments")
	}
	if funcType.NumOut() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must return 1 value")
	}

	fnIn2 := reflect.New(funcType.In(1)).Elem()

	if fnIn2.Type() != op.itemsType {
		return errors.InvalidArgument(op.name(), "The type of the "+
			"second argument in the provided function must be %s", op.itemsType.String())
	}
	fnIn1 := reflect.New(funcType.In(0)).Elem()
	fnOut := reflect.New(funcType.Out(0)).Elem()
	if fnIn1.Type() != fnOut.Type() {
		return errors.InvalidArgument(op.name(), "The type of the first argument and "+
			"the output in the provided function must be the same")
	}
	cacheReduce[key{reflect.TypeOf(op.fn), op.itemsType}] = &CachedReduce{
		fnIn1.Type()}

	return nil
}

// Reduce
func (s stream) Reduce(fn interface{}) output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return newReduce(current.itemsValue, current.itemsType, fn).run()
}

type CachedReduce struct {
	accType reflect.Type
}
