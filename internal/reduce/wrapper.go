package reduce

// reduceIntToInt function that iterates over function like this func(int int)int
func reduceIntToInt(items []int, fn func(int, int) int) int {
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

// reduceStringToString function that iterates over function like this func(int int)int
func reduceStringToInt(items []string, fn func(int, string) int) int {
	acc := 0
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}

// reduceStringToString function that iterates over function like this func(string int)string
func reduceStringToString(items []string, fn func(string, string) string) string {
	acc := ""
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}
