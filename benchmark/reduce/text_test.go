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
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString100SumLenKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString1000SumLenKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}

func BenchmarkString10000SumLenKoazee(b *testing.B) {
	b.StopTimer()
	var result stream.Output
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	reduceOutput = result
}
