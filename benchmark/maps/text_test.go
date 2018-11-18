package maps

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"strings"
	"testing"
)

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings10000 = utils.ArrayOfString(1, 10, 10000)

func BenchmarkString10UppercaseKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(strings.ToUpper).Out()
	}
}
func BenchmarkString100UppercaseKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(strings.ToUpper).Out()
	}
}
func BenchmarkString1000UppercaseKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(strings.ToUpper).Out()
	}
}
func BenchmarkString10000UppercaseKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(strings.ToUpper).Out()
	}
}

func BenchmarkString10LenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(func(val string) int { return len(val) }).Out()
	}
}

func BenchmarkString100LenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(func(val string) int { return len(val) }).Out()
	}
}

func BenchmarkString1000LenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(func(val string) int { return len(val) }).Out()
	}
}

func BenchmarkString10000LenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Map(func(val string) int { return len(val) }).Out()
	}
}
