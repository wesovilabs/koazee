package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

// We can define out reduceStream pipelines and then be reused for different input data

var mapStream = koazee.Stream().Map(func(val int) int {
	return val * 2
})
var reduceStream = mapStream.Reduce(func(acc, val int) int {
	return acc + val
})

var firstStream = mapStream.First()

func main() {
	fmt.Println(reduceStream.With([]int{1, 1, 3, 4, 5}).Int())
	fmt.Println(reduceStream.With([]int{1, 2, 3, 4, 5}).Int())
	fmt.Println(reduceStream.With([]int{1, 3, 3, 4, 5}).Int())
	fmt.Println(reduceStream.Do().Error())
	res := koazee.StreamOf([]int{1, 2, 3, 4, 5}).Reduce(func(acc, val int) int {
		return acc + val
	}).Do().Int()
	fmt.Println(res)
	res = koazee.StreamOf([]int{1, 2, 3, 4, 5}).Reduce(func(acc, val int) int {
		return acc + val
	}).With([]int{1, 2}).Int()
	fmt.Println(res)

	fmt.Println(koazee.StreamOf([]int{1, 2, 3}).First().Do().Int())
	fmt.Println(firstStream.With([]int{3, 2, 1}).Int())
}
