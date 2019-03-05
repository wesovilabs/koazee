package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/utils"
)

func TestStream_DropLeft(t *testing.T) {
	expect := []int{34, 45, 67}
	actual := koazee.StreamOf([]int{44, 55, 66, 34, 45, 67}).
		DropLeft(34).
		Out().Val()
	assert.Equal(t, expect, actual)
}

func TestStream_DropLeft_With_Struct(t *testing.T) {
	expect := []utils.Person{
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		DropLeft(utils.Person{"Jim", "Doe", 56, true}).
		Out().Val()
	assert.Equal(t, expect, actual)
}

func TestStream_DropLeft_With_Struct_NotFound(t *testing.T) {
	expect := []utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		Filter(func(person utils.Person) bool { return person.Age < 25 }).
		DropLeft(utils.Person{"Jamie", "Doe", 89, false}).
		Out().Val()
	assert.Equal(t, expect, actual)
}
