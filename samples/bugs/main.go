package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

func main() {
	for i := 0; i < 2; i++ {
		numbers := koazee.StreamOf([]int{1, 2, 3})
		n := 0
		fmt.Println(n)
		numbers.ForEach(func(num int) {
			fmt.Println("n = ", n, " and num = ", num)
			n += num // koazee's bug: n seems to be cached badly.
		}).Do()
		fmt.Println("n = ", n)
	}
}