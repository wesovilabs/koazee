// Package koazee contains code for library koazee
package koazee

import (
	"reflect"

	"github.com/wesovilabs/koazee/stream"
)

type streamAdapter struct {
	elementType reflect.Type
}

func (adapter *streamAdapter) toStream(data interface{}) stream.S {
	elemType := reflect.TypeOf(data).Elem()
	v := reflect.ValueOf(data)
	slice := reflect.MakeSlice(reflect.SliceOf(elemType), v.Len(), v.Len())
	for index := 0; index < v.Len(); index++ {
		field := v.Index(index)
		slice.Index(index).Set(field)
	}
	items := slice.Interface()
	s := stream.New(items)
	return s
}
