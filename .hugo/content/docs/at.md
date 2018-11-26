+++
title = "stream.At"
description = "Returns the element in the stream in the given position"
weight = 12
draft = false
toc = true
bref = "Returns the element in the stream in the given position"
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func At(index int) (out output)
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
        <td>index</td>
        <td>int</td>
        <td>Index in the stream, from 0-len(stream)-1</td>
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
        <td>It contains both value and error</td>
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
        <td>It can not be taken an element from a nil stream</td>
      </tr>
      <tr>
        <td>err.items-nil</td>
        <td>It can not be taken an element from an empty stream</td>
      </tr>
      <tr>
        <td>err.invalid-index</td>
        <td>he length of this stream is %d, so the index must be  between 0 and %d</td>
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
	out := koazee.Stream().
		With(numbers).
		At(2)
	fmt.Printf("%d\n", out.Val())

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
	primate := koazee.StreamOf(primates).
		At(3).Val().(*primate)

	fmt.Printf("%s was invited to the party\n", primate.name)

}
{{< /highlight >}}
</div>
