package set

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
)

// OpCode identifier for operation set
const OpCode = "set"

// Set operation struct
type Set struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
	ItemValue  interface{}
	Len        int
	Index      int
}

// Run performs the operation
func (op *Set) Run() (reflect.Value, *errors.Error) {
	info, err := op.validate()
	if err != nil {
		return reflect.ValueOf(nil), err
	}
	info.index = op.Index

	if found, result := dispatch(op.ItemsValue, op.ItemValue, info); found {
		return reflect.ValueOf(result), nil
	}
	sliceType := reflect.SliceOf(op.ItemsType)
	out := reflect.MakeSlice(sliceType, op.Len, op.Len)
	reflect.Copy(out, op.ItemsValue)
	v := out.Index(op.Index)
	v.Set(reflect.ValueOf(op.ItemValue))
	return out, nil
}

func (op *Set) validate() (*setInfo, *errors.Error) {
	item := &setInfo{}
	itemType := reflect.TypeOf(op.ItemValue)
	if val := cache.get(op.ItemsType, itemType); val != nil {
		return val, nil
	}
	if op.Len == 0 {
		return nil, errors.EmptyStream(OpCode, "It can not be taken an element "+
			"from an empty Stream")
	}
	if op.Index < 0 || op.Len-1 < op.Index {
		return nil, errors.InvalidIndex(OpCode,
			"The length of this Stream is %d, so the index must be "+
				"between 0 and %d", op.Len, op.Len-1)
	}
	if op.ItemValue == nil && op.ItemsType.Kind() == reflect.Ptr {
		item.itemType = &op.ItemsType
		cache.add(op.ItemsType, itemType, item)
		return item, nil
	}
	if itemType != op.ItemsType {
		return nil, errors.InvalidArgument(OpCode,
			"The type of the argument must be %s",
			op.ItemsType.String())
	}
	item.itemType = &itemType
	cache.add(op.ItemsType, itemType, item)
	return item, nil
}

/**
BenchmarkSetString10-4     	 2000000	       784 ns/op
BenchmarkSetString100-4    	 1000000	      1407 ns/op
BenchmarkSetString1000-4   	  200000	      6966 ns/op
BenchmarkSetString5000-4   	   50000	     36694 ns/op
*/
