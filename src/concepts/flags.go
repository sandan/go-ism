package main

import (
	"flag"
	"fmt"
    "strings"
)

// The flag package parses command line arguments
// read $go doc flag

// define two flags
// the flag variables are pointers of specified type
// to get the value off the flag, reference the pointer
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	flag.Parse()

	// arguments are placed after the binary and given in flag.Args()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
/*
Example usage:

~/go/golang/src:$ ../bin/flags -help
Usage of ../bin/flags:
  -n	omit trailing newline
  -s string
		separator (default " ")


$ ../bin/flags -n=false -s=\\t a b c
a\tb\tc

~/go/golang/src:$ ../bin/flags -n -s=\\t a b c
a\tb\tc~/go/golang/src:$

*/
