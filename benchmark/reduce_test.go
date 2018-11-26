package benchmark

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"testing"
)

func BenchmarkReduceString10SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings10)
		b.StartTimer()
		result = stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	expected := 0
	for _, val := range strings10 {
		expected += len(val)
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceString100SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings100)
		b.StartTimer()
		result = stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	expected := 0
	for _, val := range strings100 {
		expected += len(val)
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceString1000SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings1000)
		b.StartTimer()
		result = stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	expected := 0
	for _, val := range strings1000 {
		expected += len(val)
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceString5000SumLen(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(strings5000)
		b.StartTimer()
		result = stream.Reduce(func(acc int, val string) int { return acc + len(val) })
	}
	expected := 0
	for _, val := range strings5000 {
		expected += len(val)
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceInt10Sum(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers10)
		b.StartTimer()
		result = stream.Reduce(func(acc, val int) int { return acc + val })
	}
	expected := 0
	for _, val := range numbers10 {
		expected += val
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceInt100Sum(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers100)
		b.StartTimer()
		result = stream.Reduce(func(acc, val int) int { return acc + val })
	}
	expected := 0
	for _, val := range numbers100 {
		expected += val
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceInt1000Sum(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers1000)
		b.StartTimer()
		result = stream.Reduce(func(acc, val int) int { return acc + val })
	}
	expected := 0
	for _, val := range numbers1000 {
		expected += val
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}

func BenchmarkReduceInt5000Sum(b *testing.B) {
	b.StopTimer()
	var result stream.Output

	for i := 0; i < b.N; i++ {
		stream := koazee.
			StreamOf(numbers2500)
		b.StartTimer()
		result = stream.Reduce(func(acc, val int) int { return acc + val })
	}
	expected := 0
	for _, val := range numbers2500 {
		expected += val
	}
	if result.Int() != expected {
		b.Fatalf("Invalid reduce function, Expected %d but retrieved %d", expected, result.Int())
	}
}
