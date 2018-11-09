package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OperationMapIdentifier identifier for operation map
const OperationMapIdentifier = ":map"

type streamMap struct {
	fn interface{}
}

func (op *streamMap) name() string {
	return OperationMapIdentifier
}

func (op *streamMap) run(s *stream) *stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}
	function := reflect.ValueOf(op.fn)
	outputType := reflect.New(function.Type().Out(0)).Elem().Type()
	newItems := reflect.MakeSlice(reflect.SliceOf(outputType), 0, 0)
	v := reflect.ValueOf(s.items)
	for index := 0; index < v.Len(); index++ {
		argv := make([]reflect.Value, 1)
		if isPointer(v.Index(index)) {
			argv[0] = reflect.ValueOf(reflect.ValueOf(v.Index(index).Interface()).Elem().Addr().Interface())
		} else {
			argv[0] = v.Index(index)
		}
		result := function.Call(argv)
		newItems = reflect.Append(newItems, result[0])
	}
	return s
}

func (op *streamMap) validate(s *stream) *errors.Error {
	if s.items == nil {
		return errors.ItemsNil(op.name(), "You can not iterate over a nil stream")
	}
	itemsType := reflect.TypeOf(s.items)
	function := reflect.ValueOf(op.fn)
	if function.Type().NumIn() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must retrieve 1 argument")
	}
	if function.Type().NumOut() != 1 {
		return errors.InvalidArgument(op.name(), "The provided function must return 1 value")
	}
	fnIn := reflect.New(function.Type().In(0)).Elem()
	if fnIn.Type() != itemsType {
		return errors.InvalidArgument(op.name(), "The type of the argument in the provided function must be %s", itemsType.String())
	}
	return nil
}

// Map performs a mutation over all the elements in the stream and return a new stream
func (s *stream) Map(fn interface{}) S {
	s.operations = append(s.operations, &streamMap{fn})
	return s
}

