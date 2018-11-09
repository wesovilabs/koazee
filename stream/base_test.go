package stream_test

import (
	"os"
	"testing"

	"github.com/wesovilabs/koazee/stream"
)

type person struct {
	firstName string
	age       int
}

var numbersArray = []int{2, 10, 3, 2}
var numberPointersArray = []*int{intPtr(2),
	intPtr(10), intPtr(3), intPtr(2)}
var textsArray = []string{"home", "welcome", "software", "tech", "geek"}
var textPointersArray = []*string{stringPtr("home"),
	stringPtr("welcome"), stringPtr("software"),
	stringPtr("tech"), stringPtr("geek")}
var booleansArray = []bool{false, false, true, true}
var booleanPointersArray = []*bool{booleanPtr(false),
	booleanPtr(false), booleanPtr(true), booleanPtr(true)}
var multiTypesArray = []interface{}{false, "home", true, 10, nil}
var multiTypePointersArray = []interface{}{booleanPtr(false),
	stringPtr("home"), booleanPtr(true), intPtr(10), nil}
var structsArray = []person{
	{
		firstName: "John",
		age:       25,
	},
	{
		firstName: "Jane",
		age:       50,
	},
	{
		firstName: "Jean",
		age:       13,
	},
}

var structPointersArray = []*person{
	{
		firstName: "John",
		age:       25,
	},
	{
		firstName: "Jane",
		age:       50,
	},
	{
		firstName: "Jean",
		age:       13,
	},
}
var numberStream = stream.New(numbersArray)
var numberPointerStream = stream.New(numberPointersArray)
var textStream = stream.New(textsArray)
var textPointerStream = stream.New(textPointersArray)
var booleanStream = stream.New(booleansArray)
var booleanPointerStream = stream.New(booleanPointersArray)
var multiTypesStream = stream.New(multiTypesArray)
var multiTypesPointerStream = stream.New(multiTypePointersArray)
var structStream = stream.New(structsArray)
var structPointerStream = stream.New(structPointersArray)

func booleanPtr(value bool) *bool {
	val := value
	return &val
}

func stringPtr(value string) *string {
	val := value
	return &val
}

func intPtr(value int) *int {
	val := value
	return &val
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
