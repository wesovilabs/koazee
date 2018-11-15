package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeForEach identifier for operation forEach
const OpCodeForEach = "forEach"

type forEach struct {
	fn interface{}
}

func (op *forEach) name() string {
	return OpCodeForEach
}

func (op *forEach) run(s *Stream) *Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	function := reflect.ValueOf(op.fn)
	v := reflect.ValueOf(s.items)
	for index := 0; index < v.Len(); index++ {
		item := v.Index(index)
		argv := make([]reflect.Value, 1)
		argv[0] = item
		function.Call(argv)
	}
	logger.DebugInfo("%s", op.name())
	return s
}

func (op *forEach) validate(s *Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "A nil Stream can not be used to perform ForEach operation")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	function := reflect.ValueOf(op.fn)
	if function.Type().Kind() != reflect.Func {
		return errors.InvalidArgument(op.name(), "The forEach operation requires a function as argument")
	}
	if function.Type().NumIn() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 1 argument")
	}
	if function.Type().NumOut() != 0 {
		return errors.InvalidArgument(op.name(), "The provided function can not return any value")
	}
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != itemsType {
		return errors.InvalidArgument(op.name(),
			"The type of the argument in the provided function "+
				"must be %s", itemsType.String())
	}
	return nil
}

// ForEach executes the provided function for all the elements in the Stream
func (s *Stream) ForEach(fn interface{}) *Stream {
	s.operations = append(s.operations, &forEach{fn})
	return s
}
