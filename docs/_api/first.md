---
title: First()
position: 1.0
type: access
description: It returns the first element in the stream.
content_markdown: |-

left_code_blocks:
  - code_block: |-
      package main
      
      import (
        "fmt"
        "github.com/wesovilabs/koazee"
      )
      
      var elements = []int{4,5,6,7}
      
      /**
        The output will be
        First element: 4
      **/
      func main(){
          result,_:=koazee.Collection(elements).First()
          fmt.Printf("First element: %#v",result)
      }
      
    title: Simple
    language: golang
  - code_block: |-
      package main
      
      import (
        "fmt"
        "github.com/wesovilabs/koazee"
      )
      
      type todo struct {
        title       string
        priority    int
        effort       int
      }
      
      var todoList = []todo{
      	{
      		title:       "Buy food",
      		priority:    5,
      		effort:      2,
      	},
      	{
      		title:       "Buy beer",
      		priority:    10,
      		effort:       1,
      	},
      	{
      		title:       "Pay the bills",
      		priority:    6,
      		effort:       2,
      	},
      }
      
      /**
        The output will be
        First element: main.todo{title:"Buy food", priority:5, effort:2}
      **/
      func main(){
          result,_:=koazee.Collection(todoList).First()
          fmt.Printf("First element: %#v",result)
      }
      
    title: Struct
    language: golang

  - code_block: |-
      package main
      
      import (
        "fmt"
        "github.com/wesovilabs/koazee"
      )
      
      type todo struct {
        title       string
        priority    int
        effort       int
      }
      
      var todoList = []*todo{
      	{
      		title:       "Buy food",
      		priority:    5,
      		effort:      2,
      	},
      	{
      		title:       "Buy beer",
      		priority:    10,
      		effort:       1,
      	},
      	{
      		title:       "Pay the bills",
      		priority:    6,
      		effort:       2,
      	},
      }
      
      /**
        The output will be
        First element: &main.todo{title:"Buy food", priority:5, effort:2}
      **/
      func main(){
          result,_:=koazee.Collection(todoList).First()
          fmt.Printf("First element: %#v",result)
      }
      
    title: Pointer
    language: golang
---