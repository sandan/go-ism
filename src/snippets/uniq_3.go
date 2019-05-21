package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Count each line in a file
// Read a file, split into lines, process each line
// O(#files * max #lines in files)
func main() {

	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// Read a file
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uniq_3: %v\n", err)
		}

		// Split into lines
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	// output lines with multiple counts
	for key, val := range counts {
		if val > 1 {
			fmt.Printf("%s\t%d\n", key, val)
		}
	}

}
