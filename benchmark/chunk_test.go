package benchmark

import (
	"testing"

	"github.com/wesovilabs/koazee"
)

func BenchmarkChunkString10(b *testing.B) {
	b.StopTimer()
	input := strings10
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Chunk(2).Do()
	}
}

func BenchmarkChunkString100(b *testing.B) {
	b.StopTimer()
	input := strings100
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Chunk(2).Do()
	}
}

func BenchmarkChunkString1000(b *testing.B) {
	b.StopTimer()
	input := strings1000
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Chunk(2).Do()
	}
}

func BenchmarkChunkString5000(b *testing.B) {
	b.StopTimer()
	input := strings5000
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Chunk(2).Do()
	}
}
