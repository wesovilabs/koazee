// Package koazee contains code for library koazee
package koazee

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"
)

// StreamOf loads the user data into the stream
func Stream() stream.S {
	return stream.New(nil)
}

// StreamOf loads the user data into the stream
func StreamOf(data interface{}) stream.S {
	nature := natureOf(data)
	switch nature {
	case natureArray:
		adapter := streamAdapter{}
		s:=adapter.toStream(data)
		return s
	default:
		s := stream.Error(errors.InvalidType(":load", "Unsupported type! Only arrays are permitted"))
		return s
	}
}
