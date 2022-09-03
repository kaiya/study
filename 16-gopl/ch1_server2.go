package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path) // Fprintf(Writer, format, ...interface{})
	w.Header().Add("Server", "nginx")
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprintf(w, "Count %d\n", count)
}
