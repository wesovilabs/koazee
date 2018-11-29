package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkLastString10(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result = stream.Last()
	}
	if strings10[9] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings10[9], result.String())
	}
}

func BenchmarkLastString100(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result = stream.Last()
	}
	if strings100[99] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings100[99], result.String())
	}
}

func BenchmarkLastString1000(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result = stream.Last()
	}
	if strings1000[999] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings1000[999], result.String())
	}
}

func BenchmarkLastString5000(b *testing.B) {
	var result *stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		result = stream.Last()
	}
	if strings5000[4999] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings5000[4999], result.String())
	}
}
