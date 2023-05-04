package main

import "fmt"

func main() {

	// you can only do this inside a func
	message := "Hello World!"
	a, b, c := 1, false, 3 //using the shorthand notation to iniatialize three var at once
	d := 4
	e := true

	fmt.Println(message, a, b, c, d, e)
}
