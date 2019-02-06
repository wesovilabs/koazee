// Package koazee contains code for library koazee
package koazee

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"
	"reflect"
)

// Stream Initialize an empty stream
func Stream() stream.Stream {
	return stream.New([]interface{}{})
}

// StreamOf loads the data into the stream
func StreamOf(data interface{}) stream.Stream {
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		return stream.New(data)
	}
	return stream.Error(errors.InvalidType(":load",
		"Unsupported type! Only arrays are permitted"))

}
