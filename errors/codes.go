//Package errors implements utilities to handle errors in koazee
package errors

// ErrCode type
type ErrCode string

const ErrInvalidType ErrCode = "err.invalid-type"
const ErrUnknown ErrCode = "err.unknown"
const ErrInvalidIndex ErrCode = "err.invalid-index"
const ErrItemsNil ErrCode = "err.items-nil"
const ErrInvalidArgument ErrCode = "err.invalid-argument"

// InvalidType creates a invalid type error
func InvalidType(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidType, msg, args...)
}

// InvalidArgument creates a invalid type error
func InvalidArgument(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidArgument, msg, args...)
}

// ItemsNil creates a invalid type error
func ItemsNil(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrItemsNil, msg, args...)
}

// InvalidIndex creates a invalid type error
func InvalidIndex(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidIndex, msg, args...)
}

// String print string of error code
func (c ErrCode) String() string {
	return string(c)
}
