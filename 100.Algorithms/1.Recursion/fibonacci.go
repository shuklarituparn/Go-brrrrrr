package main

import "fmt"

//underscores are used to silence the unused import

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("The fibonacci number of the integer is: ")
	fmt.Printf("%v\n", fibonacci(35)) //for the formatted output that contains the placeholders
	v := fibonacci(45)
	fmt.Print(v)
}

//the package main corresponds to entry point of the program
//while making a library we can name it not-main
