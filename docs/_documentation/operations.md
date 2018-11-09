---
title: Operations
position: 4
parameters:
  - name:
    content:
content_markdown: |-
  
  As it was mentioned on the above sections, koazee takes
  the arrays as streams. The below Golang interface represents
  the already implemented functions.
  
  We will go through any of the provided operations next.
  
left_code_blocks:
  - code_block: |-
      type S interface {
      	First() ()interface{},error)
      	Last() ()interface{},error)
      	At(int) ()interface{},error)
      	Count() ()int,error)
      	Array() ()interface{},error)
      	Contains(interface{}) ()bool,error)
      	Reduce(function interface{}) ()interface{},error)
      	Apply(interface{}) S
      	Filter(function interface{}) S
      	RemoveDuplicates() S
      	Error() error
      }
    title: Stream Interface
    language: golang
---