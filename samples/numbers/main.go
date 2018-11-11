package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 3, 5, 7, 9}

func main() {
	newList := koazee.Stream().
		Drop(3).
		With(numbers).
		Out().
		Val().([]int)

	for _, number := range newList {
		fmt.Printf("%d\n", number)
	}
}
