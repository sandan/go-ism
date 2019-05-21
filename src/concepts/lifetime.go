package main

import (
	"fmt"
)

var global *int

/*
The lifetime of a variable is the interval of time during which it exists as the program executes.

The lifetime of a package-level variable is the entire execution of the program.

By contrast, local variables have dynamic lifetimes: a new instance is created each time the declaration
statement is executed, and the variable lives on until it becomes unreachable, at which point its
storage may be recycled.

Function parameters and results are local variables too; they are created
each time their enclosing function is called.
*/
func main() {

	x()
	y()


	fmt.Println("address: ", global, "value: ", *global)
}

func x() {

	// x is a local variable but may "escape from the function x()",
	// the enclosing function, by aliasing it to a variable
	// outside the scope of the function
	var x int = 1
	global = &x

	// x will not be garbage collected just yet since there is an alias around
	// in the package level, so x is allocated on the heap
}

func y() {


	// the garbage collector looks at whether or not the variable
	// is reachable or not. If it is no longer reachable,
	// then it can no longer affect the computation, so it gc's the variable, 
    // which it likely allocated on the stack frame of the function y()
	y := new(string)
	*y = "will be garbage collected"

}

//The lifetime of a variable is determined only by whether or not it is reachable

/*
  variable lifetime good to keep in mind during performance optimization,
  since each variable that escapes requires an extra memory al location.

  Garbage collection is a tremendous help in writing correct programs, but it does not relieve
  you of the burden of thinking about memory.

  You don’t need to explicitly allocate and free memory, but to write efficient programs
  you still need to be aware of the lifetime of variables.

For example, keeping unnecessary pointers to short-lived objects within long-lived objects,
especially global variables, will prevent the garbage collector from reclaiming the short-lived objects
*/
