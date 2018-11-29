package contains

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestContains_validateSuccessPtr(t *testing.T) {
	contains := &Contains{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	err := contains.validate()
	assert.Nil(t, err)
}

func TestContains_validateSuccessValue(t *testing.T) {
	contains := &Contains{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   10,
	}
	err := contains.validate()
	assert.Nil(t, err)
}

func TestContains_validateInvalidArgument(t *testing.T) {
	contains := &Contains{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   utils.IntPtr(10),
	}
	err := contains.validate()
	assert.NotNil(t, err)
	assert.Equal(t, "[contains:err.invalid-argument] The Stream contains elements of type int and the passed argument has type *int", err.Error())
}

func TestContains_Run(t *testing.T) {
	contains := &Contains{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	found, err := contains.Run()
	assert.Nil(t, err)
	assert.True(t, found)

	person := utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	contains = &Contains{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	found, err = contains.Run()
	assert.Nil(t, err)
	assert.True(t, found)


	person = utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	contains = &Contains{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 20, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	found, err = contains.Run()
	assert.Nil(t, err)
	assert.False(t, found)

	personPtr := &utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	contains = &Contains{
		Items: reflect.ValueOf([]*utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(personPtr),
		Element:   personPtr,
	}
	found, err = contains.Run()
	assert.Nil(t, err)
	assert.True(t, found)
}
