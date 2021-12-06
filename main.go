package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8090", nil)
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello word (changed)\n")
}
