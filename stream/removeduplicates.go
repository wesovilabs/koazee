package stream

import (
	"reflect"

	"github.com/wesovilabs/koazee/errors"
)

// OpCodeRemoveDuplicates identifier for operation removeDuplicates
const OpCodeRemoveDuplicates = "removeDuplicates"

type removeDuplicates struct {
}

func (op *removeDuplicates) name() string {
	return OpCodeRemoveDuplicates
}

func (op *removeDuplicates) run(s Stream) Stream {
	if err := op.validate(s); err != nil {
		s.err = err
		return s
	}

	items := reflect.ValueOf(s.items)
	itemsType := reflect.TypeOf(s.items).Elem()
	newItems := reflect.MakeSlice(reflect.SliceOf(itemsType), 0, 0)
	for index := 0; index < items.Len(); index++ {

		if !arrayContains(newItems, items.Index(index)) {
			newItems = reflect.Append(newItems, items.Index(index))
		}
	}
	s.items = newItems.Interface()
	return s
}

func (op *removeDuplicates) validate(s Stream) *errors.Error {
	if s.items == nil {
		return errors.EmptyStream(op.name(), "A nil Stream can not be iterated")
	}
	return nil
}

// RemoveDuplicates remove all thoese elements are duplicated in the Stream, leaving only
// an element with the same value
func (s Stream) RemoveDuplicates() Stream {
	s.operations = append(s.operations, &removeDuplicates{})
	return s
}

func arrayContains(items reflect.Value, value reflect.Value) bool {
	for index := 0; index < items.Len(); index++ {

		if items.Index(index).Kind() == reflect.Ptr {
			if items.Index(index).Elem().Interface() == value.Elem().Interface() {
				return true
			}
		} else {
			if items.Index(index).Interface() == value.Interface() {
				return true
			}
		}
	}
	return false
}
