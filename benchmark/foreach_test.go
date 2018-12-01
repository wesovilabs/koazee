package benchmark

import (
	"github.com/wesovilabs/koazee"
	"testing"
)

func BenchmarkForEachNumbers10Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers10)
		b.StartTimer()
		stream.ForEach(func(_ int) {}).Do()
	}
}

func BenchmarkForEachNumbers100Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers100)
		b.StartTimer()
		stream.ForEach(func(_ int) {}).Do()
	}
}

func BenchmarkForEachNumbers1000Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers1000)
		b.StartTimer()
		stream.ForEach(func(_ int) {}).Do()
	}
}

func BenchmarkForEachNumbers5000Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers2500)
		b.StartTimer()
		stream.ForEach(func(_ int) {}).Do()
	}
}
