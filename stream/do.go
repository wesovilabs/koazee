package stream

// At returns the element in the Stream in the given position
func (s Stream) Do() Stream {
	sOut := s.run()
	return sOut
}
