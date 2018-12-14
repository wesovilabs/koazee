package indexof

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"testing"
)

func TestIndexOf_validateSuccessPtr(t *testing.T) {
	indexOf := &IndexOf{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	err := indexOf.validate()
	assert.Nil(t, err)
}

func TestIndexOf_validateSuccessValue(t *testing.T) {
	indexOf := &IndexOf{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   10,
	}
	err := indexOf.validate()
	assert.Nil(t, err)
}

func TestIndexOf_validateInvalidArgument(t *testing.T) {
	indexOf := &IndexOf{
		Items:     reflect.ValueOf([]int{1, 10, 7}),
		ItemsType: reflect.TypeOf(1),
		Element:   utils.IntPtr(10),
	}
	err := indexOf.validate()
	assert.NotNil(t, err)
	assert.Equal(t, "[indexOf:err.invalid-argument] The Stream indexOf elements of type int and the passed argument has type *int", err.Error())
}

func TestIndexOf_Run(t *testing.T) {
	indexOf := &IndexOf{
		Items:     reflect.ValueOf([]*int{utils.IntPtr(1), utils.IntPtr(7), utils.IntPtr(10)}),
		ItemsType: reflect.TypeOf(utils.IntPtr(1)),
		Element:   utils.IntPtr(10),
	}
	index, err := indexOf.Run()
	assert.Nil(t, err)
	assert.Equal(t,2, index)

	person := utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexOf = &IndexOf{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	index, err = indexOf.Run()
	assert.Nil(t, err)
	assert.Equal(t,0, index)


	person = utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexOf = &IndexOf{
		Items: reflect.ValueOf([]utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 20, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(person),
		Element:   person,
	}
	index, err = indexOf.Run()
	assert.NotNil(t, err)
	assert.Equal(t,InvalidIndex, index)

	personPtr := &utils.Person{FirstName: "John", LastName: "Doe", Age: 25, Male: true}
	indexOf = &IndexOf{
		Items: reflect.ValueOf([]*utils.Person{
			{FirstName: "John", LastName: "Doe", Age: 25, Male: true},
			{FirstName: "Jim", LastName: "Doe", Age: 20, Male: true},
		}),
		ItemsType: reflect.TypeOf(personPtr),
		Element:   personPtr,
	}
	index, err = indexOf.Run()
	assert.Nil(t, err)
	assert.Equal(t,0, index)

}
