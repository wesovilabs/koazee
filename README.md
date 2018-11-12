[![Build Status](https://travis-ci.org/wesovilabs/koazee.svg?branch=master)](https://travis-ci.org/wesovilabs/koazee)
[![Go Report Card](https://goreportcard.com/badge/github.com/wesovilabs/koazee)](https://goreportcard.com/report/github.com/wesovilabs/koazee)
[![codecov](https://codecov.io/gh/wesovilabs/koazee/branch/master/graph/badge.svg)](https://codecov.io/gh/wesovilabs/koazee)
[![godoc](https://godoc.org/github.com/wesovilabs/koazee?status.svg)](http://godoc.org/github.com/wesovilabs/koazee)





# Koazee

## Lazy like a koala, smart like a chimpanzee

Koazee full documentation is hosted on [http://wesovilabs.com/koazee](http://wesovilabs.com/koazee)


### What is Koazee?

Koazee is a handy Golang library focused on helping developers and make their life easier by taking the hassle out of working with arrays.
It takes an array and creates a stream. The stream can be easily manipulated by making use of the provided operations by Koazee.

### How does Koazee work?


**Koazee is inspired in two of the most powerful, and well-known, techniques in Software development.**

- **Lazy evaluation**: Don't do things unless they are need to be done!
- **Functional programming**: Let's write clearer code more compressed and predictable


### Samples

If you like how look the code below, that means that you should be using Koazee in your projects.


```golang
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
```


#### How can I check an error in runtime?

You always can check if any operation fail during the runtime

```golang
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
	array := []string{"koazee", "telescope", "is", "created", "so", "great"}
	out:=streamFlow.With(array).Out()
	if out.Err()!=nil{
		fmt.Println(out.Err())
		return
	}
	fmt.Println(out.Val().([]int))
}
```

If we run the above code the output will be

```bash
[koazee] 05:56:18.821227 with  [koazee telescope is created so great]
[koazee] 05:56:18.821596 removeDuplicates [koazee telescope is created so great] -> [koazee telescope is created so great]
[filter:err.invalid-argument] The type of the output in the provided function must be bool
```


You can find more documentation and examples on [http://wesovilabs.com/koazee](http://wesovilabs.com/koazee)

*If you like this project and you think I should provide more functionality to Koazee, please feel free to star the repository*
