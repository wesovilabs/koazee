package pop_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/utils"
	"testing"
)

func TestPop_Run(t *testing.T) {
	popStringWithElements(t)
	popIntWith2Elements(t)
	popIntWithASingleElement(t)
	popIntWithNonElements(t)
	popArrayWithStructs(t)
}

func popStringWithElements(t *testing.T) {
	stream := koazee.StreamOf([]string{"Vegan", "Veggie", "Animals", "Life", "Peace"})
	out, stream := stream.Pop()
	assert.Nil(t, out.Err())
	assert.Equal(t, out.String(), "Vegan")
	assert.Equal(t, []string{"Veggie", "Animals", "Life", "Peace"}, stream.Out().Val().([]string))
}

func popIntWith2Elements(t *testing.T) {
	stream := koazee.StreamOf([]int{10, 10})
	out, stream := stream.Pop()
	assert.Nil(t, out.Err())
	assert.Equal(t, out.Int(), 10)
	assert.Equal(t, []int{10}, stream.Out().Val().([]int))
}

func popIntWithASingleElement(t *testing.T) {
	stream := koazee.StreamOf([]int{10})
	out, stream := stream.Pop()
	assert.Nil(t, out.Err())
	assert.Equal(t, out.Int(), 10)
	assert.Equal(t, []int{}, stream.Out().Val().([]int))
}

func popIntWithNonElements(t *testing.T) {
	stream := koazee.StreamOf([]int{})
	out, stream := stream.Pop()
	assert.Nil(t, out)
	assert.NotNil(t, stream.Out().Err())
}

func popArrayWithStructs(t *testing.T) {
	stream := koazee.StreamOf([]utils.Person{
		{"John", "Doe", 23, true},
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	})
	out, stream := stream.Pop()
	assert.Nil(t, out.Err())
	assert.Equal(t, out.Val().(utils.Person), utils.Person{"John", "Doe", 23, true})
	assert.Equal(t, 3, len(stream.Out().Val().([]utils.Person)))
	assert.Equal(t, []utils.Person{
		{"Jane", "Doe", 22, false},
		{"Jim", "Doe", 56, true},
		{"Jamie", "Doe", 89, false},
	}, stream.Out().Val().([]utils.Person))
}
