package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
	. "github.com/wesovilabs/koazee/reflection"
)

// OpCodeMap identifier for operation map
const OpCodeMap = "map"

var mapCache = make(map[reflect.Type]*mapItem)
var mapArgv = make([]reflect.Value, 1)

type streamMap struct {
	fn interface{}
}

func (m *streamMap) run(s *Stream) *Stream {
	mInfo, err := m.validateMap(s)
	if err != nil {
		s.err = err
		return s
	}
	if output := runPrimitiveMap(s.itemsValue, m.fn, s.itemsType, mInfo.outputType); output != nil {
		s.items = output
		return s
	}
	newItems := mInfo.items
	if mInfo.isPtr {
		for index := 0; index < s.itemsValue.Len(); index++ {
			mapArgv[0] = reflect.ValueOf(reflect.ValueOf(s.itemsValue.Index(index).Interface()).Elem().Addr().Interface())
			result := mInfo.fnValue.Call(mapArgv)
			newItems = reflect.Append(newItems, result[0])
		}
	} else {
		for index := 0; index < s.itemsValue.Len(); index++ {
			mapArgv[0] = s.itemsValue.Index(index)
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
	if i, ok := mapCache[fnType]; ok {
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

func runPrimitiveMap(itemsValue *reflect.Value, fn interface{}, fnIn, fnOut reflect.Type) interface{} {
	if fnIn.Kind() == reflect.Int {
		switch fnOut.Kind() {
		case reflect.Int:
			return IntToInt(itemsValue, fn)
		case reflect.String:
			return IntToString(itemsValue, fn)
		case reflect.Bool:
			return IntToBool(itemsValue, fn)
		case reflect.Float32:
			return IntToFloat32(itemsValue, fn)
		case reflect.Float64:
			return IntToFloat64(itemsValue, fn)
		}
	} else if fnIn.Kind() == reflect.String {
		switch fnOut.Kind() {
		case reflect.Int:
			return StringToInt(itemsValue, fn)
		case reflect.String:
			return StringToString(itemsValue.Interface().([]string), fn)
		case reflect.Bool:
			return StringToBool(itemsValue, fn)
		case reflect.Float32:
			return StringToFloat32(itemsValue, fn)
		case reflect.Float64:
			return StringToFloat64(itemsValue, fn)
		}
	} else if fnIn.Kind() == reflect.Bool {
		switch fnOut.Kind() {
		case reflect.Int:
			return BoolToInt(itemsValue, fn)
		case reflect.String:
			return BoolToString(itemsValue, fn)
		case reflect.Bool:
			return BoolToBool(itemsValue, fn)
		case reflect.Float32:
			return BoolToFloat32(itemsValue, fn)
		case reflect.Float64:
			return BoolToFloat64(itemsValue, fn)
		}
	} else if fnIn.Kind() == reflect.Float32 {
		switch fnOut.Kind() {
		case reflect.Int:
			return Float32ToInt(itemsValue, fn)
		case reflect.String:
			return Float32ToString(itemsValue, fn)
		case reflect.Bool:
			return Float32ToBool(itemsValue, fn)
		case reflect.Float32:
			return Float32ToFloat32(itemsValue, fn)
		case reflect.Float64:
			return Float32ToFloat64(itemsValue, fn)
		}
	}
	return nil
}
