package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"strings"
	"testing"
)

/**
BenchmarkSortString10SumLen-4     	   50000	     20361 ns/op
BenchmarkSortString100SumLen-4    	    3000	    427507 ns/op
BenchmarkSortString1000SumLen-4   	     200	   5973410 ns/op
BenchmarkSortString5000SumLen-4   	      30	  42205850 ns/op


BenchmarkSortString10SumLen-4     	 2000000	       815 ns/op
BenchmarkSortString100SumLen-4    	  200000	      8023 ns/op
BenchmarkSortString1000SumLen-4   	   10000	    140208 ns/op
BenchmarkSortString5000SumLen-4   	    2000	    896392 ns/op

*/
func BenchmarkSortString10SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result = stream.Sort(strings.Compare).Do()
	}
	slice := result.Out().Val().([]string)
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			b.Fatalf("Element %s should be positioned after %s", slice[index], slice[index+1])
		}
	}
}

func BenchmarkSortString100SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result = stream.Sort(strings.Compare).Do()
	}
	slice := result.Out().Val().([]string)
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			b.Fatalf("Element %s should be positioned after %s", slice[index], slice[index+1])
		}
	}
}

func BenchmarkSortString1000SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result = stream.Sort(strings.Compare).Do()
	}
	slice := result.Out().Val().([]string)
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			b.Fatalf("Element %s should be positioned after %s", slice[index], slice[index+1])
		}
	}
}

func BenchmarkSortString5000SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		result = stream.Sort(strings.Compare).Do()
	}
	slice := result.Out().Val().([]string)
	for index := 0; index < len(slice)-1; index++ {
		if slice[index] > slice[index+1] {
			b.Fatalf("Element %s should be positioned after %s", slice[index], slice[index+1])
		}
	}
}
