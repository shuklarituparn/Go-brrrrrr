package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	fmt.Println("Response headers:", res.Header)
}
