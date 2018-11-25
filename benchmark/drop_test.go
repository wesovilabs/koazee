package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

/**
BenchmarkDropString10UnExistingElement-4      	  500000	      3242 ns/op
BenchmarkDropString100UnExistingElement-4     	  100000	     24672 ns/op
BenchmarkDropString1000UnExistingElement-4    	    5000	    216298 ns/op
BenchmarkDropString10000UnExistingElement-4   	    1000	   2270868 ns/op

pkg: github.com/wesovilabs/koazee/benchmark
BenchmarkDropString10UnExistingElement-4      	  500000	      3584 ns/op
BenchmarkDropString100UnExistingElement-4     	   50000	     23317 ns/op
BenchmarkDropString1000UnExistingElement-4    	    5000	    207377 ns/op
BenchmarkDropString10000UnExistingElement-4   	     500	   2393964 ns/op

BenchmarkDropString10UnExistingElement-4      	 1000000	      1279 ns/op
BenchmarkDropString100UnExistingElement-4     	  500000	      3405 ns/op
BenchmarkDropString1000UnExistingElement-4    	  100000	     15891 ns/op
BenchmarkDropString10000UnExistingElement-4   	    3000	    559096 ns/op

BenchmarkDropString10UnExistingElement-4      	 2000000	       674 ns/op
BenchmarkDropString100UnExistingElement-4     	 2000000	       741 ns/op
BenchmarkDropString1000UnExistingElement-4    	 1000000	      1359 ns/op
BenchmarkDropString10000UnExistingElement-4   	  200000	      9149 ns/op

*/
func BenchmarkDropString10(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	var dropped = strings10[0]
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Drop(dropped).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == dropped {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString100(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	var dropped = strings100[0]
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		outStream = stream.Drop(dropped).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == dropped {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString1000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	var dropped = strings1000[0]
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		outStream = stream.Drop(dropped).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == dropped {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString10000(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	var dropped = strings10000[0]
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10000)
		b.StartTimer()
		outStream = stream.Drop(dropped).Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == dropped {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString10UnExistingElement(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		outStream = stream.Drop("unExisting").Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == "unExisting" {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString100UnExistingElement(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		outStream = stream.Drop("unExisting").Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == "unExisting" {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}
func BenchmarkDropString1000UnExistingElement(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		outStream = stream.Drop("unExisting").Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == "unExisting" {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}

func BenchmarkDropString10000UnExistingElement(b *testing.B) {
	b.StopTimer()
	var outStream stream.Stream
	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10000)
		b.StartTimer()
		outStream = stream.Drop("unExisting").Do()
	}
	for _, val := range outStream.Out().Val().([]string) {
		if val == "unExisting" {
			b.Fatalf("The output stream should not contain %s", val)
		}
	}
}
