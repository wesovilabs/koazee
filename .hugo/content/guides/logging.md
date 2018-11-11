+++
title = "Logging with Koazee"
description = "Learn to log what koazee is doing"
draft = false
weight = 200
bref="Learn to log what koazee is doing"
toc = true
script = 'animation'
+++
<h3 class="section-head" id="h-whatis"><a href="#h-whatis">Enabling Koazee logging</a></h3>
<p>
Koazee provides support to log what it is happening in the stream. In the below
example we will see how of easy is enable it. 

Actually, we just need to set a global variable

```golang
logger.Enabled=true
```
</p>

<h3 class="section-head" id="h-whatis"><a href="#h-whatis">Example</a></h3>

<div id="logging">
{{< highlight golang >}}
package main

import (
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/logger"
)

var numbers = []int{1, 2, 3, 4, 4, 1}

func main() {
	logger.Enabled = true
	koazee.Stream().
		With(numbers).
		Filter(func(val int) bool {
			return val%2 == 0
		}).
		Map(func(val int) int {
			return val * 2
		}).
		RemoveDuplicates().
		Reduce(func(acc, val int) int {
			return acc + val
		})
}
{{< /highlight >}}
<p>
If we run the above code could see in the stdout the below info:
{{< highlight bash >}}
[koazee] 19:11:45.668742 with  [1 2 3 4 4 1]
[koazee] 19:11:45.669116 filter [1 2 3 4 4 1] -> [2 4 4]
[koazee] 19:11:45.669141 map [2 4 4] -> [4 8 8]
[koazee] 19:11:45.669152 removeDuplicates [4 8 8] -> [4 8]
[koazee] 19:11:45.669168 reduce [4 8] -> 12
{{< /highlight >}}
</p>
</div>
