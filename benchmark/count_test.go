package benchmark

import (
	"github.com/wesovilabs/koazee"
	"testing"
)

func BenchmarkCountString10(b *testing.B) {
	var count = 0
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		count, _ = stream.Count()
	}
	if count != 10 {
		b.Fatalf("It should return value %d but retrieved %d", 10, count)
	}
}

func BenchmarkCountString100(b *testing.B) {
	var count = 0
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		count, _ = stream.Count()
	}
	if count != 100 {
		b.Fatalf("It should return value %d but retrieved %d", 100, count)
	}
}

func BenchmarkCountString1000(b *testing.B) {
	var count = 0
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		count, _ = stream.Count()
	}
	if count != 1000 {
		b.Fatalf("It should return value %d but retrieved %d", 1000, count)
	}
}

func BenchmarkCountString5000(b *testing.B) {
	var count = 0
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		count, _ = stream.Count()
	}
	if count != 5000 {
		b.Fatalf("It should return value %d but retrieved %d", 5000, count)
	}
}
