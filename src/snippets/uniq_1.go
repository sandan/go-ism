package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)      // make a map with key of string and val of int
	input := bufio.NewScanner(os.Stdin) // read from stdin

	for input.Scan() {
		counts[input.Text()]++ // count the number of lines with same data
	}
    // use CTRL + d to send EOF

	for data, occurrence := range counts { // range over key, value pairs in map
		if occurrence > 1 {
			fmt.Printf("%d\t%s\n", occurrence, data)
		}
	}
}
