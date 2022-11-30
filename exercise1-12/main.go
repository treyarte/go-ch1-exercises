package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/treyarte/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/counter", counter)
	http.HandleFunc("/animate", animate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter: %d\n", count)
	mu.Unlock()
}

func animate(w http.ResponseWriter, r *http.Request) {
	lissajous.Lissajous(w)
}