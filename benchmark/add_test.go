package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkAddString10(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Add("hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 11 {
		b.Fatalf("Stream should have %d but retrieved %d", 11, len(elements))
	}
}

func BenchmarkAddString100(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Add("hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 101 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
}

func BenchmarkAddString1000(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Add("hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 1001 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
}

func BenchmarkAddString5000(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings5000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Add("hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 5001 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
}
