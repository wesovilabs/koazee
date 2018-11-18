package first

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

func BenchmarkString10Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.First()

	}
	firstOutput = result
}

func BenchmarkString100Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.First()

	}
	firstOutput = result
}

func BenchmarkString1000Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.First()

	}
	firstOutput = result
}

func BenchmarkString10000Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.First()

	}
	firstOutput = result
}
