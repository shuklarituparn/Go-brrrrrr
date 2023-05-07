package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/path/to/static/files")))
	http.ListenAndServe(":8000", nil)
}
