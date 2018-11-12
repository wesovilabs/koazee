package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/logger"
	"strings"
)

var lenLowerThan6 = func(val string) bool { return len(val) <= 6 }
var concatStrings = func(acc, val string) string {
	if len(acc) == 0 {
		return val
	}
	return fmt.Sprintf("%s %s", acc, val)
}
var streamFlow = koazee.
	Stream().
	RemoveDuplicates().
	Filter(func(val string)string{
		return "This should be an invalid operation!"
	}).
	Filter(lenLowerThan6).
	Map(strings.ToUpper)

func main() {
	logger.Enabled=true
	array := []string{"koazee", "telescope", "is", "fucking", "so", "great"}
	out:=streamFlow.With(array).Out()
	if out.Err()!=nil{
		fmt.Println(out.Err())
		return
	}
	fmt.Println(out.Val().([]int))
}
