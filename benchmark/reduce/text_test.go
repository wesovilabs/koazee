package reduce

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings10000 = utils.ArrayOfString(1, 10, 10000)

func BenchmarkString10SumLenKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings10).
			Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString100SumLenKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings100).
			Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString1000SumLenKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings1000).
			Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString10000SumLenKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings10000).
			Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}
