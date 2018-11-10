package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/utils"

	"github.com/wesovilabs/koazee/logger"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeWith identifier for operation with
const OpCodeWith = "with"

// with struct for defining add operation
type with struct {
	data interface{}
}

func (op *with) name() string {
	return OpCodeWith
}

// Run performs the operations whenever is called
func (op *with) run(s stream) stream {
	nature := utils.NatureOf(op.data)
	switch nature {
	case utils.NatureArray:
		out := items(op.data)
		logger.DebugInfo("%s  %v", op.name(), out)
		s.items = out

	default:
		s.err = errors.InvalidType(op.name(), "Unsupported type! Only arrays are permitted")
		return s
	}

	return s
}

func (s stream) With(data interface{}) S {
	return (&with{data}).run(s)
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
