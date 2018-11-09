---
title: Count()
position: 1.5
type: value
description: It returns the number of elements in the stream
content_markdown: |-

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
        The output will be 3
      **/
      func main(){
          result:=koazee.Collection(todoList).Count()
          fmt.Printf("Number of elements: %d",result)
      }
      
    title: Sample
    language: golang
---