package stream

const OperationTraceIDIdentifier = ":traceID"

// with struct for defining add operation
type traceID struct {
	value string
}

func (op *traceID) name() string {
	return OperationTraceIDIdentifier
}

// Run performs the operations whenever is called
func (op *traceID) run(s stream) stream {
	s.traceID = op.value
	return s
}

func (s stream) SetTraceID(value string) S {
	return (&traceID{value}).run(s)
}
