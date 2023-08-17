package main

import "fmt"

func main() {
	m := make([]string, 1, 25)  //length of the splice is 1, capacity is 25
	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) // [Todd]
}

//A splice has a three thing in itself. a pointer to the data, the length of the slice, and the capacity of the slice.