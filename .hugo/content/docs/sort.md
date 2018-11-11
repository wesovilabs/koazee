+++
title = "stream.Sort"
description = "Sort the elements in the stream"
weight = 21
draft = false
toc = true
bref ="Sort the elements in the stream"
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Sor(fn interface{})
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
        <td>item</td>
        <td>func</td>
        <td>This function must receive two arguments with the same tup that the elements in the stream and must return an int (-1,0,1) taht
        will determine which of the two items must be ordered before the other</td>
      </tr>
    </tbody>
</table>
<h4>Output</h4>
None

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
        <td>A nil stream can not be sorted</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The filter operation requires a function as argument</td>
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
        <td>The type of the both arguments must be  %s</td>
      </tr>         
      <tr>
        <td>err.invalid-argument</td>
        <td>The type of the output in the provided function must be int</td>
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
	fmt.Println(koazee.Stream().
		With(numbers).
		Sort(func(val1, val2 int) int {
			if val1 > val2 {
				return 1
			} else if val1 < val2 {
				return -1
			}
			return 0
		}).Out().Val())
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
	primates := koazee.StreamOf(primates).
		Sort(func(primate1, primate2 *primate) int {
			if primate1.age > primate2.age {
				return -1
			} else if primate1.age < primate2.age {
				return 1
			}
			return 0
		}).Out().Val().([]*primate)
	for _, primate := range primates {
		fmt.Println(primate.name)
	}

}
{{< /highlight >}}
</div>
