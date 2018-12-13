package stream

// Do perform the queue of operations in the stream
func (s Stream) Do() Stream {
	sOut := s.run()
	return sOut
}
