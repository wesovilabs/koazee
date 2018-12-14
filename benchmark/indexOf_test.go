package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/internal/indexof"

	"testing"
)

/*
BenchmarkIndexOfString10FirstElement-4     	 5000000	       275 ns/op
BenchmarkIndexOfString100FirstElement-4    	 5000000	       277 ns/op
BenchmarkIndexOfString1000FirstElement-4   	 5000000	       280 ns/op
BenchmarkIndexOfString5000FirstElement-4   	 5000000	       276 ns/op
BenchmarkIndexOfString10LastElement-4      	 5000000	       273 ns/op
BenchmarkIndexOfString100LastElement-4     	 5000000	       327 ns/op
BenchmarkIndexOfString1000LastElement-4    	 2000000	       764 ns/op
BenchmarkIndexOfString5000LastElement-4    	  300000	      3687 ns/op
BenchmarkIndexOfString10NotFound-4         	 5000000	       266 ns/op
BenchmarkIndexOfString100NotFound-4        	 5000000	       337 ns/op
BenchmarkIndexOfString1000NotFound-4       	 1000000	      1248 ns/op
BenchmarkIndexOfString5000NotFound-4       	  200000	      7288 ns/op
*/

func BenchmarkIndexOfString10FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.IndexOf(strings10[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString100FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.IndexOf(strings100[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString1000FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.IndexOf(strings1000[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString5000FirstElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.IndexOf(strings5000[0])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString10LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.IndexOf(strings10[9])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString100LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.IndexOf(strings100[99])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString1000LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.IndexOf(strings1000[999])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString5000LastElement(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.IndexOf(strings5000[4999])
	}
	if index == indexof.InvalidIndex {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkIndexOfString10NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		index, _ = stream.IndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkIndexOfString100NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		index, _ = stream.IndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkIndexOfString1000NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		index, _ = stream.IndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkIndexOfString5000NotFound(b *testing.B) {
	b.StopTimer()
	var index = indexof.InvalidIndex
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		index, _ = stream.IndexOf("notfound")
	}
	if index != indexof.InvalidIndex {
		b.Fatalf("Element should not be found")
	}
}
