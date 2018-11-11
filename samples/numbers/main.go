package main

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/logger"
)

var numbers = []int{1, 2, 3, 4, 4, 1}

func main() {
	logger.Enabled = true
	koazee.Stream().
		With(numbers).
		Filter(func(val int) bool {
			return val%2 == 0
		}).
		Map(func(val int) int {
			return val * 2
		}).
		RemoveDuplicates().
		Reduce(func(acc, val int) int {
			return acc + val
		})
}
