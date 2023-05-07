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
		log.Fatal(err) //to log an error and terminate the program
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

//fmt.Fprintf(os.Stderr, format, a...) or fmt.Fprintln(os.Stderr, a...)  can also use this instead of log
