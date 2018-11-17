package _map

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/benchmark/utils"
	"github.com/wesovilabs/koazee/stream"
	"strings"
	"testing"
)

var strings10 = utils.ArrayOfString(1, 10, 10)
var strings100 = utils.ArrayOfString(1, 10, 100)
var strings1000 = utils.ArrayOfString(1, 10, 1000)
var strings10000 = utils.ArrayOfString(1, 10, 10000)

func BenchmarkString10UppercaseKoazee(b *testing.B) {
	var result *stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings10).
			Map(strings.ToUpper).
			Out()
	}
	mapOutput = result
}
func BenchmarkString10LenKoazee(b *testing.B) {
	var result *stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings10).
			Map(func(val string) int { return len(val) }).
			Out()
	}
	mapOutput = result
}

func BenchmarkString100LenKoazee(b *testing.B) {
	var result *stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings100).
			Map(func(val string) int { return len(val) }).
			Out()
	}
	mapOutput = result
}

func BenchmarkString1000LenKoazee(b *testing.B) {
	var result *stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings1000).
			Map(func(val string) int { return len(val) }).
			Out()
	}
	mapOutput = result
}

func BenchmarkString10000LenKoazee(b *testing.B) {
	var result *stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings10000).
			Map(func(val string) int { return len(val) }).
			Out()
	}
	mapOutput = result
}

/**
func BenchmarkString100UppercaseKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings100).
			Map(strings.ToUpper).
			Out()
	}
	mapOutput = result
}

func BenchmarkString1000UppercaseKoazee(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.
			StreamOf(strings1000).
			Map(strings.ToUpper).
			Out()
	}
	mapOutput = result
}
**/
