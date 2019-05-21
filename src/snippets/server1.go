package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//register a function with a request
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8181", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
