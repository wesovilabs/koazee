package _map

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

const opCode = "map"

type Map struct {
	ItemsType  reflect.Type
	ItemsValue *reflect.Value
	Func       interface{}
}

func (m *Map) Run() (interface{}, *errors.Error) {
	mInfo, err := m.validate()
	if err != nil {
		return nil, err
	}
	if found, result := dispatch(m.ItemsValue, m.Func, mInfo); found {
		return result, nil
	}
	newItems := mInfo.items
	var argv = make([]reflect.Value, 1)
	if mInfo.isPtr {
		for index := 0; index < m.ItemsValue.Len(); index++ {
			argv[0] = reflect.ValueOf(reflect.ValueOf(m.ItemsValue.Index(index).Interface()).Elem().Addr().Interface())
			result := mInfo.fnValue.Call(argv)
			newItems = reflect.Append(newItems, result[0])
		}
	} else {
		for index := 0; index < m.ItemsValue.Len(); index++ {
			argv[0] = m.ItemsValue.Index(index)
			result := mInfo.fnValue.Call(argv)
			newItems = reflect.Append(newItems, result[0])
		}
	}
	return newItems.Interface(), nil
}

func (m *Map) validate() (*mapInfo, *errors.Error) {
	item := &mapInfo{}
	item.fnInputType = m.ItemsType
	fnType := reflect.TypeOf(m.Func)
	if val := cache.get(m.ItemsType, fnType); val != nil {
		return val, nil
	}
	if m.ItemsValue == nil {
		return nil, errors.EmptyStream(opCode, "A nil Stream can not be iterated")
	}
	item.fnValue = reflect.ValueOf(m.Func)
	if item.fnValue.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(opCode, "The map operation requires a function as argument")
	}
	if item.fnValue.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(opCode, "The provided function must retrieve 1 argument")
	}
	if item.fnValue.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(opCode, "The provided function must return 1 value")
	}
	fnIn := reflect.New(item.fnValue.Type().In(0)).Elem()
	if fnIn.Type() != m.ItemsType {
		return nil, errors.InvalidArgument(opCode,
			"The type of the argument in the provided "+
				"function must be %s", m.ItemsType.String())
	}
	item.fnOutputType = reflect.New(item.fnValue.Type().Out(0)).Elem().Type()
	// item.items = reflect.MakeSlice(reflect.SliceOf(item.fnOutputType), 0, 0)
	item.isPtr = m.ItemsValue.Index(0).Kind() == reflect.Ptr
	cache.add(m.ItemsType, fnType, item)
	return item, nil
}
