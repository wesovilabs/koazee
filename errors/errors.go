package errors

import (
	"fmt"
)

// Error encapsulates error info
type Error struct {
	op   string
	code ErrCode
	msg  string
	meta map[string]interface{}
}

// Error converts the error into a readable message
func (e Error) Error() string {
	out := fmt.Sprintf("[%s:%s] %s", e.op, e.code, e.msg)
	if e.meta != nil {
		for k, v := range e.meta {
			out += fmt.Sprintf("\n  - %s: %v", k, v)
		}
	}
	return out
}

// New creates a new instance of Error
func New(operation string, code ErrCode, msgFormat string, args ...interface{}) *Error {
	return &Error{
		op:   operation,
		code: code,
		msg:  fmt.Sprintf(msgFormat, args...),
	}
}

// Operation returns the operation associated to the error
func (e Error) Operation() string {
	return e.op
}

// Code returns the code for the error
func (e Error) Code() string {
	return e.code.String()
}

// With permits add extra info to show displayed when printing the error
func (e Error) With(key string, value interface{}) Error {
	if e.meta == nil {
		e.meta = make(map[string]interface{})
	}
	e.meta[key] = value
	return e
}
