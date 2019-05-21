package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg // also quadratic, see echo_1
		sep = " "
	}
	fmt.Println(s)
	ex2()
}

func ex2() {

	for i, val := range os.Args[1:] {
		fmt.Println(i, val)
	}

}
