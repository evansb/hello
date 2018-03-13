package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World Forever")
}

func main() {
	var _ = 2
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
