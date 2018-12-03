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
