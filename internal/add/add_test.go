package add

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)


func TestAdd_RunSuccess_primitives(t *testing.T) {

	add := &Add{
		Item:       10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(1),
	}
	info, err := add.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, add.ItemsType, *info.itemType)
}

func TestAdd_RunSuccess_primitivesPtr(t *testing.T) {
	cache = cacheType{}
	add := &Add{
		Item:       utils.IntPtr(10),
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(10)}),
		ItemsType:  reflect.TypeOf(utils.IntPtr(10)),
	}
	info, err := add.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, add.ItemsType, *info.itemType)
}

func TestAdd_RunSuccess_Success_structure(t *testing.T) {
	cache = cacheType{}
	person := utils.Person{"John", "Doe", 20, true}
	people := []utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	add := &Add{
		Item:       person,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
	}
	info, err := add.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, add.ItemsType, *info.itemType)
}

func TestAdd_RunSuccess_structurePtr(t *testing.T) {
	cache = cacheType{}
	person := &utils.Person{"John", "Doe", 20, true}
	people := []*utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	add := &Add{
		Item:       person,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
	}
	info, err := add.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, add.ItemsType, *info.itemType)
}

func TestAdd_RunError_invalidArgument(t *testing.T) {
	cache = cacheType{}
	newElement := 10
	people := []*utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	add := &Add{
		Item:       newElement,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
	}
	info, err := add.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, "[add:err.invalid-argument] An element whose type is int can not be added in a Stream of type *utils.Person", err.Error())
}

func TestAdd_validateSuccessPtr(t *testing.T) {
	cache = cacheType{}
	val := 5
	add := &Add{
		Item:       nil,
		ItemsValue: reflect.ValueOf([]*int{&val}),
		ItemsType:  reflect.TypeOf(&val),
	}
	info, err := add.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, add.ItemsType, *info.itemType)
}

func TestAdd_validateErrorInvalidArgument(t *testing.T) {
	cache = cacheType{}
	add := &Add{
		Item:       10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf("hello"),
	}
	info, err := add.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidArgument.String(), err.Code())
	assert.Equal(t, "[add:err.invalid-argument] An element whose type is int can not be added in a Stream of type string", err.Error())
}

func TestAdd_validateErrorNilValueInNonPointerStream(t *testing.T) {
	cache = cacheType{}
	add := &Add{
		Item:       nil,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(0),
	}
	info, err := add.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidArgument.String(), err.Code())
	assert.Equal(t, "[add:err.invalid-argument] A nil value can not be added in a Stream of non-pointers values", err.Error())
}
