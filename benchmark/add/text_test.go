package at

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

/**
Titi
BenchmarkAddString10-4      	 5000000	       280 ns/op
BenchmarkString100Koazee-4     	 5000000	       259 ns/op
BenchmarkString1000Koazee-4    	 5000000	       267 ns/op
BenchmarkString10000Koazee-4   	 5000000	       270 ns/op
*/
func BenchmarkAddString10(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Add("hello").Out()
	}
	atOutput = result
}

func BenchmarkString100Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Add("hello")

	}
	atOutput = result
}

func BenchmarkString1000Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Add("hello")
	}
	atOutput = result
}

func BenchmarkString10000Koazee(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Add("hello")
	}
	atOutput = result
}
