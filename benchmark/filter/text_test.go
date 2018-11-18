package filter

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"testing"
)

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings10000 = utils.ArrayOfString(1, 10, 10000)

var filterFn = func(val string) bool { return len(val) < 6 }

/**
koazee Titi
BenchmarkString10TextLenKoazee-4      	  300000	      3976 ns/op
BenchmarkString100TextLenKoazee-4     	   50000	     27080 ns/op
BenchmarkString1000TextLenKoazee-4    	    5000	    243762 ns/op
BenchmarkString10000TextLenKoazee-4   	    1000	   2431268 ns/op
*/
/**
Koazee Gelada
BenchmarkString10TextLenKoazee-4      	 2000000	       753 ns/op
BenchmarkString100TextLenKoazee-4     	 1000000	      2081 ns/op
BenchmarkString1000TextLenKoazee-4    	  100000	     18506 ns/op
BenchmarkString10000TextLenKoazee-4   	    5000	    253658 ns/op
 */
func BenchmarkString10TextLenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Filter(filterFn).Out()
	}
}
func BenchmarkString100TextLenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings100)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Filter(filterFn).Out()
	}
}
func BenchmarkString1000TextLenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings1000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Filter(filterFn).Out()
	}
}
func BenchmarkString10000TextLenKoazee(b *testing.B) {
	b.StopTimer()
	stream := koazee.
		StreamOf(strings10000)
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		stream.Filter(filterFn).Out()
	}
}
