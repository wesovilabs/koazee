package at

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestAt_Run(t *testing.T) {
	at := At{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4), utils.IntPtr(40)}),
		Len:        2,
		Index:      1,
	}
	val, err := at.Run()
	assert.Nil(t, err)
	assert.NotEqual(t, reflect.Zero(reflect.TypeOf(utils.IntPtr(5))), val)
	assert.Equal(t, 40, val.Elem().Interface())
}

func TestAt_validate(t *testing.T) {
	at := At{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4)}),
		Len:        10,
		Index:      1,
	}
	err := at.validate()
	assert.Nil(t, err)
}

func TestAt_validateInvalidIndex(t *testing.T) {
	at := At{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4)}),
		Len:        10,
		Index:      10,
	}
	err := at.validate()
	assert.NotNil(t, err)
	assert.Equal(t, "[at:err.invalid-index] The length of this Stream is 10, so the index must be between 0 and 9", err.Error())
}

func TestAt_validateLen0(t *testing.T) {
	at := At{
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(4)}),
		Len:        0,
		Index:      10,
	}
	err := at.validate()
	assert.NotNil(t, err)
	assert.Equal(t, "[at:err.items-nil] It can not be taken an element from an empty Stream", err.Error())
}
