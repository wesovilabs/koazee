package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 3, 5, 7, 9}

func main() {
	elements := koazee.Stream().
		With(numbers).
		Filter(func(val int) bool {
			return val > 5
		}).Out().Val().([]int)
	for _, element := range elements {
		fmt.Printf("Number %d elements is in the list\n", element)
	}


}
