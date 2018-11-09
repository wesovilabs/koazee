// Package koazee contains code for library koazee
package koazee

import (
	"reflect"
	"strings"
)

// streamType struct used to classify the types
type streamType int

const (
	// unsupported used for those cases in which koazee can not work with
	unsupported streamType = iota
	// natureArray type for any arrays
	natureArray
)

// natureOf returns a streamType for the given parameter.
func natureOf(data interface{}) streamType {
	if isArray(data) {
		return natureArray
	}
	return unsupported
}

func isArray(data interface{}) bool {
	return data != nil && strings.HasPrefix(reflect.TypeOf(data).String(), "[]")
}
