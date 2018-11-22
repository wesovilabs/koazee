+++
title = "stream.Reduce"
description = "Reduces the stream to a single value by executing a provided function for each value of the stream"
weight = 18
draft = false
toc = true
bref = "Reduces the stream to a single value by executing a provided function for each value of the stream "
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Reduce(fn interface{}) output
{{< /highlight >}}

<h4>Arguments</h4>
<table>
    <thead>
        <tr>
        <th>Name</th>
        <th>Type</th>
        <th>Description</th>
        </tr>
    </thead>
    <tbody>
      <tr>
        <td>fn</td>
        <td>func</td>
        <td>This function must receive two arguments, the first argument is an accumulator and the second a element of the same
        type that the elements in the stream. The output of this function will have the same type than the accumulator</td>
      </tr>
    </tbody>
</table>
<h4>Output</h4>
<table>
    <thead>
        <tr>
        <th>Name</th>
        <th>Type</th>
        <th>Description</th>
        </tr>
    </thead>
    <tbody>
      <tr>
        <td>output</td>
        <td>stream.output</td>
        <td>It returns the output (value and error)</td>
      </tr>
    </tbody>
</table>

<h3 class="section-head" id="h-errors"><a href="#h-errors">Errors</a></h3>
<table>
    <thead>
        <tr>
        <th>Type</th>
        <th>Message</th>
        </tr>
    </thead>
    <tbody>
      <tr>
        <td>err.items-nil</td>
        <td>A nil stream can not be reduced</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The reduce operation requires a function as argument</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The provided function must retrieve 2 arguments</td>
      </tr>      
      <tr>
        <td>err.invalid-argument</td>
        <td>The provided function must return 1 value</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The type of the second argument in the provided function must be %s</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The type of the first argument and the output in the provided function must be the same</td>
      </tr>
    </tbody>
</table>
<h3 class="section-head" id="h-examples"><a href="#h-examples">Examples</a></h3>
<nav class="tabs" data-component="tabs">
    <ul>
      <li class="active">
        <a href="#numbers">Numbers</a>
      </li>
      <li>
        <a href="#struct_pointers">Pointers</a>
      </li>
    </ul>
</nav>
<div id="numbers">
{{< highlight golang >}}
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 3, 5, 7, 9}

func main() {
	avg := koazee.Stream().
		With(numbers).
		Reduce(func(acc, item int) int {
			return acc  + item
		}).Int()
	fmt.Printf("The average is %d\n", avg)
}
{{< /highlight >}}
</div>
<div id="struct_pointers">
{{< highlight golang >}}
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
	newPrimate("Benjamin", 23, "Spider monkey", male),
	newPrimate("George", 19, "Golden Lion Tamarin", male),
	newPrimate("Jane", 33, "Orangutan", female),
	newPrimate("Sarah", 7, "Gibbon", female),
}

func main() {
	filteredPrimates := koazee.StreamOf(primates).
		Filter(func(primate *primate) bool {
			return primate.age > 10 && primate.genre == female
		}).Out().Val().([]*primate)
	for _, primate := range filteredPrimates {
		fmt.Printf("%s is a female and is %d years\n", primate.name, primate.age)
	}

}
{{< /highlight >}}
</div>
