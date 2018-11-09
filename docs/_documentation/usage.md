---
title: Usage
position: 3
content_markdown: |-
  
  Once koazee is added to the project you just need to
  import it 

left_code_blocks:
 - code_block: |-
    import "github.com/wesovilabs/koazee"
    
    ...
    
    var todoList = []todo{
    	{
    		title:       "Buy food",
    		priority:    mediumPriority,
    		timeToSpend: 2,
    	},
    	{
    		title:       "Buy beer",
    		priority:    highPriority,
    		timeToSpend: 1,
    	},
    }
    
    func main(){
        stream:=koazee.Collection(todoList)
        ...
    }
   language: golang
    
---