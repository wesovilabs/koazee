package main

import (
	"fmt"

	"github.com/wesovilabs/koazee"
)

var numbers = []int{0, 2, 3, 10, 3, 7, 0, -1, -9, 1, -7}

func main() {

	isEven := func(val int) bool { return val%2 == 0 }
	duplicate := func(val int) int { return val * 2 }
	desc := func(left, right int) int {
		if left > right {
			return 1
		}
		if left < right {
			return -1
		}
		return 0
	}
	asc := func(left, right int) int {
		return desc(left, right) * -1
	}
	printNumber := func(val int) {
		fmt.Printf("Value %d\n", val)
	}
	sum := func(acc, value int) int {
		return acc + value
	}
	val, _ := koazee.StreamOf(numbers).
		RemoveDuplicates().
		Filter(isEven).
		Map(duplicate).
		Sort(asc).
		Add(3).
		Drop(20).
		Sort(asc).
		ForEach(printNumber).
		Reduce(sum)
	fmt.Println(val)
}
