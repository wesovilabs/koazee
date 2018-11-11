package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type genre int

const (
	female genre = iota
	male
)

type primate struct {
	name   string
	age    int
	family string
	genre  genre
}

func newPrimate(name string, age int, family string, genre genre) *primate {
	return &primate{
		name:   name,
		age:    age,
		family: family,
		genre:  genre,
	}
}

var primates = []*primate{
	newPrimate("John", 15, "Capuchin", male),
	newPrimate("Laura", 12, "Squirrel monkey", female),
}

var primates2 = []*primate{
	newPrimate("Benjamin", 23, "Spider monkey", male),
	newPrimate("George", 19, "Golden Lion Tamarin", male),
	newPrimate("Jane", 33, "Orangutan", female),
	newPrimate("Sarah", 7, "Gibbon", female),
}

func main() {
	stream1 := koazee.StreamOf(primates).Filter(func(primate *primate) bool {
		return primate.genre == male
	})
	stream2 := koazee.StreamOf(primates2).Filter(func(primate *primate) bool {
		return primate.genre == male
	})
	koazee.Stream().Compose(stream1, stream2).ForEach(func(primate *primate) {
		fmt.Printf("Hi there, this is %s\n", primate.name)
	}).Out()
}
