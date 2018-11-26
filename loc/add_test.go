package loc

import (
	"os"
	"testing"
)

const maxElements = 1000

var elements = make([]int, maxElements)

func TestMain(m *testing.M) {
	for i := 0; i < maxElements; i++ {
		elements[i] = i
	}
	os.Exit(m.Run())
}

func Add(items []int, item int) []int {
	return append(items, item)
}

func Add2(items []int, item int) []int {
	n := len(items)
	output := make([]int, n+1)
	output[n] = item
	copy(items[0:n], output)
	return output
}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	var result []int
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		result = Add(elements, 10)
	}
	if len(result) != maxElements+1 {
		b.Fatalf("Invalid number of elements, expected %d but retrieved %d", maxElements+1, len(result))
	}
}

func BenchmarkAdd2(b *testing.B) {
	b.ReportAllocs()
	var result []int
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		result = Add2(elements, 10)
	}
	if len(result) != maxElements+1 {
		b.Fatalf("Invalid number of elements, expected %d but retrieved %d", maxElements+1, len(result))
	}
}
