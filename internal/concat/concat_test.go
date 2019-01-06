package concat_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/utils"
)

func concatStrings(t *testing.T) {
	stream := koazee.StreamOf([]string{"wolf", "ox", "dog", "lion", "hamster"})
	concatStream := stream.Concat([]string{"cat", "pig", "mouse"})
	assert.Equal(t, concatStream, []string{"wolf", "ox", "dog", "lion", "hamster", "cat", "pig", "mouse"})
}

func concatIntPtr(t *testing.T) {
	stream := koazee.StreamOf([]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4)})
	concatStream := stream.Concat([]*int{utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)})
	assert.Equal(t, concatStream, []*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4), utils.IntPtr(5), utils.IntPtr(6), utils.IntPtr(7)})
}

func concatComplex(t *testing.T) {
	stream := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	})
	concatStream := stream.Concat([]utils.Person{
		{"Mike", "Doe", 13, false},
		{"Tom", "Doe", 43, true},
		{"Mary", "Doe", 33, false},
		{"Jeff", "Doe", 11, true},
	})
	assert.Equal(t, concatStream, []utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
		{"Mike", "Doe", 13, false},
		{"Tom", "Doe", 43, true},
		{"Mary", "Doe", 33, false},
		{"Jeff", "Doe", 11, true},
	})
}

func Testconcat_Run(t *testing.T) {
	concatStrings(t)
	concatIntPtr(t)
	concatComplex(t)
}
