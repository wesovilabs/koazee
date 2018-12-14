package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/internal/indexof"

	"testing"
)

/*
BenchmarkLastIndexOfString10FirstElement-4     	 5000000	       274 ns/op
BenchmarkLastIndexOfString100FirstElement-4    	 5000000	       325 ns/op
BenchmarkLastIndexOfString1000FirstElement-4   	 2000000	       582 ns/op
BenchmarkLastIndexOfString5000FirstElement-4   	  300000	      3832 ns/op
BenchmarkLastIndexOfString10LastElement-4      	 5000000	       275 ns/op
BenchmarkLastIndexOfString100LastElement-4     	 5000000	       279 ns/op
BenchmarkLastIndexOfString1000LastElement-4    	 5000000	       280 ns/op
BenchmarkLastIndexOfString5000LastElement-4    	 5000000	       279 ns/op
BenchmarkLastIndexOfString10NotFound-4         	 5000000	       243 ns/op
BenchmarkLastIndexOfString100NotFound-4        	 5000000	       361 ns/op
BenchmarkLastIndexOfString1000NotFound-4       	 1000000	      1433 ns/op
BenchmarkLastIndexOfString5000NotFound-4       	  200000	      7267 ns/op
*/

func BenchmarkLastIndexOfString10FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings10[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString100FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings100[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString1000FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings1000[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString5000FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings5000[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString10LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings10[9])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString100LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings100[99])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString1000LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings1000[999])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString5000LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.LastIndexOf(strings5000[4999])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkLastIndexOfString10NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.LastIndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkLastIndexOfString100NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.LastIndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkLastIndexOfString1000NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.LastIndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkLastIndexOfString5000NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.LastIndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}
