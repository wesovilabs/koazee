package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"math/rand"
	"testing"
)

func BenchmarkAtString10(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	index := rand.Intn(10)
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result = stream.At(index)
	}
	if strings10[index] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings10[index], result.String())
	}
}

func BenchmarkAtString100(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	index := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result = stream.At(index)
	}
	if strings100[index] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings100[index], result.String())
	}
}

func BenchmarkAtString1000(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	index := rand.Intn(1000)
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result = stream.At(index)
	}
	if strings1000[index] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings1000[index], result.String())
	}
}

func BenchmarkAtString10000(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	index := rand.Intn(10000)
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10000)
		b.StartTimer()
		result = stream.At(index)
	}
	if strings10000[index] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings10000[index], result.String())
	}
}
