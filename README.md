[![Build Status](https://travis-ci.org/wesovilabs/koazee.svg?branch=master)](https://travis-ci.org/wesovilabs/koazee)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/koazee)](https://goreportcard.com/report/github.com/wesovilabs/koazee)
[![godoc](https://godoc.org/github.com/wesovilabs/koazee?status.svg)](http://godoc.org/github.com/wesovilabs/koazee)
[![codecov](https://codecov.io/gh/wesovilabs/koazee/branch/master/graph/badge.svg)](https://codecov.io/gh/wesovilabs/koazee)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go#Utilities)

# Koazee

> Lazy like a koala, smart like a chimpanzee


## What is Koazee?

Koazee is a StreamLike, Immutable, Lazy Loading and smart Golang Library to deal with slices. 


Visit the [Koazee wiki](https://github.com/wesovilabs/koazee/wiki) to find out what Koazee can do.

## Koazee Highlights

- **Immutable**: Koazee won't modify your inputs. 
- **StreamLike**: We can combine operations up our convenience. 
- **Lazy loading**: Operations are not performed until they're required
- **Generic**: Koazee provides a generic interface able to deal with slice of any type without creating custom functions.
- **Focusing on performance**: First rule for implementing a new operation is providing the best possible performance.



## Getting started

### Installing
> Add Koazee to your project

**Go modules**

```
module github.com/me/project
require ( 
  github.com/wesovilabs/koazee vX.Y.Z
)
```

**Glide**

```
glide get github.com/wesovilabs/koazee
```

**Go dep**

```
go get github.com/wesovilabs/koazee
```

### Usage

#### Stream creation

Let's first obtain a stream from an existing array.

```golang
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("slice: %v\n", numbers)
	stream := koazee.StreamOf(numbers)
	fmt.Printf("stream: %v\n", stream.Out().Val())
}

/**
go run main.go

slice: [1 5 4 3 2 7 1 8 2 3]
stream: [1 5 4 3 2 7 1 8 2 3]
*/
```

#### Stream operations

Current release v0.0.3 (Gibbon) brings us 20 generic operations that are showed below

##### stream.At / stream.First / stream.Last
These operations return an element from the stream

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	stream := koazee.StreamOf(numbers)
	fmt.Printf("stream.At(4): %d\n", stream.At(4).Int())
	fmt.Printf("stream.First: %d\n", stream.First().Int())
	fmt.Printf("stream.Last: %d\n", stream.Last().Int())
}

/**
go run main.go

stream.At(4): 2
stream.First: 1
stream.Last: 3
*/
```

##### stream.Add / stream.Drop / stream.DropWhile / stream.DeleteAt / stream.Pop / stream.Set
These operations add or delete elements from the stream.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)

	stream := koazee.StreamOf(numbers)
	fmt.Print("stream.Add(10): ")
	fmt.Println(stream.Add(10).Do().Out().Val())

	fmt.Print("stream.Drop(5): ")
	fmt.Println(stream.Drop(5).Do().Out().Val())
	
	fmt.Print("stream.DropWhile(val<=5): ")
	fmt.Println(stream.DropWhile(func(element int)bool{return element<=5}).Do().Out().Val())

	fmt.Print("stream.DeleteAt(4): ")
	fmt.Println(stream.DeleteAt(4).Do().Out().Val())

	fmt.Print("stream.Set(0,5): ")
	fmt.Println(stream.Set(0, 5).Do().Out().Val())

	fmt.Print("stream.Pop(): ")
	val, newStream := stream.Pop()
	fmt.Printf("%d ... ", val.Int())
	fmt.Println(newStream.Out().Val())

}

/**
go run main.go

input: [1 5 4 3 2 7 1 8 2 3]
stream.Add(10): [1 5 4 3 2 7 1 8 2 3 10]
stream.Drop(5): [1 4 3 2 7 1 8 2 3]
stream.DropWhile(val<=5): [7 8]
stream.DeleteAt(4): [1 5 4 3 7 1 8 2 3]
stream.Set(0,5): [5 5 4 3 2 7 1 8 2 3]
stream.Pop(): 1 ... [5 4 3 2 7 1 8 2 3]
*/
```

##### tream.Count / stream.IndexOf / stream.IndexesOf / stream.LastIndexOf / stream.Contains
These operations return info from the elements in the stream

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)
	stream := koazee.StreamOf(numbers)
	count, _ := stream.Count()
	fmt.Printf("stream.Count(): %d\n", count)
	index, _ := stream.IndexOf(2)
	fmt.Printf("stream.IndexOf(2): %d\n", index)
	indexes, _ := stream.IndexesOf(2)
    fmt.Printf("stream.IndexesOf(2): %d\n", indexes)
	index, _ = stream.LastIndexOf(2)
	fmt.Printf("stream.LastIndexOf(2): %d\n", index)
	contains, _ := stream.Contains(7)
	fmt.Printf("stream.Contains(7): %v\n", contains)
}

/**
go run main.go

input: [1 5 4 3 2 7 1 8 2 3]
stream.Count(): 10
stream.IndexOf(2): 4
stream.IndexesOf(2): [4 8]
stream.LastIndexOf(2): 8
stream.Contains(7): true
*/
```

##### stream.Sort / stream.Reverse
These operations organize the elements in the stream.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "fox", "tiger", "lion"}

func main() {
	fmt.Print("input: ")
	fmt.Println(animals)
	stream := koazee.StreamOf(animals)

	fmt.Print("stream.Reverse(): ")
	fmt.Println(stream.Reverse().Out().Val())

	fmt.Print("stream.Sort(strings.Compare): ")
	fmt.Println(stream.Sort(strings.Compare).Out().Val())

}

/**
go run main.go

input: [lynx dog cat monkey fox tiger lion]
stream.Reverse(): [lion tiger fox monkey cat dog lynx]
stream.Sort(strings.Compare): [cat dog fox lion lynx monkey tiger]
*/
```

##### stream.Take / stream.Filter / stream.RemoveDuplicates
These operations return a filtered stream.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Print("input: ")
	fmt.Println(animals)
	stream := koazee.StreamOf(animals)

	fmt.Print("stream.Take(1,4): ")
	fmt.Println(stream.Take(1, 4).Out().Val())

	fmt.Print("stream.Filter(len==4): ")
	fmt.Println(stream.
		Filter(
			func(val string) bool {
				return len(val) == 4
			}).
		Out().Val(),
	)
	fmt.Print("stream.RemoveDuplicates(): ")
	fmt.Println(stream.RemoveDuplicates().Out().Val())
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.Take(1,4): [dog cat monkey dog]
stream.Filter(len==4): [lynx lion]
stream.RemoveDuplicates(): [lynx dog cat monkey fox tiger lion]
*/
```
##### stream.GroupBy
This operation creates groups depending the returned value by the function

You can now optionally return an error as the second parameter to stop processing of the stream. The error will be available in `stream.Out().Err().UserError()`.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Printf("input: %v\n", animals)
	stream := koazee.StreamOf(animals)
	fmt.Print("stream.GroupBy(strings.Len): ")
	out, _ := stream.GroupBy(func(val string)int{return len(val)})
	fmt.Println(out)
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.GroupBy(strings.Len): map[5:[tiger] 4:[lynx lion] 3:[dog cat dog fox] 6:[monkey]]
*/
```


##### stream.Map
This operation performs a modification over all the elements in the stream.

You can now optionally return an error as the second parameter to stop processing of the stream. The error will be available in `stream.Out().Err().UserError()`.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Printf("input: %v\n", animals)
	stream := koazee.StreamOf(animals)
	fmt.Print("stream.Map(strings.Title): ")
	fmt.Println(stream.Map(strings.Title).Do().Out().Val())
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.Map(strings.Title): [Lynx Dog Cat Monkey Dog Fox Tiger Lion]
*/
```

##### stream.Reduce
This operation give us a single output after iterating over the elements in the stream.

You can now optionally return an error as the second parameter to stop processing of the stream. The error will be available in `stream.Out().Err().UserError()`.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)
	stream := koazee.StreamOf(numbers)
	fmt.Print("stream.Reduce(sum): ")
	fmt.Println(stream.Reduce(func(acc, val int) int {
		return acc + val
	}).Int())
}

/**
go run main.go

input: [1 5 4 3 2 7 1 8 2 3]
stream.Reduce(sum): 36
*/
```

##### stream.ForEach
This operation iterates over the element in the stream.

You can now optionally return an error as the second parameter to stop processing of the stream. The error will be available in `stream.Out().Err().UserError()`.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type message struct {
	user    string
	message string
}

var messages = []*message{
	{user: "John", message: "Hello Jane"},
	{user: "Jane", message: "Hey John, how are you?"},
	{user: "John", message: "I'm fine! and you?"},
	{user: "Jane", message: "Me too"},
}

func main() {

	stream := koazee.StreamOf(messages)
	stream.ForEach(func(m *message) {
		fmt.Printf("%s: \"%s\"\n", m.user, m.message)
	}).Do()
}

/**
go run main.go

John: "Hello Jane"
Jane: "Hey John, how are you?"
John: "I'm fine! and you?"
Jane: "Me too"
*/
```

#### Combine operations and evaluate them lazily
The main goal of Koazee is providing a set of operations that can be combined and being evaluated lazily.

```go
package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

type Person struct {
	Name string
	Male bool
	Age  int
}

var people = []*Person{
	{"John Smith", true, 32},
	{"Peter Pan", true, 17},
	{"Jane Doe", false, 20},
	{"Anna Wallace", false, 35},
	{"Tim O'Brian", true, 13},
	{"Celia Hills", false, 15},
}

func main() {
	stream := koazee.
		StreamOf(people).
		Filter(func(person *Person) bool {
			return !person.Male
		}).
		Sort(func(person, otherPerson *Person) int {
			return strings.Compare(person.Name, otherPerson.Name)
		}).
		ForEach(func(person *Person) {
			fmt.Printf("%s is %d years old\n", person.Name, person.Age)
		})

	fmt.Println("Operations are not evaluated until we perform stream.Do()\n")
	stream.Do()
}

/**
go run main.go

Operations are not evaluated until we perform stream.Do()

Anna Wallace is 35 years old
Celia Hills is 15 years old
Jane Doe is 20 years old
 */
```

## Available Operations

| Operation  | Description  | Since  |
|---|---|---|
| Add | It adds a new element in the last position |  v0.0.1 |
| At | It returns the element in the given position |  v0.0.1 |
| Contains | It checks if the given element is found in the stream.|  v0.0.1 |
| Count | It returns the number of elements in a stream|  v0.0.1 |
| DeleteAt| It remove the elements in the given position |  v0.0.3 |
| Drop | It removes an element from the stream |  v0.0.1 |
| DropWhile | It removes the elements in the stream that match with the given input function | master |
| Filter | It discards those elements that doesn't match with the provided filter|  v0.0.1 |
| First | It returns the element in the first position |  v0.0.1 |
| ForEach | It does something over all the elements in the stream.|  v0.0.1 |
| GroupBy | It returns a map of array of elements grouped by the function.|  master |
| IndexOf | It returns the first index of the  element in the stream.|  v0.0.3 |
| IndexesOf | It returns the index for all the occurrences of the  element in the stream.|  master |
| Last | It returns the element in the last position |  v0.0.1 |
| LastIndexOf | It returns the last occurrence for the  element in the stream.|  v0.0.3 |
| Map | It converts the element in the stream |  v0.0.1 |
| Pop | It extracts the first element in the stream and return this and the new stream | v0.0.3 |
| Reduce | It reduceshe stream to a single value by executing a provided function for each value of the stream|  v0.0.1 |
| RemoveDuplicates | It removes duplicated elements.|  v0.0.1 |
| Reverse| It reverses the sequence of elements in the stream.| v0.0.3 |
| Set | It replaces the element in the given index by the provided value |  v0.0.3 |
| Sort | It sorts the elements in the stream|  v0.0.1 |
| Take | It returns a stream with the elements between the given indexes |  v0.0.3 |

## Samples

A rich and growing set of examples can be found on [koazee-samples](https://github.com/wesovilabs/koazee-samples)

## Benchmark

You can check the Benchmark for the Koazee operations [here](https://github.com/wesovilabs/koazee/wiki/Benchmark-Report)

A benchmark comparison with other frameworks can be found in [Koazee vs Go-Funk vs Go-Linq](https://medium.com/@ivan.corrales.solera/koazee-vs-go-funk-vs-go-linq-caf8ef18584e)

## Guides & Tutorials

[Shopping cart with Koazee](https://medium.com/wesovilabs/koazee-the-shopping-cart-a381bba32955)


## Roadmap

This is only the beginning! By the way, If you missed any operation in Koazee v0.0.3,  or you found a bug, please [create a new issue on Github or vote the existing ones](https://github.com/wesovilabs/koazee/issues)!


## Contributors
- [@ivancorrales](https://github.com/ivancorrales)
- [@xuyz](https://github.com/xuyz)
- [@u5surf](https://github.com/u5surf)
- [@flowonyx](https://github.com/flowonyx)
