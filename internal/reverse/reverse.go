package reverse

import (
	"github.com/wesovilabs/koazee/errors"
	"reflect"
	"sort"
)

// OpCode identifier for operation Reverse
const OpCode = "reverse"

type Reverse struct {
	ItemsValue reflect.Value
	ItemsType  reflect.Type
}

func (rev SortSliceString) Len() int           { return len(rev.value) }
func (rev SortSliceString) Less(i, j int) bool { return i < j }
func (rev SortSliceString) Swap(i, j int)      { rev.value[i], rev.value[j] = rev.value[j], rev.value[i] }

type SortSliceString struct {
	value []string
}

func (op *Reverse) Run() (reflect.Value, *errors.Error) {
	if found, result := dispatch(op.ItemsValue, op.ItemsType); found {
		v := reflect.ValueOf(result)
		return v, nil
	}
	input := op.ItemsValue.Interface()
	sort.SliceStable(input, func(i, j int) bool { return true })
	return reflect.ValueOf(input), nil
}

func (op *Reverse) prepareMapWithKeys() map[interface{}]int {
	keys := make(map[interface{}]int)

	if op.ItemsValue.Index(0).Kind() == reflect.Ptr {
		for index := 0; index < op.ItemsValue.Len(); index++ {
			v := op.ItemsValue.Index(index).Elem().Interface()
			keys[v]++
		}
	} else {
		for index := 0; index < op.ItemsValue.Len(); index++ {
			v := op.ItemsValue.Index(index).Interface()
			keys[v]++
		}
	}
	return keys
}
