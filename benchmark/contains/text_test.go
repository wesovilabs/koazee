package contains

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"testing"
)

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings10000 = utils.ArrayOfString(1, 10, 10000)

/*
BenchmarkString10Koazee-4      	 1000000	      1025 ns/op
BenchmarkString100Koazee-4     	  200000	      8568 ns/op
BenchmarkString1000Koazee-4    	   20000	     87407 ns/op
BenchmarkString10000Koazee-4   	    2000	    886315 ns/op

*/

func BenchmarkString10Koazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains(strings10[1])
	}
}

func BenchmarkString100Koazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains(strings100[50])

	}
}

func BenchmarkString1000Koazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains(strings1000[500])
	}
}
func BenchmarkString10000Koazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	element := strings10000[4000]
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains(element)
	}
}

func BenchmarkString10NotFoundKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains("undefined")
		// stream.Contains("oyetu")
	}
}

func BenchmarkString100NotFoundKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains("undefined")

	}

}

func BenchmarkString1000NotFoundKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains("undefined")

	}
}

func BenchmarkString10000NotFoundKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Contains("undefined")
	}
}
