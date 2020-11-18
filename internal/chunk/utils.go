package chunk

func min(value1, value2 int) int {
	if value1 < value2 {
		return value1
	}
	return value2
}

func outerCapacity(len, size int) int {
	return len/size + min(len%size, 1)
}

func chunkRangeLast(start, size, max int) int {
	return min(start+size, max)
}
