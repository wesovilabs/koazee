package first

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFirst_Run(t *testing.T) {
	input := First{
		ItemsValue: reflect.ValueOf([]string{"Hello", "my", "friend", "?", "How", "are", "you", "?"}),
		Len:        8,
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Equal(t, output.Interface(), "Hello")
}

func TestFirstPtr_Run(t *testing.T) {
	valTrue := true
	valFalse := false
	input := First{
		ItemsValue: reflect.ValueOf([]*bool{&valTrue, &valTrue, &valFalse, &valFalse, &valTrue}),
		Len:        5,
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Equal(t, output.Interface(), &valTrue)
}

func TestFirstError_Run(t *testing.T) {
	input := First{
		ItemsValue: reflect.ValueOf([]*bool{}),
		Len:        0,
	}
	output, err := input.Run()
	assert.Equal(t, reflect.ValueOf(nil), output)
	assert.NotNil(t, err)
}
