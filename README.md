[![Build Status](https://travis-ci.org/wesovilabs/koazee.svg?branch=master)](https://travis-ci.org/wesovilabs/koazee)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/koazee)](https://goreportcard.com/report/github.com/wesovilabs/koazee)
[![godoc](https://godoc.org/github.com/wesovilabs/koazee?status.svg)](http://godoc.org/github.com/wesovilabs/koazee)
[![codecov](https://codecov.io/gh/wesovilabs/koazee/branch/master/graph/badge.svg)](https://codecov.io/gh/wesovilabs/koazee)


# Koazee

> Lazy like a koala, smart like a chimpanzee


### What is Koazee?

Koazee is a handy Golang library focused on helping developers and make their life easier by taking the hassle out of working with arrays. It takes an array and creates an stream. The stream can be easily manipulated by making use of the provided operations by Koazee. 

Koazee is inspired in two of the most powerful, and well-known, techniques in Software development.

Visit the [Koazee wiki](https://github.com/wesovilabs/koazee/wiki) to find out what Koazee can do.

### How does Koazee work?

**Koazee is inspired in two of the most powerful, and well-known, techniques in Software development.**

- **Lazy evaluation**: Don't do things unless they are need to be done!
- **Functional programming**: Let's write clearer code more compressed and predictable


### Koazee Operations

| Operation  | Description  | Since  |
|---|---|---|
| [At](https://github.com/wesovilabs/koazee/wiki/Operation-at) | It returns the element in the given position |  v0.0.1 |
| [First](https://github.com/wesovilabs/koazee/wiki/Operation-first) | It returns the element in the first position |  v0.0.1 |
| [Last](https://github.com/wesovilabs/koazee/wiki/Operation-last) | It returns the element in the last position |  v0.0.1 |
| [Add](https://github.com/wesovilabs/koazee/wiki/Operation-add) | It adds a new element in the last position |  v0.0.1 |
| [Pop](https://github.com/wesovilabs/koazee/wiki/Operation-pop) | It extracts the first element in the stream and return this and the new stream | master |
| [Drop](https://github.com/wesovilabs/koazee/wiki/Operation-drop) | It removes an element from the stream |  v0.0.1 |
| [DeleteAt](https://github.com/wesovilabs/koazee/wiki/Operation-deleteAt) | It remove the elements in the given position |  master |
| [Take](https://github.com/wesovilabs/koazee/wiki/Operation-take) | It returns a stream with the elements between the given indexes |  master |
| [Map](https://github.com/wesovilabs/koazee/wiki/Operation-map) | It converts the element in the stream |  v0.0.1 |
| [Filter](https://github.com/wesovilabs/koazee/wiki/Operation-filter) | It discards those elements that doesn't match with the provided filter|  v0.0.1 |
| [Reduce](https://github.com/wesovilabs/koazee/wiki/Operation-reduce) | It reduceshe stream to a single value by executing a provided function for each value of the stream|  v0.0.1 |
| [ForEach](https://github.com/wesovilabs/koazee/wiki/Operation-foreach) | It does something over all the elements in the stream.|  v0.0.1 |
| [RemoveDuplicates](https://github.com/wesovilabs/koazee/wiki/Operation-removeDuplcates) | It removes duplicated elements.|  v0.0.1 |
| [Count](https://github.com/wesovilabs/koazee/wiki/Operation-count) | It returns the number of elements in a stream|  v0.0.1 |
| [Contains](https://github.com/wesovilabs/koazee/wiki/Operation-contains) | It checks if the given element is found in the stream.|  v0.0.1 |
| [Sort](https://github.com/wesovilabs/koazee/wiki/Operation-sort) | It sorts the elements in the stream by the given function.|  v0.0.1 |
| [Reverse](https://github.com/wesovilabs/koazee/wiki/Operation-reverse) | It reverses the sequence of elements in the stream.|  master |


### Benchmark

You can check the Benchmark for the Koazee operations [here](https://github.com/wesovilabs/koazee/wiki/Benchmark-Report)

A benchmark compison with other frameworks can be found in [Koazee vs Go-Funk vs Go-Linq](https://medium.com/@ivan.corrales.solera/koazee-vs-go-funk-vs-go-linq-caf8ef18584e)

### Guides & Tutorials

[Shopping cart with Koazee](https://medium.com/wesovilabs/koazee-the-shopping-cart-a381bba32955)

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

*If you like this project and you think I should provide more functionality to Koazee, please feel free to star the repository*
