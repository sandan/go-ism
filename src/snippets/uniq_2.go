package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int) //make a map with keys of type string and values of type int
	files := os.Args[1:]           // read filenames from stdin

	if len(files) == 0 {

		countLines(os.Stdin, counts)  // if no files are given then use stdin

	} else {
		for _, arg := range files {

			// open a file
			f, err := os.Open(arg)

			if err != nil {
				//print an error message if unable to open
				fmt.Fprintf(os.Stderr, "uniq_2: %v\n", err)
				continue // process next file, note that the file does not have to be closed
			}

			countLines(f, counts)
			f.Close()
		}
	}

	// print entries in map with more than one count
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}

// helper function for counting lines in a file f given a map counts
func countLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
