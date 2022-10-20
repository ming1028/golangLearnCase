package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")
}
