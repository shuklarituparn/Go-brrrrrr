package main

import "fmt"

func main() {

	var message = "Hello World!" //here the type of the variables is inferred
	var a, b, c = 1, 2, 3        //same here

	fmt.Println(message, a, b, c)
}
