package groupby

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode code
const OpCode = "groupBy"

// Sort operation struct
type GroupBy struct {
	ItemsType  reflect.Type
	ItemsValue reflect.Value
	Func       interface{}
}

// Run performs the operation
func (op *GroupBy) Run() (reflect.Value, *errors.Error) {
	gInfo, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	// dispatch functions do not handle errors
	if !gInfo.hasError {
		mapType := reflect.MapOf(gInfo.fnOutputType, op.ItemsType)
		output := reflect.MakeMap(mapType)
		fn := reflect.ValueOf(op.Func)
		var argv = make([]reflect.Value, 1)
		for index := 0; index < op.ItemsValue.Len(); index++ {
			val := op.ItemsValue.Index(index)
			argv[0] = val
			result := fn.Call(argv)
			if gInfo.hasError {
				if !result[1].IsNil() {
					return reflect.ValueOf(nil), errors.UserError(OpCode, result[1].Interface().(error))
				}
			}
			keyContent := output.MapIndex(result[0])
			if keyContent.Len() == 0 {
				output.SetMapIndex(result[0], reflect.ValueOf([]interface{}{val}))
			} else {
				output.SetMapIndex(result[0], reflect.AppendSlice(keyContent, val))
			}
		}
		return output, nil
	}
	return reflect.ValueOf(nil), err
}

// Run performs the operation
func (m *GroupBy) validate() (*groupByInfo, *errors.Error) {
	item := &groupByInfo{}
	item.fnInputType = m.ItemsType
	fnType := reflect.TypeOf(m.Func)
	if val := cache.get(m.ItemsType, fnType); val != nil {
		return val, nil
	}
	item.fnValue = reflect.ValueOf(m.Func)
	if item.fnValue.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCode, "The groupBy operation requires a function as argument")
	}
	if item.fnValue.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(OpCode, "The provided function must retrieve 1 argument")
	}
	numOut := item.fnValue.Type().NumOut()
	errType := reflect.TypeOf((*error)(nil)).Elem()
	if !(numOut == 1 || (numOut == 2 && item.fnValue.Type().Out(1).Implements(errType))) {
		return nil, errors.InvalidArgument(OpCode, "The provided function must return 1 value or the second value must be an error")
	}
	fnIn := reflect.New(item.fnValue.Type().In(0)).Elem()
	if fnIn.Type() != m.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument in the provided "+
				"function must be %s", m.ItemsType.String())
	}
	item.fnOutputType = reflect.New(item.fnValue.Type().Out(0)).Elem().Type()
	if m.ItemsValue.Index(0).Kind() == reflect.Ptr {
		return nil, errors.InvalidArgument(OpCode,
			"The key of groupBy can not be a pointer")
	}
	item.hasError = numOut == 2
	cache.add(m.ItemsType, fnType, item)
	return item, nil
}
