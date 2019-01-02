package indexesof

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func TestContains_equals(t *testing.T) {
	assert.True(t, equalsValues(reflect.ValueOf(utils.IntPtr(3)), reflect.ValueOf(utils.IntPtr(3))))
	assert.False(t, equalsValues(reflect.ValueOf(utils.IntPtr(3)), reflect.ValueOf(3)))
	assert.True(t, equalsValues(reflect.ValueOf("hola"), reflect.ValueOf("hola")))
	assert.False(t, equalsValues(reflect.ValueOf("hola"), reflect.ValueOf("HOla")))
}
