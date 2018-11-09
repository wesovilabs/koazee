---
title: Last()
position: 1.1
type: access    
description: It returns the last element in the stream.
content_markdown: |-

left_code_blocks:
  - code_block: |-
        package main
            
        import (
          "fmt"
          "github.com/wesovilabs/koazee"
        )
                
        var elements = []string{ "Freedom","animals"}
        
        /**
          The output will be
          Last element: "animals"
        **/
        func main(){
            result,_:=koazee.Collection(elements).Last()
            fmt.Printf("Last element: %#v",result)
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
          main.todo{title:"Pay the bills", priority:6, effort:2}
        **/
        func main(){
            result,_:=koazee.Collection(todoList).Last()
            fmt.Printf("Last element: %#v",result)
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
          main.todo{title:"Pay the bills", priority:6, effort:2}
        **/
        func main(){
            result,_:=koazee.Collection(todoList).Last()
            fmt.Printf("Last element: %#v",result)
        }
      
    title: Pointer
    language: golang
---