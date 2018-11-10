package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeReduce identifier for operation reduce
const OpCodeReduce = "reduce"

type reduce struct {
	items interface{}
	fn    interface{}
}

func (op *reduce) name() string {
	return OpCodeReduce
}

func (op *reduce) run() output {

	if err := op.validate(); err != nil {
		return output{nil, err}
	}
	function := reflect.ValueOf(op.fn)
	items := reflect.ValueOf(op.items)
	acc := reflect.New(function.Type().In(0)).Elem()
	for index := 0; index < items.Len(); index++ {
		argv := make([]reflect.Value, 2)
		argv[0] = acc
		argv[1] = items.Index(index)
		result := function.Call(argv)
		acc = reflect.ValueOf(result[0].Interface())
	}
	logger.DebugInfo("%s %v -> %v", op.name(), op.items, acc.Interface())
	return output{acc.Interface(), nil}
}

func (op *reduce) validate() *errors.Error {
	if op.items == nil {
		return errors.ItemsNil(op.name(), "A nil stream can not be reduced")
	}
	itemsType := reflect.TypeOf(op.items).Elem()
	function := reflect.ValueOf(op.fn)
	if function.Type().Kind() != reflect.Func {
		return errors.InvalidArgument(op.name(), "The filter operation requires a function as argument")
	}
	if function.Type().NumIn() != 2 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 2 arguments")
	}
	if function.Type().NumOut() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must return 1 value")
	}
	fnIn1 := reflect.New(function.Type().In(0)).Elem()
	fnIn2 := reflect.New(function.Type().In(1)).Elem()
	fnOut := reflect.New(function.Type().Out(0)).Elem()
	if fnIn2.Type() != itemsType {
		return errors.InvalidArgument(op.name(), "The type of the "+
			"second argument in the provided function must be %s", itemsType.String())
	}
	if fnIn1.Type() != fnOut.Type() {
		return errors.InvalidArgument(op.name(), "The type of the first argument and "+
			"the output in the provided function must be the same")
	}
	return nil
}

// Reduce
func (s stream) Reduce(fn interface{}) output {
	current := s.run()
	if current.err != nil {
		return output{nil, current.err}
	}
	return (&reduce{current.items, fn}).run()
}
