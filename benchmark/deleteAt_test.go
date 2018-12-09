package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkDeleteAtString10(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.DeleteAt(0).Do()
	}
	if len(outStream.Out().Val().([]string)) != 9 {
		b.Fatal("The output stream should contains 9 elements")
	}
	for index, val := range outStream.Out().Val().([]string) {
		if val != strings10[index+1] {
			b.Fatalf("The element at index %d should be %s", index, strings10[index+1])
		}
	}
}

func BenchmarkDeleteAtString100(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		outStream = stream.DeleteAt(0).Do()
	}
	if len(outStream.Out().Val().([]string)) != 99 {
		b.Fatal("The output stream should contains 99 elements")
	}
	for index, val := range outStream.Out().Val().([]string) {
		if val != strings100[index+1] {
			b.Fatalf("The element at index %d should be %s", index, strings100[index+1])
		}
	}
}

func BenchmarkDeleteAtString1000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		outStream = stream.DeleteAt(0).Do()
	}
	if len(outStream.Out().Val().([]string)) != 999 {
		b.Fatal("The output stream should contains 999 elements")
	}
	for index, val := range outStream.Out().Val().([]string) {
		if val != strings1000[index+1] {
			b.Fatalf("The element at index %d should be %s", index, strings1000[index+1])
		}
	}
}

func BenchmarkDeleteAtString5000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		outStream = stream.DeleteAt(0).Do()
	}
	if len(outStream.Out().Val().([]string)) != 4999 {
		b.Fatal("The output stream should contains 4999 elements")
	}
	for index, val := range outStream.Out().Val().([]string) {
		if val != strings5000[index+1] {
			b.Fatalf("The element at index %d should be %s", index, strings5000[index+1])
		}
	}
}
