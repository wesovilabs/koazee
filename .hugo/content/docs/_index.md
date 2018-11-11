+++
title = "Documentation"
description = "Find out the list of available operations"
draft = false
weight = 200
bref="Find out what Koazee can do for you"
toc = true
script = 'animation'
+++
<h3 class="section-head" id="h-plugin"><a href="#h-plugin">What is Koazee?</a></h3>
<p>
    Koazee is a handy Golang library focused on helping developers and make their life easier by taking the hassle out of working with arrays.
    It takes an array and creates an stream. The stream can be easily manipulated by making use of the provided operations by Koazee. 
</p>

<p>
    Koazee is inspired in two of the most powerful, and well-known, techniques in Software development. 
</p>
<p>
    <b>Lazy evaluation</b>: Your stream operations will only be performed when required.
    <ul>
        <li>Don't evaluate unnecessary code</li>
        <li>Do the things when they need to be done!</li>
    </ul>
</p>
<p>    
    <b>Functional programming</b>: It helps us to write better code.
    <ul>
        <li>It allows you to write more compressed and predictable code.</li>
        <li>It’s easier to test.</li>
    </ul>
</p>   

<h3 class="section-head" id="h-plugin"><a href="#h-plugin">How does Koazee work?</a></h3>
<p>
 Koazee takes advantage of reflection and simplify the way to deal with arrays. It provides us with a great and growing 
 set of powerful operations over arrays.
</p>
<p>
    The full list of available operations can be found on <a href="/koazee/docs/">Koazee Documentation</a>
</p>
<h3 class="section-head" id="h-plugin"><a href="#h-plugin">Installing</a></h3>
<p>
    Add Koazee to your project 
</p>
<nav class="tabs" data-component="tabs">
    <ul>
      <li class="active">
        <a href="#modules">Go modules</a>
      </li>
      <li>
        <a href="#glide">Glide</a>
      </li>
      <li>
        <a href="#godep">Go dep</a>
      </li>
    </ul>
</nav>
<div id="modules">    
{{< highlight text >}}
    module github.com/me/project
    require ( 
        github.com/wesovilabs/koazee vX.Y.Z
    )
{{< /highlight >}}
</div>
<div id="glide">
{{< highlight text >}}
    glide get github.com/wesovilabs/koazee
{{< /highlight >}}
</div>
<div id="godep">
{{< highlight text >}}
    go get github.com/wesovilabs/koazee
{{< /highlight >}}
</div>

<h3 class="section-head" id="h-plugin"><a href="#h-plugin">Usage</a></h3>
<p>
Once Koazee is added to the project you just need to import it 
</p>
<div>
```golang 
    
    package main
    
    import (
    	"fmt"
    	"github.com/wesovilabs/koazee"
    )
    
    var numbers = []int{1, 3, 5, 7, 9}
    var lowerThan5 = func(val int) bool { return val < 5 }
    var duplicate = func(val int) int { return val * 2 }
    var sum = func(acc, value int) int { return acc + value }
    
    func main() {
    
    	result := koazee.StreamOf(numbers).
    		Filter(lowerThan5).
    		Map(duplicate).
    		Add(2).
    		Reduce(sum).
    		Int()
    		
    	fmt.Printf("Result is %d\n", result)
    }

```
</div>
<br/>
<div>
    <blockquote>
    As a friend of mine says: "Developoing is like write poetry", and the above code looks like that
    </blocquote>
</div>