package main

import (
	"fmt"
	"io"
	//"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	for _, url := range os.Args[1:] {
		var formatted_url string

		if !strings.HasPrefix("http://", url) {
			formatted_url = "http://" + url
		} else {
			formatted_url = url
		}

		resp, err := http.Get(formatted_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch Get error: %v\n", err)
			os.Exit(1)
		}

		//body, err := ioutil.ReadAll(resp.Body)

		// io.Copy will copy bytes from resp.Body until an EOF is reached.
		// A buffer large enough to hold the entire stream is not needed
		// since the destination is to Stdout, unlike ioutil.ReadAll
		_, err1 := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch Read Body error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status: %s\n", resp.Status)
		//fmt.Printf("status: %s\nbody:\n %s\n", resp.Status, body)

	}

}
