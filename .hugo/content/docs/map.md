+++
title = "stream.Map"
description = "Convert the current elements in the stream into a different type"
weight = 19
draft = false
toc = true
bref = "Convert the current elements in the stream into a different type"
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Map(fn interface{}) output
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
        <td>This function must receive 1 arguments, that must have the same type that current elements in the stream. The output of this function will be the new type op elements in the stream</td>
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
        <td>A nil stream can not be iterated</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The map operation requires a function as argument</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The provided function must retrieve 1 argument</td>
      </tr>      
      <tr>
        <td>err.invalid-argument</td>
        <td>The provided function must return 1 value</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>The type of the argument in the provided function must be %s</td>
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
	"strings"
)

var numbers = []int{1, 3, 5, 7, 9}

func main() {
	output := koazee.Stream().
		With(numbers).
		Map(func(item int) string {
			return strings.Repeat("a", item)
		}).Out().Val().([]string)
	fmt.Printf("The output is %v\n", output)
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
	ages := koazee.StreamOf(primates).
		Map(func(primate *primate) int {
			return primate.age
		}).Out().Val()

	fmt.Printf("The ages of the primates is %v years\n", ages)

}

{{< /highlight >}}
</div>
