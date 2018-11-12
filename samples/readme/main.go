package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var lenLowerThan6 = func(val string) bool { return len(val) <= 6 }
var concatStrings = func(acc, val string) string {
	if len(acc) == 0 {
		return val
	}
	return fmt.Sprintf("%s %s", acc, val)
}
var concatStringsWitDash = func(acc, val string) string {
	if len(acc) == 0 {
		return val
	}
	return fmt.Sprintf("%s-%s", acc, val)
}
var streamFlow = koazee.
	Stream().
	RemoveDuplicates().
	Filter(lenLowerThan6).
	Map(strings.ToUpper)

func main() {
	array := []string{"koazee", "telescope", "is", "created", "so", "great"}
	fmt.Println(streamFlow.With(array).Reduce(concatStrings).String())
	fmt.Println(streamFlow.With(array).Reduce(concatStringsWitDash).String())
}
