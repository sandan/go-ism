package main

import (
  "fmt"
  "text/template"
)

func main(){

 // a template is a string or file containing actions: {{..}}
 // actions trigger behaviors expressed using a dsl
 // an action has a notion of current value shown with a dot (.)
 //   it refers to the template's initial parameter

 // within an action | is the pipe symbol like in unix
 //  it makes the expression on the left the first argument to the expression on the right

 // printf is a synonym for fmt.Sprintf in all templates
 // you can create your own functions and use them in the template as well
}
