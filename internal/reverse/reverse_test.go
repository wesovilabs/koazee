package reverse_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/utils"
	"testing"
)

func reverseStrings(t *testing.T) {
	stream := koazee.StreamOf([]string{"wolf", "ox", "dog", "lion", "hamster"})
	reverseStream := stream.Reverse().Do()
	elements := reverseStream.Out().Val().([]string)
	fmt.Println(elements)
	assert.Equal(t, elements[0], "hamster")
	assert.Equal(t, elements[1], "lion")
	assert.Equal(t, elements[2], "dog")
	assert.Equal(t, elements[3], "ox")
	assert.Equal(t, elements[4], "wolf")
}

func reverseIntPtr(t *testing.T) {
	stream := koazee.StreamOf([]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(4)})
	reverseStream := stream.Reverse().Do()
	elements := reverseStream.Out().Val().([]*int)
	fmt.Println(elements)
	assert.Equal(t, *elements[0], 4)
	assert.Equal(t, *elements[1], 3)
	assert.Equal(t, *elements[2], 2)
}

func reverseComplex(t *testing.T) {
	stream := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	})
	reverseStream := stream.Reverse().Do()
	elements := reverseStream.Out().Val().([]utils.Person)
	fmt.Println(elements)
	assert.Equal(t, elements[0].FirstName, "Jamie")
	assert.Equal(t, elements[1].FirstName, "Jim")
	assert.Equal(t, elements[2].FirstName, "Jane")
	assert.Equal(t, elements[3].FirstName, "John")
}

func TestReverse_Run(t *testing.T) {
	reverseStrings(t)
	reverseIntPtr(t)
	reverseComplex(t)
}
