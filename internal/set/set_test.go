package set

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/errors"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestSet_RunSuccess_primitives(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(1),
		Index:      1,
		Len:        3,
	}
	value, err := set.Run()
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, 10, value.Interface().([]int)[1])
}

func TestSet_RunSuccess_primitivesPtr(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  utils.IntPtr(10),
		ItemsValue: reflect.ValueOf([]*int{utils.IntPtr(2), utils.IntPtr(3), utils.IntPtr(10)}),
		ItemsType:  reflect.TypeOf(utils.IntPtr(10)),
		Index:      1,
		Len:        3,
	}
	value, err := set.Run()
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, 10, *value.Interface().([]*int)[1])
}

func TestSet_RunSuccess_Success_structure(t *testing.T) {
	cache = cacheType{}
	person := utils.Person{"John", "Doe", 20, true}
	people := []utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	set := &Set{
		ItemValue:  person,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
		Index:      1,
		Len:        2,
	}
	value, err := set.Run()
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, person, value.Interface().([]utils.Person)[1])
}

func TestSet_RunSuccess_structurePtr(t *testing.T) {
	cache = cacheType{}
	person := &utils.Person{"John", "Doe", 20, true}
	people := []*utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	set := &Set{
		ItemValue:  person,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
		Index:      1,
		Len:        2,
	}
	value, err := set.Run()
	assert.Nil(t, err)
	assert.NotNil(t, value)
	assert.Equal(t, person, value.Interface().([]*utils.Person)[1])
}

func TestSet_RunError_invalidArgument(t *testing.T) {
	cache = cacheType{}
	newElement := 10
	people := []*utils.Person{{"Jane", "Doe", 18, false}, {"Tom", "Doe", 1, true}}
	set := &Set{
		ItemValue:  newElement,
		ItemsValue: reflect.ValueOf(people),
		ItemsType:  reflect.TypeOf(people).Elem(),
		Index:      1,
		Len:        2,
	}
	value, err := set.Run()
	assert.Equal(t, reflect.ValueOf(nil), value)
	assert.NotNil(t, err)
	assert.Equal(t, "[set:err.invalid-argument] The type of the argument must be *utils.Person", err.Error())
}

func TestSet_validateSuccessPtr(t *testing.T) {
	cache = cacheType{}
	val := 5
	set := &Set{
		ItemValue:  nil,
		ItemsValue: reflect.ValueOf([]*int{&val}),
		ItemsType:  reflect.TypeOf(&val),
		Index:      0,
		Len:        1,
	}
	info, err := set.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, set.ItemsType, *info.itemType)
}

func TestSet_validateErrorInvalidArgument(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf("hello"),
		Index:      1,
		Len:        3,
	}
	info, err := set.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidArgument.String(), err.Code())
	assert.Equal(t, "[set:err.invalid-argument] The type of the argument must be string", err.Error())
}

func TestSet_validateErrorInvalidIndex(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(10),
		Index:      4,
		Len:        3,
	}
	info, err := set.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidIndex.String(), err.Code())
	assert.Equal(t, "[set:err.invalid-index] The length of this Stream is 3, so the index must be between 0 and 2", err.Error())
}

func TestSet_validateErrorInvalidIndexNegative(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  10,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(10),
		Index:      -1,
		Len:        3,
	}
	info, err := set.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidIndex.String(), err.Code())
	assert.Equal(t, "[set:err.invalid-index] The length of this Stream is 3, so the index must be between 0 and 2", err.Error())
}

func TestSet_validateErrorNilValueInNonPointerStream(t *testing.T) {
	cache = cacheType{}
	set := &Set{
		ItemValue:  nil,
		ItemsValue: reflect.ValueOf([]int{5, 10, 15}),
		ItemsType:  reflect.TypeOf(0),
		Index:      0,
		Len:        3,
	}
	info, err := set.validate()
	assert.Nil(t, info)
	assert.NotNil(t, err)
	assert.Equal(t, OpCode, err.Operation())
	assert.Equal(t, errors.ErrInvalidArgument.String(), err.Code())
	assert.Equal(t, "[set:err.invalid-argument] The type of the argument must be int", err.Error())
}

func TestSet_validateSuccessPtrCache(t *testing.T) {
	cache = cacheType{}
	val := 5
	set := &Set{
		ItemValue:  nil,
		ItemsValue: reflect.ValueOf([]*int{&val}),
		ItemsType:  reflect.TypeOf(&val),
		Index:      0,
		Len:        1,
	}
	info, err := set.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, set.ItemsType, *info.itemType)
	info, err = set.validate()
	assert.Nil(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, set.ItemsType, *info.itemType)
}
