package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkTakeNumbers1000Between1and9(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers10)
		b.StartTimer()
		outStream = stream.Take(1, 9).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if val != numbers10[index+1] {
			b.Fatal("Invalid value for the index")
		}
	}
}

func BenchmarkTakeNumbers1000Between1and99(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers100)
		b.StartTimer()
		outStream = stream.Take(1, 99).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if val != numbers100[index+1] {
			b.Fatal("Invalid value for the index")
		}
	}
}

func BenchmarkTakeNumbers1000Between1and999(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers1000)
		b.StartTimer()
		outStream = stream.Take(1, 999).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if val != numbers1000[index+1] {
			b.Fatal("Invalid value for the index")
		}
	}
}

func BenchmarkTakeNumbers2500etween1and2499(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers2500)
		b.StartTimer()
		outStream = stream.Take(1, 2499).Do()
	}
	for index, val := range outStream.Out().Val().([]int) {
		if val != numbers2500[index+1] {
			b.Fatal("Invalid value for the index")
		}
	}
}
