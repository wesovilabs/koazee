package benchmark

import (
	"github.com/wesovilabs/koazee"
	"testing"
)

func BenchmarkReverseString10SumLen(b *testing.B) {
	b.StopTimer()
	input := strings10
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Reverse().Do()
	}
}

func BenchmarkReverseString100SumLen(b *testing.B) {
	b.StopTimer()
	input := strings100
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Reverse().Do()
	}
}

func BenchmarkReverseString1000SumLen(b *testing.B) {
	b.StopTimer()
	input := strings1000
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Reverse().Do()
	}
}

func BenchmarkReverseString5000SumLen(b *testing.B) {
	b.StopTimer()
	input := strings5000
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(input)
		b.StartTimer()
		stream.Reverse().Do()
	}
}
