package errors

// ErrCode type used to define the error codes in koazee
type ErrCode string

// ErrInvalidType code for errors related to invalid type
const ErrInvalidType ErrCode = "err.invalid-type"

// ErrInvalidIndex code for errors related to invalid index when accessing to a position in the stream
const ErrInvalidIndex ErrCode = "err.invalid-index"

// ErrItemsNil code for errors produced by a nil stream
const ErrItemsNil ErrCode = "err.items-nil"

// ErrInvalidArgument code for errors produced by an invalid input in the operation
const ErrInvalidArgument ErrCode = "err.invalid-argument"

// InvalidType creates an invalid type error
func InvalidType(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidType, msg, args...)
}

// InvalidArgument creates a invalid argument error
func InvalidArgument(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidArgument, msg, args...)
}

// EmptyStream creates an item nil error
func EmptyStream(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrItemsNil, msg, args...)
}

// InvalidIndex creates a invalid index error
func InvalidIndex(op string, msg string, args ...interface{}) *Error {
	return New(op, ErrInvalidIndex, msg, args...)
}

// String print string of error code
func (c ErrCode) String() string {
	return string(c)
}
