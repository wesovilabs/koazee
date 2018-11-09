//Package errors implements utilities to handle errors in koazee execution
package errors

import (
	"fmt"
)

// Error encapsulates error info
type Error struct {
	Op   string
	Code ErrCode
	Msg  string
}

// Error function
func (e Error) Error() string {
	return fmt.Sprintf("%s - %s: %s",e.Op, e.Code, e.Msg)
}

// New creates an Error
func New(operation string, code ErrCode, msgFormat string, args ...interface{}) *Error {
	return &Error{
		Op:   operation,
		Code: code,
		Msg:  fmt.Sprintf(msgFormat, args...),
	}
}
