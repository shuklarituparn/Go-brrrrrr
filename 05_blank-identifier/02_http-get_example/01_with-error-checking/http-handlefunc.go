package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8000", nil)
}

/*
This example creates a http.HandleFunc() handler that responds with the string "Hello, World!" to any HTTP requests on
the /hello URL path. It then registers the handler and starts an HTTP server on port 8000.
*/
