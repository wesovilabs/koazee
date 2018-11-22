+++
title = "stream.Compose"
description = "Join 2 or more streams in a single one"
weight = 23
draft = false
toc = true
bref ="Join 2 or more streams in a single one"
+++

<h3 class="section-head" id="h-signature"><a href="#h-signature">Function signature</a></h3>
{{< highlight golang >}}
    func Compose(streams []stream.S) (stream stream.S) 
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
        <td>streams</td>
        <td>Array of Streams</td>
        <td>List of streams to be joined</td>
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
        <td>strea,</td>
        <td>stream.S</td>
        <td>A new stream that will contain both the items and operations in all the the joined streams</td>
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
        <td>err.invalid-argument</td>
        <td>To compose a new stream, 2 or more streams mut be passed as argument</td>
      </tr>
      <tr>
        <td>err.invalid-argument</td>
        <td>An stream can not be composed with stream of different type</td>
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

var numbers = []int{1, 2, 3, 4}
var numbers2 = []int{5, 6, 7, 8}

func main() {
	stream1 := koazee.Stream().
		With(numbers).
		Map(func(val int) int {
			return val * 2
		})
	stream2 := koazee.StreamOf(numbers2)
	fmt.Println(koazee.Stream().Compose(stream1, stream2).Reduce(func(acc,val int)int{return acc+val}).Int())

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
		fmt.Println("Hi there, this is %s", primate.name)
	}).Out()
}
{{< /highlight >}}
</div>
