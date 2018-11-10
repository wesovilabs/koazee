// Package koazee contains code for library koazee
package utils_test

import (
	"testing"

	"github.com/wesovilabs/koazee/utils"

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

func Test_NatureOf(t *testing.T) {
	for _, c := range arraysTypes {
		streamType := utils.NatureOf(c.data)
		assert.Equal(t, utils.NatureArray, streamType)
	}
	for _, c := range unsupportedTypes {
		streamType := utils.NatureOf(c.data)
		assert.Equal(t, utils.Unsupported, streamType)
	}
}
