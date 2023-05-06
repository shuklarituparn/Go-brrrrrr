package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.geekwiseacademy.com/")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(res.Body)
	err = res.Body.Close()
	if err != nil {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
