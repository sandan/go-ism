package main

import (
	"fmt"
)

func main() {

	// The new(T) expression creates an unnamed variable of type T
	// with the zero value of types for T,
	// and returns  a pointer to type T (*T)

    // The go compiler decides basd on the lifetime of the variable,
    // whether new allocates the variable on the stack or heap (see variable_lifetime)
	var x = new(string)
	fmt.Println("*x:", *x)

	*x = "hello!"
	fmt.Println("*x:", *x)
}

/*
A variable created with new is no different from an ordinary local variable whose address is
taken, except that there’s no need to invent (and declare) a dummy name, and we can use
new(T) in an expression. Thus new is only a syntactic convenience, not a fundamental notion
*/
