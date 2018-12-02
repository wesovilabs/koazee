package sort_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee/internal/sort"
	"github.com/wesovilabs/koazee/utils"
	"reflect"
	"strings"
	"testing"
)

func testSortArrayOfPeople(t *testing.T) {
	items := []utils.Person{
		{FirstName: "John", LastName: "Doe", Age: 34, Male: true},
		{FirstName: "Jane", LastName: "Doe", Age: 17, Male: false},
		{FirstName: "Jim", LastName: "Garcia", Age: 20, Male: true},
		{FirstName: "Jamie", LastName: "Garcia", Age: 12, Male: false},
	}
	sort := sort.Sort{
		ItemsValue: reflect.ValueOf(items),
		ItemsType:  reflect.TypeOf(utils.Person{}),
		Func: func(a, b utils.Person) int {
			return strings.Compare(a.FirstName, b.FirstName)
		},
	}
	values, err := sort.Run()
	assert.Nil(t, err)
	assert.NotNil(t, values)
	assert.Equal(t, "Jamie", values.Interface().([]utils.Person)[0].FirstName)
	assert.Equal(t, "Jane", values.Interface().([]utils.Person)[1].FirstName)
	assert.Equal(t, "Jim", values.Interface().([]utils.Person)[2].FirstName)
	assert.Equal(t, "John", values.Interface().([]utils.Person)[3].FirstName)
}

func testSortArrayOfPtrInt(t *testing.T) {
	items := []*int{
		utils.IntPtr(10),
		utils.IntPtr(1),
		utils.IntPtr(21),
		utils.IntPtr(3),
		utils.IntPtr(30),
	}
	sort := sort.Sort{
		ItemsValue: reflect.ValueOf(items),
		ItemsType:  reflect.TypeOf(utils.IntPtr(2)),
		Func: func(a, b *int) int {
			if *a <= *b {
				return -1
			}
			return 1
		},
	}
	values, err := sort.Run()
	assert.Nil(t, err)
	assert.NotNil(t, values)
	assert.Equal(t, 1, *(values.Interface().([]*int)[0]))
	assert.Equal(t, 3, *(values.Interface().([]*int)[1]))
	assert.Equal(t, 10, *(values.Interface().([]*int)[2]))
	assert.Equal(t, 21, *(values.Interface().([]*int)[3]))
	assert.Equal(t, 30, *(values.Interface().([]*int)[4]))

}

func testSortArrayOfStringInvalidFunctionArgType(t *testing.T) {
	items := []string{"a", "b", "h", "d"}
	sort := sort.Sort{
		ItemsValue: reflect.ValueOf(items),
		ItemsType:  reflect.TypeOf("a"),
		Func: func(a, b *int) int {
			if *a <= *b {
				return -1
			}
			return 1
		},
	}
	values, err := sort.Run()
	assert.Equal(t, reflect.ValueOf(nil), values)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "[sort:err.invalid-argument] The type of the both arguments must be  string")
}
func testSortArrayOfStringInvalidFunctionOutType(t *testing.T) {
	items := []string{"a", "b", "h", "d"}
	sort := sort.Sort{
		ItemsValue: reflect.ValueOf(items),
		ItemsType:  reflect.TypeOf("a"),
		Func: func(a, b string) bool {
			return true
		},
	}
	values, err := sort.Run()
	assert.Equal(t, reflect.ValueOf(nil), values)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "[sort:err.invalid-argument] The type of the Output in the provided function must be int")
}

func testSortArrayOfStringInvalidFunctionArgNumbers(t *testing.T) {
	items := []string{"a", "b", "h", "d"}
	sort := sort.Sort{
		ItemsValue: reflect.ValueOf(items),
		ItemsType:  reflect.TypeOf("a"),
		Func: func(a, b, c string) int {
			if a <= b {
				return -1
			}
			return len(c)
		},
	}
	values, err := sort.Run()
	assert.Equal(t, reflect.ValueOf(nil), values)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "[sort:err.invalid-argument] The provided function must retrieve 2 arguments")
}

func TestSort_Run(t *testing.T) {
	testSortArrayOfPeople(t)
	testSortArrayOfPtrInt(t)
	testSortArrayOfStringInvalidFunctionArgType(t)
	testSortArrayOfStringInvalidFunctionOutType(t)
	testSortArrayOfStringInvalidFunctionArgNumbers(t)
}

