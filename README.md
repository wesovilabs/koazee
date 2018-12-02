[![Build Status](https://travis-ci.org/wesovilabs/koazee.svg?branch=master)](https://travis-ci.org/wesovilabs/koazee)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/koazee)](https://goreportcard.com/report/github.com/wesovilabs/koazee)
[![codecov](https://codecov.io/gh/wesovilabs/koazee/branch/master/graph/badge.svg)](https://codecov.io/gh/wesovilabs/koazee)
[![godoc](https://godoc.org/github.com/wesovilabs/koazee?status.svg)](http://godoc.org/github.com/wesovilabs/koazee)





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
