package koazee

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringArraysAsStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dataInput := []string{"Home", "word", "bye bye", "test", "mean", "pop", "music", "telescope"}

		output := StreamOf(dataInput).
			Filter(func(element string) bool {
				return len(element) < 5
			}).
			Map(strings.ToUpper).
			Reduce(func(acc, val string) string {
				return fmt.Sprintf("%s%s", acc, val)
			}).String()
		if output != "HOMEWORDTESTMEANPOP" {

			b.Errorf("Invalid response %s", output)
		}
	}

}
