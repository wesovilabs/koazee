//Package errors implements utilities to handle errors in koazee execution
package errors

const ()

// ErrCode type
type ErrCode string

var errInvalidType ErrCode = "koazee:invalid-type"
var errUnknown ErrCode = "koazee:unknown"
var errInvalidIndex ErrCode = "koazee:index"
var errItemsNil ErrCode = "koazee:items-nil"
var errInvalidArgument ErrCode = "koazee:argument"

// InvalidType creates a invalid type error
func InvalidType(op string, msg string, args ...interface{}) *Error {
	return New(op, errInvalidType, msg, args...)
}

// InvalidArgument creates a invalid type error
func InvalidArgument(op string, msg string, args ...interface{}) *Error {
	return New(op, errInvalidArgument, msg, args...)
}

// ItemsNil creates a invalid type error
func ItemsNil(op string, msg string, args ...interface{}) *Error {
	return New(op, errItemsNil, msg, args...)
}

// InvalidStreamIndex creates a invalid type error
func InvalidStreamIndex(op string, msg string, args ...interface{}) *Error {
	return New(op, errInvalidIndex, msg, args...)
}

// String print string of error code
func (c ErrCode) String() string {
	return string(c)
}
