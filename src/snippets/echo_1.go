package main

import (
	"fmt"
	"os"
)

func main() {

	var s, sep string // var declares variables.
	// The default zero value is used for the type if not explicitly given.

	for i := 1; i < len(os.Args); i++ {

		/*
		   The for loop is the only loop statement in go

		   for initialization; condition; post {
		   // zero or more statements
		   }
		   Parentheses are never used around the three components of a for loop. The braces are
		   mandatory, however, and the opening brace must be on the same line as the post statement

		   - initialization is optional
		   - omitting initialization and post gives a while loop
		   - omitting all three gives an infinite loop
		*/
		// for loop decl uses := for short variable declaration
		// declares variable(s) with implied types based on literals
		// ++ and -- are postfix operators only not prefix,
		// Increment/Decrement are statements (no values) not expressions (has a value)

		s += sep + os.Args[i] // + concats strings,
		// note that this is a quadratic process
		// new elements are not appended to the end of a string (immutable)
		// but a new string is built after every iteration
		// performance is likely fine for cmd line args...
		sep = " "

	}
	fmt.Println(s)

}
