package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var stream = koazee.Stream().Map(func(val int) int {
	return val * 2
}).Reduce(func(acc, val int) int {
	return acc + val
})

func main() {
	var result = stream.With([]int{1,2,3,4,5}).Int()
	fmt.Println(result)
	fmt.Println(stream.With([]int{1,3,3,4,5}).Int())
	fmt.Println(stream.With([]int{1,2,3,4,5}).Int())
}
