// Package koazee contains code for library koazee
package koazee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var arraysTypes = []struct {
	data interface{}
}{
	{
		[]string{"a", "home"},
	},
	{
		[]int{2, 10},
	},
	{
		[]bool{false, false},
	},
	{
		[]int32{},
	},
	{
		[]struct{ firstName string }{{firstName: "John"}, {firstName: "David"}},
	},
	{
		[]interface{}{1, "home", false, nil},
	},
	{
		make([]string, 2),
	},
}

var unsupportedTypes = []struct {
	data interface{}
}{
	{
		"home",
	},
	{
		10,
	},
	{
		false,
	},
	{
		nil,
	},
	{
		struct{ firstName string }{firstName: "John"},
	},
	{
		make(map[string]int),
	},
	{
		&struct{ firstName string }{firstName: "John"},
	},
}

func TestNatureOf(t *testing.T) {
	for _, c := range arraysTypes {
		streamType := natureOf(c.data)
		assert.Equal(t, natureArray, streamType)
	}
	for _, c := range unsupportedTypes {
		streamType := natureOf(c.data)
		assert.Equal(t, unsupported, streamType)
	}
}

func TestIsArray(t *testing.T) {
	for _, c := range arraysTypes {
		result := isArray(c.data)
		assert.True(t, result)
	}
	for _, c := range unsupportedTypes {
		result := isArray(c.data)
		assert.False(t, result)
	}
}
