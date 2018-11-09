---
title: Installation
position: 2
parameters:
  - name:
    content:
content_markdown: |-
  
  Add koazee to your project 

left_code_blocks:
 - code_block: |-
    module github.com/my/thing
    
    require (
        github.com/wesovilabs/koazee v0.0.1
    )
   title: Go modules
   language: bash
 - code_block: |-
    glide get github.com/wesovilabs/koazee
   title: Glide
   language: bash
 - code_block: |-
    go get github.com/wesovilabs/koazee
   title: Go dep
   language: bash
 
---