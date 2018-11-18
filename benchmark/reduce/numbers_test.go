package reduce

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

var numbers10 = utils.ArrayOfInt(0, 6, 10)
var numbers100 = utils.ArrayOfInt(0, 25, 100)
var numbers1000 = utils.ArrayOfInt(0, 200, 1000)
var numbers10000 = utils.ArrayOfInt(0, 2000, 10000)

func BenchmarkNumbers10SumKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(numbers10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc, val int) int { return acc + val })
	}
	reduceOutput = result
}

func BenchmarkNumbers100SumKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(numbers100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc, val int) int { return acc + val })
	}
	reduceOutput = result
}

func BenchmarkNumbers1000SumKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(numbers1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc, val int) int { return acc + val })
	}
	reduceOutput = result
}

func BenchmarkNumbers10000SumKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(numbers10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc, val int) int { return acc + val })
	}
	reduceOutput = result
}
