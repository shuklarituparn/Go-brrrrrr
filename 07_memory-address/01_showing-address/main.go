package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)                   //the value of a
	fmt.Println("a's memory address - ", &a) //the memory address of a
	fmt.Printf("%d \n", &a)                  //memory address of a in decimal notation
}
