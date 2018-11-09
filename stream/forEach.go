package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OperationFilterIdentifier identifier for operation filter
const OperationForEachIdentifier = ":forEach"

type forEach struct {
	fn interface{}
}

func (op *forEach) name() string {
	return OperationForEachIdentifier
}

func (op *forEach) run(s *stream) *stream {
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

func (op *forEach) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "You can not iterate over a nil stream")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	function := reflect.ValueOf(op.fn)
	if function.Type().NumIn() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 1 argument")
	}
	if function.Type().NumOut() != 0 {
		return errors.InvalidArgument(op.name(), "The provided function can not return any value")
	}
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != itemsType {
		return errors.InvalidArgument(op.name(), "The type of the argument in the provided function must be %s", itemsType.String())
	}
	return nil
}

// ForEach executes the provided function for all the elements in the stream
func (s stream) ForEach(fn interface{}) S {
	s.operations = append(s.operations, &forEach{fn})
	return s
}
