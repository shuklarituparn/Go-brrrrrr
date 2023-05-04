package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 42, 42, 42) //%x just prints hexadecimal without format
	//	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42) //#x prints the hexadecimal in format
}
