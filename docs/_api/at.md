---
title: At( index )
position: 1.2
type: access
description: It returns the element in the stream, in the given position
parameters:
  - name: index
    content: int value for taking the element in this position.
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
        Element at position 1 is 5
      **/
      func main(){
          index:=1
          result,_:=koazee.Collection(elements).At(index)
          fmt.Printf("Element at position %d is %d", index, result)
      }
      
    title: Basic
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
          Element at position 1 is main.todo{title:"Buy beer", priority:10, effort:1}
        **/
        func main(){
            index:=1
            result,_:=koazee.Collection(todoList).At(index)
            fmt.Printf("Element at position %d is %#v", index, result)
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
            Element at position 1 is &main.todo{title:"Buy beer", priority:10, effort:1}
          **/
          func main(){
              index:=1
              result,_:=koazee.Collection(todoList).At(index)
              fmt.Printf("Element at position %d is %#v", index, result)
          }
        
    title: Pointer
    language: golang
---