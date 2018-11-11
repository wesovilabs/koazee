+++
title = "stream.Drop"
description = "Drop an element from the stream"
weight = 11
draft = false
toc = true
bref = "Drop an element from the stream."
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>

{{< highlight golang >}}
    func Drop(item interface{}) (stream S)
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
        <td>Element to be dropped from the stream</td>
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
        <td>An element can not be dropped from a nil stream</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>A nil value can not be dropped from a stream of non-pointers values</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>An element whose type is %s can not be dropped from a stream of type %s</td>
      </tr>
    </tbody>
</table>
<h3 class="section-head" id="h-examples"><a href="#h-examples">Examples</a></h3>
<nav class="tabs" data-component="tabs">
    <ul>
      <li>
        <a href="#numbers">Numbers</a>
      </li>
      <li class="active">
        <a href="#struct_pointers">Pointers</a>
      </li>
    </ul>
</nav>
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
		Drop(newPrimate("Benjamin", 23, "Spider monkey", male)).
		Out().
		Val().([]*primate)

	for _, primate := range newList {
		fmt.Printf("%s was invited to the party\n", primate.name)
	}
}
{{< /highlight >}}
</div>
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
		Drop(3).
		With(numbers).
		Out().
		Val().([]int)

	for _, number := range newList {
		fmt.Printf("%d\n", number)
	}
}

{{< /highlight >}}