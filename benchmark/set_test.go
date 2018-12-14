package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkSetString10(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Set(0, "hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 10 {
		b.Fatalf("Stream should have %d but retrieved %d", 11, len(elements))
	}
	if elements[0] != "hello" {
		b.Fatal("Invalid set value")
	}
}

func BenchmarkSetString100(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Set(0, "hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 100 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
	if elements[0] != "hello" {
		b.Fatal("Invalid set value")
	}
}

func BenchmarkSetString1000(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Set(0, "hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 1000 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
	if elements[0] != "hello" {
		b.Fatal("Invalid set value")
	}
}

func BenchmarkSetString5000(b *testing.B) {
	var outStream stream.Stream
	b.StopTimer()
	stream := koazee.
		StreamOf(strings5000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		outStream = stream.Set(0, "hello").Do()
	}
	elements := outStream.Out().Val().([]string)
	if len(elements) != 5000 {
		b.Fatalf("Stream should have %d but retrieved %d", 101, len(elements))
	}
	if elements[0] != "hello" {
		b.Fatal("Invalid set value")
	}
}
