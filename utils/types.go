package utils

import (
	"reflect"
	"strings"
)

// StreamType struct used to classify the types
type StreamType int

const (
	// Unsupported used for those cases in which koazee can not work with
	Unsupported StreamType = iota
	// NatureArray type for any arrays
	NatureArray
)

// NatureOf returns a StreamType for the given parameter.
func NatureOf(data interface{}) StreamType {
	if isArray(data) {
		return NatureArray
	}
	return Unsupported
}

func isArray(data interface{}) bool {

	return data != nil && strings.HasPrefix(reflect.TypeOf(data).String(), "[]")
}
