package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	// async concurrently call fetch for each url with the go keyword
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// retrieve messages in channel
	// block len(os.Args[1:]) times to wait for
	// len(os.Args[1:]) responses
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs seconds elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	//writes to a channel of strings
	// must write exactly once per goroutine
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		// write an error and return
		ch <- fmt.Sprintf("while GET: %s\n%v\n", url, err)
		return
	}

	// discard info from body of response
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		// write and error and return
		ch <- fmt.Sprintf("while reading: %s\n%v\n", url, err)
		return
	}

	elapsed := time.Since(start).Seconds()
	// write back results to channel
	ch <- fmt.Sprintf("%s\t%d\t%v seconds\n", url, nbytes, elapsed)
}
