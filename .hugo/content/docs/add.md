+++
title = "stream.Add"
description = "Add a new element in the stream"
weight = 13
draft = false
toc = true
bref = "Add a new element into the stream."
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Add(item interface{}) (stream S)
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
        <td>Same type of elements in the stream</td>
        <td>New item to be added into the stream</td>
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
        <td>stream</td>
        <td>koazee.S</td>
        <td>It returns the stream modified by the current operation</td>
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
        <td>An element can not be added in a nil stream</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>A nil value can not be added in a stream of non-pointers values</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>An element whose type is %s can not be added in a stream of type %s</td>
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
	newList := koazee.Stream().
		Add(10).
		With(numbers).
		Out().
		Val().([]int)

	for _, number := range newList {
		fmt.Printf("%d\n", number)
	}
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
	newList := koazee.StreamOf(primates).
		Add(newPrimate("Pepe", 16, "Gibbon", male)).
		Out().
		Val().([]*primate)

	for _, primate := range newList {
		fmt.Printf("%s was invited to the party\n", primate.name)
	}
}

{{< /highlight >}}
</div>
