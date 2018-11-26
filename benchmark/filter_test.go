package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

var filterFn = func(val int) bool { return val%2 == 0 }

func BenchmarkFilterNumbers10LenIsEven(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers10)
		b.StartTimer()
		outStream = stream.Filter(filterFn).Do()
	}
	for _, val := range outStream.Out().Val().([]int) {
		if val%2 != 0 {
			b.Fatalf("The output stream should only contains even numbers, and %d is odd", val)
		}
	}
}

func BenchmarkFilterNumbers100LenIsEven(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers100)
		b.StartTimer()
		outStream = stream.Filter(filterFn).Do()
	}
	for _, val := range outStream.Out().Val().([]int) {
		if val%2 != 0 {
			b.Fatalf("The output stream should only contains even numbers, and %d is odd", val)
		}
	}
}

func BenchmarkFilterNumbers1000LenIsEven(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers1000)
		b.StartTimer()
		outStream = stream.Filter(filterFn).Do()
	}
	for _, val := range outStream.Out().Val().([]int) {
		if val%2 != 0 {
			b.Fatalf("The output stream should only contains even numbers, and %d is odd", val)
		}
	}
}

func BenchmarkFilterNumbers5000LenIsEven(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers2500)
		b.StartTimer()
		outStream = stream.Filter(filterFn).Do()
	}
	for _, val := range outStream.Out().Val().([]int) {
		if val%2 != 0 {
			b.Fatalf("The output stream should only contains even numbers, and %d is odd", val)
		}
	}
}
