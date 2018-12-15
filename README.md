[![Build Status](https://travis-ci.org/wesovilabs/koazee.svg?branch=master)](https://travis-ci.org/wesovilabs/koazee)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/koazee)](https://goreportcard.com/report/github.com/wesovilabs/koazee)
[![godoc](https://godoc.org/github.com/wesovilabs/koazee?status.svg)](http://godoc.org/github.com/wesovilabs/koazee)
[![codecov](https://codecov.io/gh/wesovilabs/koazee/branch/master/graph/badge.svg)](https://codecov.io/gh/wesovilabs/koazee)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/avelino/awesome-go#Utilities)

# Koazee

> Lazy like a koala, smart like a chimpanzee


## What is Koazee?

Koazee is a StreamLike, Inmutable, Lazy Loading and smart Golang Library to deal with slices. 


Visit the [Koazee wiki](https://github.com/wesovilabs/koazee/wiki) to find out what Koazee can do.

## Koazee Highlights

- **Inmutable**: Koazee won't modify your inputs. 
- **StreamLike**: We can combine operations up our convenience. 
- **Lazy loading**: Operations are not performed until they're required
- **Generic**: Koazee provides a generic interface able to deal with slice of any type without creating custom functions.
- **Focusing on performance**: First rule for implementing a new operation is providing the best possible performance.



## Getting started

### Stream creation

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

### Stream operations

Current release v0.0.3 (Gibbon) brings us 20 generic operations that are showed below

#### stream.At / stream.First / stream.Last
These operation return an element from the stream

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

### Samples

If you like how look the code below, that means that you should be using Koazee in your projects.


```go
package main

import (
  "fmt"
  "github.com/wesovilabs/koazee/stream"
  "koazee-samples/database"
  "strings"
)

var quotesStream = stream.New(database.GetQuotes())

func printQuotesOrderedByAuthor() stream.Stream {
  return quotesStream.
    Sort(func(quoteLeft, quoteRight *database.Quote) int {
      return strings.Compare(quoteLeft.Author, quoteRight.Author)
    }).
    ForEach(func(quote *database.Quote) {
      fmt.Printf(" * %s said \"%s\"\n", quote.Author, quote.Text)
    })
}

func numberOfAnonymousQuotes() int {
  count, _ := quotesStream.Filter(func(quote *database.Quote) bool {
    return quote.Author == "Anonymous"
  }).Count()
  return count
}

func listOfAuthors() stream.Stream {
  return quotesStream.
    Map(func(quote *database.Quote) string {
      return quote.Author
    }).
  	RemoveDuplicates().
    Sort(strings.Compare)
}

func printNameOfAuthors() stream.Stream {
  return listOfAuthors().
    ForEach(func(author string) {
      fmt.Printf(" * %s\n", author)
    })
}

func quotesByAuthorOrderedByQuoteLen(author string) stream.Stream {
  return quotesStream.
    Filter(func(quote *database.Quote) bool {
      return quote.Author == author
    }).
    Map(func(quote *database.Quote) string {
      return quote.Text
    }).
    Sort(func(quote1, quote2 string) int {
      if len(quote1) > len(quote2) {
        return 1
      } else if len(quote1) < len(quote2) {
        return -1
      }
      return 0
    }).
    ForEach(func(quote string) {
      fmt.Println(quote)
    })
}



func main() {
  count, _ := quotesStream.Count()
  fmt.Printf("\n - Total quotes: %d\n", count)
  fmt.Printf("\n - Total anonymous quotes: %d\n", numberOfAnonymousQuotes())
  count, _ = listOfAuthors().Count()
  fmt.Printf("\n - Total authors: %d\n", count)
  fmt.Println("\nPrinting quotes ordered by author name")
  fmt.Println("--------------------------------------")
  printQuotesOrderedByAuthor().Do()
  fmt.Println("\nPrinting list of authors sorted by name")
  fmt.Println("--------------------------------------")
  printNameOfAuthors().Do()
  fmt.Println("\nPrinting list of quotes sorted bylen of quote and said by Albert Einstein")
  quotesByAuthorOrderedByQuoteLen("Albert Einstein").Do()
  fmt.Println("\nPrinting list of quotes sorted bylen of quote and said by anonymous")
  quotesByAuthorOrderedByQuoteLen("Anonymous").Do()
}
```

### Benchmark

You can check the Benchmark for the Koazee operations [here](https://github.com/wesovilabs/koazee/wiki/Benchmark-Report)

A benchmark compison with other frameworks can be found in [Koazee vs Go-Funk vs Go-Linq](https://medium.com/@ivan.corrales.solera/koazee-vs-go-funk-vs-go-linq-caf8ef18584e)

### Guides & Tutorials

[Shopping cart with Koazee](https://medium.com/wesovilabs/koazee-the-shopping-cart-a381bba32955)


*If you like this project and you think I should provide more functionality to Koazee, please feel free to star the repository*


