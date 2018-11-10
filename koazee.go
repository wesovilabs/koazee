// Package koazee contains code for library koazee
package koazee

import (
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/stream"
	"github.com/wesovilabs/koazee/utils"
)

// Stream Initialize an empty stream
func Stream() stream.S {
	return stream.New(nil)
}

// StreamOf loads the data into the stream
func StreamOf(data interface{}) stream.S {
	nature := utils.NatureOf(data)
	switch nature {
	case utils.NatureArray:
		adapter := streamAdapter{}
		s := adapter.toStream(data)
		return s
	default:
		s := stream.Error(errors.InvalidType(":load",
			"Unsupported type! Only arrays are permitted"))
		return s
	}
}
