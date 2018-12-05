package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkPopString10(b *testing.B) {
	var result *stream.Output
	var streamOut stream.Stream
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result, streamOut = stream.Pop()
	}
	if strings10[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings10[0], result.String())
	}
	if len(streamOut.Out().Val().([]string)) != 9 {
		b.Fatalf("The stream should have %d elements but it has %d", 9, len(streamOut.Out().Val().([]string)))
	}
}

func BenchmarkPopString100(b *testing.B) {
	var result *stream.Output
	var streamOut stream.Stream
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result, streamOut = stream.Pop()
	}
	if strings100[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings100[0], result.String())
	}
	if len(streamOut.Out().Val().([]string)) != 99 {
		b.Fatalf("The stream should have %d elements but it has %d", 99, len(streamOut.Out().Val().([]string)))
	}
}

func BenchmarkPopString1000(b *testing.B) {
	var result *stream.Output
	var streamOut stream.Stream
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result, streamOut = stream.Pop()
	}
	if strings1000[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings1000[0], result.String())
	}
	if len(streamOut.Out().Val().([]string)) != 999 {
		b.Fatalf("The stream should have %d elements but it has %d", 999, len(streamOut.Out().Val().([]string)))
	}
}

func BenchmarkPopString5000(b *testing.B) {
	var result *stream.Output
	var streamOut stream.Stream
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		result, streamOut = stream.Pop()
	}
	if strings5000[0] != result.String() {
		b.Fatalf("It should return value %s but retrieved %s", strings1000[0], result.String())
	}
	if len(streamOut.Out().Val().([]string)) != 4999 {
		b.Fatalf("The stream should have %d elements but it has %d", 4999, len(streamOut.Out().Val().([]string)))
	}
}
/**
BenchmarkPopString10-4     	 3000000	       347 ns/op
BenchmarkPopString100-4    	 5000000	       343 ns/op
BenchmarkPopString1000-4   	 5000000	       338 ns/op
BenchmarkPopString5000-4   	 5000000	       342 ns/op
PASS
 */