package main

import (
	"github.com/wesovilabs/koazee"
)

var numbers1 = []int{1, 3, 5, 7, 9}
var numbers2 = []int{2, 4, 6, 8}
var numbers3 = []int{3, 4, 2}

func main() {

	lowerThan5 := func(val int) bool { return val < 5 }
	duplicate := func(val int) int { return val * 2 }
	sum := func(acc, value int) int {
		return acc + value
	}

	stream1 := koazee.Stream().
		SetTraceID("stream-1").
		Filter(lowerThan5).
		Map(duplicate).
		Add(2).
		With(numbers1)


	stream2 := koazee.Stream().
		SetTraceID("stream-2").
		Add(3).
		With(numbers2)

	stream3 := koazee.StreamOf(numbers3).SetTraceID("stream-3")

	stream4:=koazee.
		Stream().
		SetTraceID("stream-4").
		Compose(stream1, stream2, stream3)

	stream4.Reduce(sum).Int()
}
