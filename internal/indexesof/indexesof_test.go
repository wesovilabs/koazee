package indexesof

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
)

func TestIndexesOf_validateSuccessPtr(t *testing.T) {
	indexesOf := &IndexesOf{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	err := indexesOf.validate()
	assert.Nil(t, err)
}

func TestIndexesOf_validateSuccessValue(t *testing.T) {
	indexesOf := &IndexesOf{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   10,
	}
	err := indexesOf.validate()
	assert.Nil(t, err)
}

func TestIndexesOf_validateInvalidArgument(t *testing.T) {
	indexesOf := &IndexesOf{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   utils.IntPtr(10),
	}
	err := indexesOf.validate()
	assert.NotNil(t, err)
	assert.Equal(t, "[indexesOf:err.invalid-argument] The Stream indexesOf elements of type int and the passed argument has type *int", err.Error())
}

func TestIndexesOf_Run(t *testing.T) {
	indexesOf := &IndexesOf{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	index, err := indexesOf.Run()
	assert.Nil(t, err)
	assert.Equal(t, []int{2}, index)

	person := utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexesOf = &IndexesOf{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	index, err = indexesOf.Run()
	assert.Nil(t, err)
	assert.Equal(t, []int{0}, index)

	person = utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexesOf = &IndexesOf{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 20, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	index, err = indexesOf.Run()
	assert.NotNil(t, err)
	assert.Equal(t, []int{}, index)

	personPtr := &utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexesOf = &IndexesOf{
		Items: reflect.ValueOf([]*utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(personPtr),
		Element:   personPtr,
	}
	index, err = indexesOf.Run()
	assert.Nil(t, err)
	assert.Equal(t, []int{0}, index)

}
