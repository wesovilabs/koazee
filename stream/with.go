package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

const OperationWithIdentifier = ":with"

// with struct for defining add operation
type with struct {
	data    interface{}
	traceID string
}

func (op *with) name() string {
	return OperationWithIdentifier
}

// Run performs the operations whenever is called
func (op *with) run(s stream) stream {
	nature := natureOf(op.data)
	switch nature {
	case natureArray:
		out := items(op.data)
		logger.DebugInfo(op.traceID, "%s  %v", op.name(), out)
		s.items = out

	default:
		s.err = errors.InvalidType(":load", "Unsupported type! Only arrays are permitted")
		return s
	}

	return s
}

func (s stream) With(data interface{}) S {
	return (&with{data, s.traceID}).run(s)
}

func items(data interface{}) interface{} {
	elemType := reflect.TypeOf(data).Elem()
	v := reflect.ValueOf(data)
	slice := reflect.MakeSlice(reflect.SliceOf(elemType), v.Len(), v.Len())
	for index := 0; index < v.Len(); index++ {
		slice.Index(index).Set(v.Index(index))
	}
	return slice.Interface()
}
