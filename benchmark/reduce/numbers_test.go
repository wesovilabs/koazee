package reduce

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"testing"
)

var output int
var numbers10 = utils.ArrayOfInt(0, 1, 10)
var numbers100 = utils.ArrayOfInt(0, 10, 100)
var numbers1000 = utils.ArrayOfInt(0, 100, 1000)
var numbers10000 = utils.ArrayOfInt(0, 1000, 10000)

/**
func BenchmarkSum10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := 0
		for _, v := range numbers10 {
			result += v
		}
		output = result
	}

}
func BenchmarkSum10Koazee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output2 := koazee.
			StreamOf(numbers10).
			Reduce(func(acc, val int) int { return acc + val }).
			Val()
		if output != output2 {
			b.Fail()
		}
	}
}

**/
func BenchmarkSum100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := 0
		for _, v := range numbers100 {
			result += v
		}
	}
}
func BenchmarkSum100Koazee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		koazee.
			StreamOf(numbers100).
			Reduce(func(acc, val int) int { return acc + val }).
			Int()
	}
}

/**
func BenchmarkSum1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := 0
		for _, v := range numbers1000 {
			result += v
		}
	}
}
func BenchmarkSum1000Koazee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		koazee.
			StreamOf(numbers1000).
			Reduce(func(acc, val int) int { return acc + val }).
			Int()
	}
}

func BenchmarkSum10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := 0
		for _, v := range numbers10000 {
			result += v
		}
	}
}
func BenchmarkSum10000Koazee(b *testing.B) {
	for i := 0; i < b.N; i++ {
		koazee.
			StreamOf(numbers10000).
			Reduce(func(acc, val int) int { return acc + val }).
			Int()
	}
}
**/
