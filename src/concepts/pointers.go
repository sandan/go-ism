package main

import (
	"fmt"
)

func main() {

	// a variable is a piece of storage that has a value
	// a pointer is a variable whose value is an address
	//   - similar to C, it can read and update a variable's value indirectly

	// not every value has an address, but every variable does have an address
	// (either on the program stack or heap)

	var x int
	var ptr_x *int = &x // *int type is a pointer to an int
	// & operator gets the address of variable x
	// ptr_x is a variable that is said to "point" to x

	// The * operator can be used to reference the value in the address
	// contained in the pointer variable ptr_x
	fmt.Println("value of x: %d", *ptr_x)

	*ptr_x = 21
	fmt.Println("value of x: %d", x)
}
