package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"koazee-samples-shopping/cart"
	"log"
	"strings"
	"time"
)

const discount = 0.8 // percentage to payed when buy 3 or more items

// ItemTotal used to append the calculated priced to the items
type ItemTotal struct {
	*cart.Item
	total float32
}

var shoppingStream = koazee.Stream().
	Filter(func(item *cart.Item) bool {
		return item.ExpirationDate.After(time.Now())
	}).
	Sort(func(itemLeft, itemRight *cart.Item) int {
		return strings.Compare(itemLeft.Name, itemRight.Name)
	}).
	Map(func(item *cart.Item) *ItemTotal {
		total := float32(item.Units) * item.PricePerUnit
		if item.Units >= 3 {
			total *= discount
		}
		return &ItemTotal{item, total}
	}).
	ForEach(func(item *ItemTotal) {
		fmt.Printf(" - %s, %d units, %.2f€\n", item.Name, item.Units, item.total)
	})

func process(items []*cart.Item) {
	myStream := shoppingStream.With(items).Do()
	output := myStream.Reduce(
		func(acc float32, item *ItemTotal) float32 {
			return acc + item.total
		})
	if output.Err() != nil {
		log.Fatal(output.Err().Error())
	}
	fmt.Printf(" Total price %.2f€\n", output.Float32())
}

func main() {
	fmt.Println(":--------------------------:")
	process(cart.Items)
	fmt.Println(":--------------------------:")
	process(cart.Items2)
	fmt.Println(":--------------------------:")
}
