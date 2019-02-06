package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkFirstString10(b *testing.B) {
	var result stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result = stream.First().Do()
	}
	if strings10[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings10[0], result.String())
	}
}

func BenchmarkFirstString100(b *testing.B) {
	var result stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result = stream.First().Do()
	}
	if strings100[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings100[0], result.String())
	}
}

func BenchmarkFirstString1000(b *testing.B) {
	var result stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result = stream.First().Do()
	}
	if strings1000[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings1000[0], result.String())
	}
}

func BenchmarkFirstString5000(b *testing.B) {
	var result stream.Output
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		result = stream.First().Do()
	}
	if strings5000[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings5000[0], result.String())
	}
}
