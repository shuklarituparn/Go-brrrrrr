package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.geekwiseacademy.com/")
	page, _ := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
