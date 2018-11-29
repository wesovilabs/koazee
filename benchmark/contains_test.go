package benchmark

import (
	"github.com/wesovilabs/koazee"
	"testing"
)

/*
BenchmarkString10Koazee-4      	 500000	      1025 ns/op
BenchmarkString100Koazee-4     	  200000	      8568 ns/op
BenchmarkString1000Koazee-4    	   20000	     87407 ns/op
BenchmarkString5000Koazee-4   	    2000	    886315 ns/op

*/

func BenchmarkContainsString10FirstElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		found, _ = stream.Contains(strings10[0])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString100FirstElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		found, _ = stream.Contains(strings100[0])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString1000FirstElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		found, _ = stream.Contains(strings1000[0])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString5000FirstElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		found, _ = stream.Contains(strings5000[0])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString10LastElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		found, _ = stream.Contains(strings10[9])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString100LastElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		found, _ = stream.Contains(strings100[99])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString1000LastElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		found, _ = stream.Contains(strings1000[999])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString5000LastElement(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		found, _ = stream.Contains(strings5000[4999])
	}
	if !found {
		b.Fatalf("Element should be found")
	}
}

func BenchmarkContainsString10NotFound(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		found, _ = stream.Contains("notfound")
	}
	if found {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkContainsString100NotFound(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		found, _ = stream.Contains("notfound")
	}
	if found {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkContainsString1000NotFound(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		found, _ = stream.Contains("notfound")
	}
	if found {
		b.Fatalf("Element should not be found")
	}
}

func BenchmarkContainsString5000NotFound(b *testing.B) {
	b.StopTimer()
	var found = false
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		found, _ = stream.Contains("notfound")
	}
	if found {
		b.Fatalf("Element should not be found")
	}
}
