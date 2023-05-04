package main

import "fmt"

func main() {

	var message = "Hello World!"
	var a, b, c = 1, false, 3 //here we can mix the types of the variables, b is bool, a and b are int

	fmt.Println(message, a, b, c)
}
