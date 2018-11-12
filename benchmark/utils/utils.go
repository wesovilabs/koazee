package utils

import "math/rand"

func arrayOfString(minLen, maxLen, cap int) []string {
	array := make([]string, cap)
	for i := 0; i < cap; i++ {
		array[i] = ""
	}
	return array
}

func ArrayOfInt(min, max, cap int) []int {
	array := make([]int, cap)
	for i := 0; i < cap; i++ {
		array[i] = rand.Intn(max-min) + min
	}
	return array
}
