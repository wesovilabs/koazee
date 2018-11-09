---
title: Array()
position: 4.0
type: array
description: It returns an array with the elements in the stream
content_markdown: |-

    This operation will be used after manipulating elements with the below operations
      {: .info }
left_code_blocks:
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
        Elements: []main.todo{main.todo{title:"Buy food", priority:5, effort:2}, main.todo{title:"Buy beer", priority:10, effort:1}, main.todo{title:"Pay the bills", priority:6, effort:2}}
      **/
      func main(){
          result:=koazee.Collection(todoList).Array()
          fmt.Printf("Elements: %#v",result)
      }
      
    title: Sample
    language: golang
---