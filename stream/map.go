package stream

import (
	"fmt"
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeMap identifier for operation map
const OpCodeMap = "map"

var mapCache = make(map[reflect.Type]*mapItem)
var mapArgv = make([]reflect.Value, 1)

type streamMap struct {
	fn interface{}
}

func (m *streamMap) run(s *Stream) *Stream {
	mInfo, err := m.validateMap(s);
	if err != nil {
		s.err = err
		return s
	}
	newItems := mInfo.items
	items := reflect.ValueOf(s.items)
	if mInfo.isPtr {
		for index := 0; index < items.Len(); index++ {
			mapArgv[0] = reflect.ValueOf(reflect.ValueOf(items.Index(index).Interface()).Elem().Addr().Interface())
			result := mInfo.fnValue.Call(mapArgv)
			newItems = reflect.Append(newItems, result[0])
		}
	} else {
		for index := 0; index < items.Len(); index++ {
			mapArgv[0] = items.Index(index)
			result := mInfo.fnValue.Call(mapArgv)
			newItems = reflect.Append(newItems, result[0])
		}
	}
	s.items = newItems.Interface()
	return s
}

func (m *streamMap) validateMap(s *Stream) (*mapItem, *errors.Error) {
	item := &mapItem{}
	fnType := reflect.TypeOf(m.fn)
	fmt.Println(fnType)
	if i, ok := mapCache[fnType]; ok {
		fmt.Println("cached")
		return i, nil
	}
	if s.items == nil {
		return nil, errors.EmptyStream(OpCodeMap, "A nil Stream can not be iterated")
	}
	itemsType := reflect.TypeOf(s.items).Elem()
	item.fnValue = reflect.ValueOf(m.fn)
	if item.fnValue.Type().Kind() != reflect.Func {
		return nil, errors.InvalidArgument(OpCodeMap, "The map operation requires a function as argument")
	}
	if item.fnValue.Type().NumIn() != 1 {
		return nil, errors.InvalidArgument(OpCodeMap, "The provided function must retrieve 1 argument")
	}
	if item.fnValue.Type().NumOut() != 1 {
		return nil, errors.InvalidArgument(OpCodeMap, "The provided function must return 1 value")
	}
	fnIn := reflect.New(item.fnValue.Type().In(0)).Elem()
	if fnIn.Type() != itemsType {
		return nil, errors.InvalidArgument(OpCodeMap,
			"The type of the argument in the provided "+
				"function must be %s", itemsType.String())
	}
	item.outputType = reflect.New(item.fnValue.Type().Out(0)).Elem().Type()
	item.items = reflect.MakeSlice(reflect.SliceOf(item.outputType), 0, 0)
	item.isPtr = isPointer(s.itemsValue.Index(0))
	mapCache[fnType] = item
	return item, nil
}

type mapItem struct {
	fnValue    reflect.Value
	outputType reflect.Type
	isPtr      bool
	items      reflect.Value
}

// Map performs a mutation over all the elements in the Stream and return a new Stream
func (s *Stream) Map(fn interface{}) *Stream {
	s.operations = append(s.operations, &streamMap{fn})
	return s
}
