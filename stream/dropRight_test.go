package stream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/utils"
)

func TestStream_DropRight(t *testing.T) {
	expect := []string{"hoge", "foo"}
	actual := koazee.StreamOf([]string{"hoge", "foo", "bar"}).
		DropRight("foo").
		Out().Val()
	assert.Equal(t, expect, actual)
}

func TestStream_DropRight_With_Struct(t *testing.T) {
	expect := []utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		DropRight(utils.Person{"Jim", "Doe", 56, true}).
		Out().Val()
	assert.Equal(t, expect, actual)
}

func TestStream_DropRight_With_Cache(t *testing.T) {
	expect := []utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		DropRight(utils.Person{"Jim", "Doe", 56, true}).
		Out().Val()
	assert.Equal(t, expect, actual)
}

func TestStream_DropRight_With_Struct_NotFound(t *testing.T) {
	expect := []utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}
	actual := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}).
		DropRight(utils.Person{"Mike", "Jamie", 89, true}).
		Out().Val()
	assert.Equal(t, expect, actual)
}
