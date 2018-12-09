package last

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestLast_Run(t *testing.T) {
	input := Last{
		ItemsValue: reflect.ValueOf([]string{"Hello", "my", "friend", "?", "How", "are", "you", "?"}),
		Len:        8,
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Equal(t, output.Interface(), "?")
}

func TestLastPtr_Run(t *testing.T) {
	valTrue := true
	valFalse := false
	input := Last{
		ItemsValue: reflect.ValueOf([]*bool{&valTrue, &valTrue, &valFalse, &valFalse, &valTrue}),
		Len:        5,
	}
	output, err := input.Run()
	assert.Nil(t, err)
	assert.Equal(t, output.Interface(), &valTrue)
}

func TestLastError_Run(t *testing.T) {
	input := Last{
		ItemsValue: reflect.ValueOf([]*bool{}),
		Len:        0,
	}
	output, err := input.Run()
	assert.Equal(t, reflect.ValueOf(nil), output)
	assert.NotNil(t, err)
}
