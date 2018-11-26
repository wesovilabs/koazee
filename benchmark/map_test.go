package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"strings"
	"testing"
)

func BenchmarkMapString10ToUpper(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Map(strings.ToUpper).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if strings.ToUpper(val) != val {
			b.Fatalf("All the elements int he stream should be in uppercase and %s is not", val)
		}
	}
}

func BenchmarkMapString100ToUpper(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		outStream = stream.Map(strings.ToUpper).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if strings.ToUpper(val) != val {
			b.Fatalf("All the elements int he stream should be in uppercase and %s is not", val)
		}
	}
}

func BenchmarkMapString1000ToUpper(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		outStream = stream.Map(strings.ToUpper).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if strings.ToUpper(val) != val {
			b.Fatalf("All the elements int he stream should be in uppercase and %s is not", val)
		}
	}
}

func BenchmarkMapString5000ToUpper(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		outStream = stream.Map(strings.ToUpper).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if strings.ToUpper(val) != val {
			b.Fatalf("All the elements int he stream should be in uppercase and %s is not", val)
		}
	}
}

func BenchmarkMapString10ToIntLen(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Map(func(val string) int { return len(val) }).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if len(strings10[index]) != val {
			b.Fatalf("Expected %d but retrieved %d in position %d", len(strings10[index]), val, index)
		}
	}
}

func BenchmarkMapString100ToIntLen(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		outStream = stream.Map(func(val string) int { return len(val) }).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if len(strings100[index]) != val {
			b.Fatalf("Expected %d but retrieved %d in position %d", len(strings100[index]), val, index)
		}
	}
}

func BenchmarkMapString1000ToIntLen(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		outStream = stream.Map(func(val string) int { return len(val) }).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if len(strings1000[index]) != val {
			b.Fatalf("Expected %d but retrieved %d in position %d", len(strings1000[index]), val, index)
		}
	}
}

func BenchmarkMapString5000ToIntLen(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		outStream = stream.Map(func(val string) int { return len(val) }).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if len(strings5000[index]) != val {
			b.Fatalf("Expected %d but retrieved %d in position %d", len(strings5000[index]), val, index)
		}
	}
}
