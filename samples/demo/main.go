package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type Person struct {
	Fullname string
	Age      int
}

var people = []Person{
	{"John Doe", 25},
	{"Jane Doe", 25},
	{"Tom Doe", 22},
	{"David Doe", 15},
}

func main() {
	v,_:=koazee.StreamOf(people).GroupBy(func(p Person) int{
		return p.Age
	})

	fmt.Println(v.(map[int][]Person))
}