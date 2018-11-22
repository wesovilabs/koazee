+++
title = "stream.Contains"
description = "Check if the given element is found in the stream"
weight = 15
draft = false
toc = true
bref = "Check if the given element is found in the stream"
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Contains(item interface{}) (found bool,err *errors.Error)
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
        <td>Item to be searched in the stream</td>
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
        <td>found</td>
        <td>bool</td>
        <td>Returns true if element is found and false if not</td>
      </tr>
      <tr>
        <td>err</td>
        <td>*errors.Error</td>
        <td>Returns nil when the operation was fine and error when something didn't work</td>
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
        <td>It can not be checked if an element is in a nil stream</td>
      </tr
      <tr>
        <td>err.invalid-argument</td>
        <td>It can not be checked if an array of non-pointers contains a nil value</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The stream contains elements of type %s and the passed argument has type %s</td>
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
	found, _ := koazee.Stream().
		With(numbers).
		Contains(5)
	if found {
		fmt.Println("The element was found!")
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
	searched:=newPrimate("Sarah", 7, "Gibbon", female)
	found,_:=koazee.StreamOf(primates).
		Contains(searched)
	if found{
		fmt.Printf("%s is in the list of monkeys invited to the party\n", searched.name)
	}

}
{{< /highlight >}}
</div>
