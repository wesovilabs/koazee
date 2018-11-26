package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

var array10WithDuplicates = append(strings10[0:5], strings10[0:5]...)
var array100WithDuplicates = append(strings100[0:50], strings100[0:50]...)
var array1000WithDuplicates = append(strings1000[0:500], strings1000[0:500]...)
var array10000WithDuplicates = append(strings10000[0:5000], strings10000[0:5000]...)

func BenchmarkDuplicatedString10(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(array10WithDuplicates)
		b.StartTimer()
		outStream = stream.RemoveDuplicates().Do()
	}
	keys := make(map[string]bool)
	for _, val := range outStream.Out().Val().([]string) {
		keys[val] = true

	}
	if len(keys) != len(outStream.Out().Val().([]string)) {
		b.Fatalf("There0s duplicated elements")
	}
}

func BenchmarkDuplicatedString100(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(array100WithDuplicates)
		b.StartTimer()
		outStream = stream.RemoveDuplicates().Do()
	}
	keys := make(map[string]bool)
	for _, val := range outStream.Out().Val().([]string) {
		keys[val] = true

	}
	if len(keys) != len(outStream.Out().Val().([]string)) {
		b.Fatalf("There0s duplicated elements")
	}
}

func BenchmarkDuplicatedString1000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(array1000WithDuplicates)
		b.StartTimer()
		outStream = stream.RemoveDuplicates().Do()
	}
	keys := make(map[string]bool)
	for _, val := range outStream.Out().Val().([]string) {
		keys[val] = true

	}
	if len(keys) != len(outStream.Out().Val().([]string)) {
		b.Fatalf("There0s duplicated elements")
	}
}

func BenchmarkDuplicatedString10000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(array10000WithDuplicates)
		b.StartTimer()
		outStream = stream.RemoveDuplicates().Do()
	}
	keys := make(map[string]bool)
	for _, val := range outStream.Out().Val().([]string) {
		keys[val] = true

	}
	if len(keys) != len(outStream.Out().Val().([]string)) {
		b.Fatalf("There0s duplicated elements")
	}
}
