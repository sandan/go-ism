package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

// Behind the scenes, the server runs
// the handler for each incoming request in
// a separate goroutine so that it can serve multiple
// requests simultaneously
func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8181", nil))
}

// echoes path of requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path: %q\n", r.URL.Path)
}

// echo number of calls so far
// note that counter still needs an *http.Request even though
// it is not used since it is called by http.HandleFunc
func counter(w http.ResponseWriter, r *http.Request) {

	// To avoid race conditions on concurrent requests to
	// increment the count variable, we must ensure that at
	// most one goroutine accesses the variable at a time
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
