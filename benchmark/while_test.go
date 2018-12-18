package benchmark

import (
	"testing"

	"github.com/wesovilabs/koazee"
)

func BenchmarkWhileNumbers10Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers10)
		b.StartTimer()
		stream.While(func(_ int) bool { return true }, func(v int) {}).Do()
	}
}

func BenchmarkWhileNumbers100Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers100)
		b.StartTimer()
		stream.While(func(_ int) bool { return true }, func(v int) {}).Do()
	}
}

func BenchmarkWhileNumbers1000Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers1000)
		b.StartTimer()
		stream.While(func(_ int) bool { return true }, func(v int) {}).Do()
	}
}

func BenchmarkWhileNumbers5000Print(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers2500)
		b.StartTimer()
		stream.While(func(_ int) bool { return true }, func(v int) {}).Do()
	}
}
