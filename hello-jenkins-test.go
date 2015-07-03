package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s", "Test")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":1337", nil)
}
