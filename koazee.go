// Package koazee contains code for library koazee
package koazee

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"
)

// StreamOf loads the user data into the stream
func StreamOf(data interface{}) stream.S {
	nature := natureOf(data)
	switch nature {
	case natureArray:
		adapter := streamAdapter{}
		return adapter.toStream(data)
	default:
		s := stream.Error(errors.InvalidType(":load", "Unsupported type! Only arrays are permitted"), )
		return s
	}
}
