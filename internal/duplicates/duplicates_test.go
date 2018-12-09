package duplicates

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRemoveDuplicates_Run(t *testing.T) {
	input := RemoveDuplicates{
		ItemsValue: reflect.ValueOf([]string{"Hello", "my", "friend", "?", "How", "are", "you", "?"}),
		ItemsType:  reflect.TypeOf("Hello"),
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Len(t, output.Interface().([]string), 7)
	assert.Equal(t, []string{"Hello", "my", "friend", "?", "How", "are", "you"}, output.Interface().([]string))
}

func TestRemoveDuplicatesPtr_Run(t *testing.T) {
	valTrue := true
	valFalse := false
	input := RemoveDuplicates{

		ItemsValue: reflect.ValueOf([]*bool{&valTrue, &valTrue, &valFalse, &valFalse, &valTrue}),
		ItemsType:  reflect.TypeOf(&valTrue),
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Len(t, output.Interface().([]*bool), 2)
	assert.Equal(t, []*bool{&valTrue, &valFalse}, output.Interface().([]*bool))
}
